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

func Header(w io.Writer, info *models.Info) error {
	return tmpls.ExecuteTemplate(w, "header", info)
}

func TestCases(w io.Writer, info *models.Info) error {
	return tmpls.ExecuteTemplate(w, "testcases", info)
}
