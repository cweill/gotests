package render

import (
	"io"
	"tester/models"
	"text/template"
)

var generatedTemplate = template.Must(template.New("testcases").Parse(`package {{.Package}}

import (
	"testing"
)
{{range .Funcs}}
func Test{{.Name}}(t *testing.T) {
	{{$multi := .ReturnsMultiple}}
	tests := []struct {
		name string{{range .Parameters}}{{if .Name}}
		{{.Name}} {{.Type}}{{end}}{{end}}{{range .Results}}{{if not .IsError}}{{if and .Name $multi}}	
		want{{.Name}} {{.Type}}{{else if $multi}}
		want{{.Type}} {{.Type}}{{else}}
		want {{.Type}}{{end}}{{end}}{{end}}{{if .ReturnsError}}
		wantErr bool{{end}}
	}{}
	for _, tt := range tests {
		t.Logf("Running: %v", tt.name)
		got{{if .ReturnsError}}, err{{end}} := {{.Name}}({{range .Parameters}}tt.{{.Name}},{{end}}){{if .ReturnsError}}
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. {{.Name}}() error: %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}{{end}}
		if got != tt.want {
			t.Errorf("%v. {{.Name}}() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
{{end}}`))

func TestCases(w io.Writer, info *models.Info) error {
	return generatedTemplate.Execute(w, info)
}
