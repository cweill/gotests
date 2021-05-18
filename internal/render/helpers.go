package render

//go:generate esc -o bindata/esc.go -pkg=bindata templates
import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/cweill/gotests/internal/models"
)

const (
	name  = "name"
	nFile = 7
)

var tmpls *template.Template

func fieldName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = f.Name
	} else {
		n = f.Type.String()
	}
	return n
}

func receiverName(f *models.Receiver) string {
	var n string
	if f.IsNamed() {
		n = f.Name
	} else {
		n = f.ShortName()
	}
	if n == "name" {
		// Avoid conflict with test struct's "name" field.
		n = "n"
	} else if n == "t" {
		// Avoid conflict with test argument.
		// "tr" is short for t receiver.
		n = "tr"
	}
	return n
}

func parameterName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = f.Name
	} else {
		n = fmt.Sprintf("in%v", f.Index)
	}
	return n
}

func wantName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = "want" + strings.Title(f.Name)
	} else if f.Index == 0 {
		n = "want"
	} else {
		n = fmt.Sprintf("want%v", f.Index)
	}
	return n
}

func gotName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = "got" + strings.Title(f.Name)
	} else if f.Index == 0 {
		n = "got"
	} else {
		n = fmt.Sprintf("got%v", f.Index)
	}
	return n
}

func Header(w io.Writer, h *models.Header) error {
	if err := tmpls.ExecuteTemplate(w, "header", h); err != nil {
		return err
	}
	_, err := w.Write(h.Code)
	return err
}

func TestFunction(w io.Writer, f *models.Function, printInputs, subtests, named, parallel bool, templateParams map[string]interface{}) error {
	return tmpls.ExecuteTemplate(w, "function", struct {
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
		TemplateParams: templateParams,
	})
}
