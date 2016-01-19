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

	testCasesTmpl = template.Must(template.New("testcases").Parse(`{{range .Funcs}}{{$multi := .ReturnsMultiple}}func {{.TestName}}(t *testing.T) {
	tests := []struct {
		name string{{range .Parameters}}{{if .Name}}
		{{.Name}} {{.Type}}{{end}}{{end}}{{range .Results}}{{if not .IsError}}{{if and .Name $multi}}	
		want{{.Name}} {{.Type}}{{else if $multi}}
		want{{.Type}} {{.Type}}{{else}}
		want {{.Type}}{{end}}{{end}}{{end}}{{if .ReturnsError}}
		wantErr bool{{end}}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Logf("Running: %v", tt.name)
		got{{if .ReturnsError}}, err{{end}} := {{.Name}}({{range $index, $element := .Parameters}}{{if $index}}, {{end}}tt.{{.Name}}{{end}}){{if .ReturnsError}}
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. {{.Name}}() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}{{end}}
		if got != tt.want {
			t.Errorf("%v. {{.Name}}() = %v, want %v", tt.name, got, tt.want)
		}
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
