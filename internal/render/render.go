package render

//go:generate go-bindata -pkg=bindata -o "./bindata/bindata.go" templates
import (
	"fmt"
	"io"
	"text/template"
	"unicode"

	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/render/bindata"
)

const name = "name"

var tmpls *template.Template

func init() {
	tmpls = template.New("render").Funcs(map[string]interface{}{
		"Field":    fieldName,
		"Receiver": receiverName,
		"Param":    parameterName,
		"Want":     wantName,
		"Got":      gotName,
	})
	for _, name := range bindata.AssetNames() {
		tmpls = template.Must(tmpls.Parse(string(bindata.MustAsset(name))))
	}
}

func fieldName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = unexport(f.Name)
	} else {
		n = unexport(f.Type.String())
	}
	if n == name {
		return "f" + n
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
	if n == name {
		return "r" + n
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
	if n == name {
		return "p" + n
	}
	return n
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
