package output

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/tools/imports"

	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/render"
)

// we do not need support for aliases in import for now.
var importsMap = map[string]string{
	"testify": "github.com/stretchr/testify/assert",
}

type Options struct {
	PrintInputs    bool
	Subtests       bool
	Parallel       bool
	Named          bool
	Template       string
	TemplateDir    string
	TemplateParams map[string]interface{}
	TemplateData   [][]byte
}

func (o *Options) providesTemplateData() bool { return o != nil && len(o.TemplateData) > 0 }
func (o *Options) providesTemplateDir() bool  { return o != nil && o.TemplateDir != "" }
func (o *Options) providesTemplate() bool     { return o != nil && o.Template != "" }

func Process(head *models.Header, funcs []*models.Function, opt *Options) ([]byte, error) {
	switch {
	case opt.providesTemplateDir():
		if err := render.LoadCustomTemplates(opt.TemplateDir); err != nil {
			return nil, fmt.Errorf("loading custom templates: %v", err)
		}
	case opt.providesTemplate():
		if err := render.LoadCustomTemplatesName(opt.Template); err != nil {
			return nil, fmt.Errorf("loading custom templates of name: %v", err)
		}
	case opt.providesTemplateData():
		render.LoadFromData(opt.TemplateData)
	default:
		render.Reset()
	}

	tf, err := ioutil.TempFile("", "gotests_")
	if err != nil {
		return nil, fmt.Errorf("ioutil.TempFile: %v", err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())
	b := &bytes.Buffer{}
	if err := writeTests(b, head, funcs, opt); err != nil {
		return nil, err
	}

	out, err := imports.Process(tf.Name(), b.Bytes(), nil)
	if err != nil {
		return nil, fmt.Errorf("imports.Process: %v", err)
	}
	return out, nil
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func writeTests(w io.Writer, head *models.Header, funcs []*models.Function, opt *Options) error {
	if path, ok := importsMap[opt.Template]; ok {
		head.Imports = append(head.Imports, &models.Import{
			Path: fmt.Sprintf(`"%s"`, path),
		})
	}

	b := bufio.NewWriter(w)
	if err := render.Header(b, head); err != nil {
		return fmt.Errorf("render.Header: %v", err)
	}

	for _, fun := range funcs {
		if err := render.TestFunction(b, fun, opt.PrintInputs, opt.Subtests, opt.Named, opt.Parallel, opt.TemplateParams); err != nil {
			return fmt.Errorf("render.TestFunction: %v", err)
		}
	}
	return b.Flush()
}
