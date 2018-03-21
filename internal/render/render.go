package render

//go:generate go-bindata -pkg=bindata -o "./bindata/bindata.go" templates
import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/cweill/gotests/internal/models"
	"github.com/cweill/gotests/internal/options"
	"github.com/cweill/gotests/internal/render/bindata"
)

const name = "name"

var (
	tmpls *template.Template
)

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

func TestFunction(w io.Writer, f *models.Function, opt *options.Options) error {
	return tmpls.ExecuteTemplate(w, "function", struct {
		*models.Function
		PrintInputs bool
		Subtests    bool
		ArgsStruct  bool
	}{
		Function:    f,
		PrintInputs: opt.PrintInputs,
		Subtests:    opt.Subtests,
		ArgsStruct:  shouldGenerateArgsStruct(opt.ArgsStructMode, f),
	})
}

func shouldGenerateArgsStruct(mode options.ArgsStruct, f *models.Function) bool {
	switch mode {
	case options.ArgsStructAlways:
		return true
	case options.ArgsStructNever:
		return false
	case options.ArgsStructSmart:
		return len(f.Parameters) > 1
	default:
		panic(fmt.Sprintf("Unhandled args struct mode %v", mode))
	}
}
