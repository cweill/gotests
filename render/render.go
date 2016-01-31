package render

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"text/template"

	"github.com/cweill/gotests/models"
)

var tmpls *template.Template

func init() {
	_, filename, _, _ := runtime.Caller(1)
	tmpls = template.Must(template.New("render").Funcs(map[string]interface{}{
		"Receiver": receiverName,
		"Param":    parameterName,
		"Want":     wantName,
		"Got":      gotName,
	}).ParseGlob(path.Join(path.Dir(filename), "templates/*.tmpl")))
}

func receiverName(f *models.Field) string {
	if f.IsNamed() {
		return f.Name
	}
	return f.ShortName()
}

func parameterName(f *models.Field, i int) string {
	if f.IsNamed() {
		return f.Name
	}
	return fmt.Sprintf("in%v", i)
}

func wantName(i int) string {
	if i == 0 {
		return "want"
	}
	return fmt.Sprintf("want%v", i)
}

func gotName(i int) string {
	if i == 0 {
		return "got"
	}
	return fmt.Sprintf("got%v", i)
}

func Header(w io.Writer, h *models.Header) error {
	if err := tmpls.ExecuteTemplate(w, "header", h); err != nil {
		return err
	}
	if _, err := w.Write(h.Code); err != nil {
		return err
	}
	return nil
}

func TestFunction(w io.Writer, f *models.Function, printInputs bool) error {
	return tmpls.ExecuteTemplate(w, "function", struct {
		*models.Function
		PrintInputs bool
	}{
		Function:    f,
		PrintInputs: printInputs,
	})
}
