package render

//go:generate go-bindata -pkg=render templates

import (
	"fmt"
	"io"
	"text/template"
	"unicode"

	"github.com/cweill/gotests/internal/models"
)

var tmpls *template.Template

func init() {
	tmpls = template.New("render").Funcs(map[string]interface{}{
		"Field":    fieldName,
		"Receiver": receiverName,
		"Param":    parameterName,
		"Want":     wantName,
		"Got":      gotName,
	})
	for _, name := range AssetNames() {
		tmpls = template.Must(tmpls.Parse(string(MustAsset(name))))
	}
}

func fieldName(f *models.Field) string {
	if f.IsNamed() {
		return unexport(f.Name)
	}
	return unexport(f.Type.String())
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

func unexport(s string) string {
	r := []rune(s)
	for i := range r {
		if i != 0 && i+1 < len(r)-1 && unicode.IsLower(r[i+1]) {
			break
		}
		r[i] = unicode.ToLower(r[i])
	}
	return string(r)
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
