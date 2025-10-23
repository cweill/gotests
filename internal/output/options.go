package output

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/cweill/gotests/internal/ai"
	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/render"
	"golang.org/x/tools/imports"
)

// Options contains configuration for generating and formatting test output.
type Options struct {
	PrintInputs    bool
	Subtests       bool
	Parallel       bool
	Named          bool
	UseGoCmp       bool
	Template       string
	TemplateDir    string
	TemplateParams map[string]interface{}
	TemplateData   [][]byte
	UseAI          bool
	AIModel        string
	AIEndpoint     string
	AIMinCases     int
	AIMaxCases     int

	render *render.Render
}

// Process generates formatted test code from the header and function signatures.
func (o *Options) Process(head *models.Header, funcs []*models.Function) ([]byte, error) {
	o.render = render.New()

	switch {
	case o.providesTemplateDir():
		if err := o.render.LoadCustomTemplates(o.TemplateDir); err != nil {
			return nil, fmt.Errorf("loading custom templates: %v", err)
		}
	case o.providesTemplate():
		if err := o.render.LoadCustomTemplatesName(o.Template); err != nil {
			return nil, fmt.Errorf("loading custom templates of name: %v", err)
		}
	case o.providesTemplateData():
		o.render.LoadFromData(o.TemplateData)
	}

	//
	tf, err := ioutil.TempFile("", "gotests_")
	if err != nil {
		return nil, fmt.Errorf("ioutil.TempFile: %v", err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	// create physical copy of test
	b := &bytes.Buffer{}
	if err := o.writeTests(b, head, funcs); err != nil {
		return nil, err
	}

	// format file
	out, err := imports.Process(tf.Name(), b.Bytes(), nil)
	if err != nil {
		return nil, fmt.Errorf("imports.Process: %v", err)
	}
	return out, nil
}

func (o *Options) providesTemplateData() bool {
	return o != nil && len(o.TemplateData) > 0
}

func (o *Options) providesTemplateDir() bool {
	return o != nil && o.TemplateDir != ""
}

func (o *Options) providesTemplate() bool {
	return o != nil && o.Template != ""
}

func (o *Options) writeTests(w io.Writer, head *models.Header, funcs []*models.Function) error {
	if path, ok := importsMap[o.Template]; ok {
		head.Imports = append(head.Imports, &models.Import{
			Path: fmt.Sprintf(`"%s"`, path),
		})
	}

	// Add go-cmp import if needed
	if o.UseGoCmp {
		head.Imports = append(head.Imports, &models.Import{
			Path: `"github.com/google/go-cmp/cmp"`,
		})
	}

	// Initialize AI provider if needed
	var provider ai.Provider
	if o.UseAI {
		cfg := &ai.Config{
			Provider:       "ollama",
			Model:          o.AIModel,
			Endpoint:       o.AIEndpoint,
			MinCases:       o.AIMinCases,
			MaxCases:       o.AIMaxCases,
			MaxRetries:     3,  // Default: 3 retries
			RequestTimeout: 60, // Default: 60 seconds
			HealthTimeout:  2,  // Default: 2 seconds
		}
		var err error
		provider, err = ai.NewOllamaProvider(cfg)
		if err != nil {
			return fmt.Errorf("failed to create AI provider: %w", err)
		}

		// Check if Ollama is available
		if !provider.IsAvailable() {
			return fmt.Errorf("AI provider %s is not available - ensure Ollama is running at %s", provider.Name(), o.AIEndpoint)
		}
	}

	b := bufio.NewWriter(w)
	if err := o.render.Header(b, head); err != nil {
		return fmt.Errorf("render.Header: %v", err)
	}

	// Use context with timeout to prevent AI generation from hanging indefinitely
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	for _, fun := range funcs {
		var aiCases []interface{}

		// Generate AI test cases if enabled
		if o.UseAI && provider != nil {
			cases, err := provider.GenerateTestCases(ctx, fun)
			if err != nil {
				// Log warning but continue with empty cases (fallback to TODO)
				fmt.Fprintf(os.Stderr, "Warning: failed to generate AI test cases for %s: %v\n", fun.Name, err)
			} else {
				// Convert []ai.TestCase to []interface{} for template
				aiCases = make([]interface{}, len(cases))
				for i, tc := range cases {
					aiCases[i] = tc
				}
			}
		}

		err := o.render.TestFunction(b, fun, o.PrintInputs, o.Subtests, o.Named, o.Parallel, o.UseGoCmp, o.TemplateParams, aiCases)
		if err != nil {
			return fmt.Errorf("render.TestFunction: %v", err)
		}
	}

	return b.Flush()
}
