package render

import (
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"text/template"

	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/render/bindata"
	"github.com/cweill/gotests/templates"
)

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
	for _, name := range bindata.AssetNames() {
		r.tmpls = template.Must(r.tmpls.Parse(bindata.FSMustString(false, name)))
	}

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
	f, err := templates.Dir(false, "/").Open(name)
	if err != nil {
		return fmt.Errorf("templates.Open: %v", err)
	}

	files, err := f.Readdir(nFile)
	if err != nil {
		return fmt.Errorf("f.Readdir: %v", err)
	}

	for _, f := range files {
		text, err := templates.FSString(false, path.Join("/", name, f.Name()))
		if err != nil {
			return fmt.Errorf("templates.FSString: %v", err)
		}

		if tmpls, err := r.tmpls.Parse(text); err != nil {
			return fmt.Errorf("tmpls.Parse: %v", err)
		} else {
			r.tmpls = tmpls
		}
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
	useGoCmp bool,
	params map[string]interface{},
) error {
	return r.tmpls.ExecuteTemplate(w, "function", struct {
		*models.Function
		PrintInputs    bool
		Subtests       bool
		Parallel       bool
		Named          bool
		UseGoCmp       bool
		TemplateParams map[string]interface{}
	}{
		Function:       f,
		PrintInputs:    printInputs,
		Subtests:       subtests,
		Parallel:       parallel,
		Named:          named,
		UseGoCmp:       useGoCmp,
		TemplateParams: params,
	})
}
