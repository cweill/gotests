package render

import (
	"gentests/models"
	"io"
	"text/template"
)

var tmpls = template.Must(template.ParseGlob("render/templates/*.tmpl"))

func Header(w io.Writer, info *models.Info) error {
	return tmpls.ExecuteTemplate(w, "header", info)
}

func TestCases(w io.Writer, info *models.Info) error {
	return tmpls.ExecuteTemplate(w, "testcases", info)
}
