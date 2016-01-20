package render

import (
	"io"
	"tester/models"
	"text/template"
)

var (
	headerTmpl = template.Must(template.New("testcases").Parse(`package {{.Package}}

import (
	"testing"
)

`))

	testCasesTmpl = template.Must(template.New("testcases").Parse(`{{range .Funcs}}{{$f := .}}func {{.TestName}}(t *testing.T) {
	tests := []struct {
		name string{{if .Receiver}}
		{{.Receiver.Name}} {{.Receiver.Type}}{{end}}{{range $index, $element := .Parameters}}{{if .Name}}
		{{.Name}} {{.Type}}{{else}}
		in{{$index}} {{.Type}}{{end}}{{end}}{{range .Results}}{{if and .Name $f.ReturnsMultiple}}	
		want{{.Name}} {{.Type}}{{else if $f.ReturnsMultiple}}
		want{{.Type}} {{.Type}}{{else}}
		want {{.Type}}{{end}}{{end}}{{if .ReturnsError}}
		wantErr bool{{end}}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		{{if .Results}}got{{if .ReturnsError}}, err{{end}} := {{else if .ReturnsError}}err := {{end}}{{if .Receiver}}tt.{{.Receiver.Name}}.{{end}}{{.Name}}({{range $index, $element := .Parameters}}{{if $index}}, {{end}}{{if .Name}}tt.{{.Name}}{{else}}tt.in{{$index}}{{end}}{{end}}){{if .ReturnsError}}
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. {{.Name}}() error = %v, wantErr: %v", tt.name, err, tt.wantErr){{if .Results}}
			continue{{end}}
		}{{end}}{{range .Results}}{{if .IsScalar}}
		if got != tt.want {
			t.Errorf("%v. {{$f.Name}}() = %v, want %v", tt.name, got, tt.want)
		}{{else}}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. {{$f.Name}}() = %v, want %v", tt.name, got, tt.want)
		}{{end}}{{end}}
	}
}

{{end}}`))
)

func Header(w io.Writer, info *models.Info) error {
	return headerTmpl.Execute(w, info)
}

func TestCases(w io.Writer, info *models.Info) error {
	return testCasesTmpl.Execute(w, info)
}
