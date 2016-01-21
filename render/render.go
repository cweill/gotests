package render

import (
	"io"
	"path"
	"runtime"
	"text/template"

	"github.com/cweill/gotests/models"
)

var tmpls *template.Template

func init() {
	_, filename, _, _ := runtime.Caller(1)
	tmpls = template.Must(template.ParseGlob(path.Join(path.Dir(filename), "templates/*.tmpl")))
}

func Header(w io.Writer, info *models.Header) error {
	return tmpls.ExecuteTemplate(w, "header", info)
}

func TestFunction(w io.Writer, f *models.Function) error {
	return tmpls.ExecuteTemplate(w, "testfunction", f)
}
