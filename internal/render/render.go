package render

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"path"
	"text/template"

	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/templates"
)

//go:embed templates/*
var data embed.FS

type Render struct {
	tmpls *template.Template
}

func New() *Render {
	r := Render{
		tmpls: template.New("render").Funcs(map[string]interface{}{
			"Field":    fieldName,
			"Receiver": receiverName,
			"Param":    parameterName,
			"Want":     wantName,
			"Got":      gotName,
		}),
	}

	// default templates first
	r.tmpls = template.Must(r.tmpls.ParseFS(data, "templates/*.tmpl"))

	return &r
}

// LoadCustomTemplates allows to load in custom templates from a specified path.
func (r *Render) LoadCustomTemplates(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("ioutil.ReadDir: %v", err)
	}

	templateFiles := []string{}
	for _, f := range files {
		templateFiles = append(templateFiles, path.Join(dir, f.Name()))
	}
	r.tmpls, err = r.tmpls.ParseFiles(templateFiles...)
	if err != nil {
		return fmt.Errorf("tmpls.ParseFiles: %v", err)
	}

	return nil
}

// LoadCustomTemplatesName allows to load in custom templates of a specified name from the templates directory.
func (r *Render) LoadCustomTemplatesName(name string) error {
	fileSystem, err := fs.Sub(templates.FS, name)
	if err != nil {
		return fmt.Errorf("templates.Sub: %w", err)
	}

	r.tmpls, err = r.tmpls.ParseFS(fileSystem, "*.tmpl")
	if err != nil {
		return fmt.Errorf("templates.ParseFS: %w", err)
	}

	return nil
}

// LoadFromData allows to load from a data slice
func (r *Render) LoadFromData(templateData [][]byte) {
	for _, d := range templateData {
		r.tmpls = template.Must(r.tmpls.Parse(string(d)))
	}
}

func (r *Render) Header(w io.Writer, h *models.Header) error {
	if err := r.tmpls.ExecuteTemplate(w, "header", h); err != nil {
		return err
	}
	_, err := w.Write(h.Code)
	return err
}

func (r *Render) TestFunction(
	w io.Writer,
	f *models.Function,
	printInputs bool,
	subtests bool,
	named bool,
	parallel bool,
	params map[string]interface{}) error {
	return r.tmpls.ExecuteTemplate(w, "function", struct {
		*models.Function
		PrintInputs    bool
		Subtests       bool
		Parallel       bool
		Named          bool
		TemplateParams map[string]interface{}
	}{
		Function:       f,
		PrintInputs:    printInputs,
		Subtests:       subtests,
		Parallel:       parallel,
		Named:          named,
		TemplateParams: params,
	})
}
