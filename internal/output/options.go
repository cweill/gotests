package output

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/render"
	"golang.org/x/tools/imports"
)

type Options struct {
	PrintInputs    bool
	Subtests       bool
	Parallel       bool
	Named          bool
	Template       string
	TemplateDir    string
	TemplateParams map[string]interface{}
	TemplateData   [][]byte
	UseGoCmp       bool

	render *render.Render
}

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

	b := bufio.NewWriter(w)
	if err := o.render.Header(b, head); err != nil {
		return fmt.Errorf("render.Header: %v", err)
	}

	for _, fun := range funcs {
		err := o.render.TestFunction(b, fun, o.PrintInputs, o.Subtests, o.Named, o.Parallel, o.UseGoCmp, o.TemplateParams)
		if err != nil {
			return fmt.Errorf("render.TestFunction: %v", err)
		}
	}

	return b.Flush()
}
