package render

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"text/template"

	"github.com/cweill/gotests/internal/models"
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

func receiverName(f *models.Receiver) string {
	if f.IsNamed() {
		return f.Name
	}
	return f.ShortName()
}

func parameterName(f *models.Field) string {
	if f.IsNamed() {
		return f.Name
	}
	return fmt.Sprintf("in%v", f.Index)
}

func wantName(f *models.Field) string {
	if f.Index == 0 {
		return "want"
	}
	return fmt.Sprintf("want%v", f.Index)
}

func gotName(f *models.Field) string {
	if f.Index == 0 {
		return "got"
	}
	return fmt.Sprintf("got%v", f.Index)
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
