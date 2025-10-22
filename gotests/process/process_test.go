package process

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		args []string
		opts *Options
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Nil options and nil args",
			args: nil,
			opts: nil,
			want: specifyFlagMessage + "\n",
		}, {
			name: "Nil options",
			args: []string{"testdata/foobar.go"},
			opts: nil,
			want: specifyFlagMessage + "\n",
		}, {
			name: "Empty options",
			args: []string{"testdata/foobar.go"},
			opts: &Options{},
			want: specifyFlagMessage + "\n",
		}, {
			name: "Non-empty options with no args",
			args: []string{},
			opts: &Options{AllFuncs: true},
			want: specifyFileMessage + "\n",
		}, {
			name: "OnlyFuncs option w/ no matches",
			args: []string{"testdata/foobar.go"},
			opts: &Options{OnlyFuncs: "FooBar"},
			want: "No tests generated for testdata/foobar.go\n",
		}, {
			name: "Invalid OnlyFuncs option",
			args: []string{"testdata/foobar.go"},
			opts: &Options{OnlyFuncs: "??"},
			want: "Invalid -only regex: error parsing regexp: missing argument to repetition operator: `??`\n",
		}, {
			name: "Invalid ExclFuncs option",
			args: []string{"testdata/foobar.go"},
			opts: &Options{ExclFuncs: "??"},
			want: "Invalid -excl regex: error parsing regexp: missing argument to repetition operator: `??`\n",
		},
	}
	for _, tt := range tests {
		out := &bytes.Buffer{}
		Run(out, tt.args, tt.opts)
		if got := out.String(); got != tt.want {
			t.Errorf("%q. Run() =\n%v, want\n%v", tt.name, got, tt.want)
		}
	}
}

func TestRun_WithValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	srcPath := filepath.Join(tmpDir, "test.go")
	content := `package mypackage

func Add(a, b int) int {
	return a + b
}
`
	if err := os.WriteFile(srcPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		opts *Options
	}{
		{
			name: "all funcs",
			opts: &Options{AllFuncs: true},
		},
		{
			name: "exported funcs",
			opts: &Options{ExportedFuncs: true},
		},
		{
			name: "with only regex",
			opts: &Options{OnlyFuncs: "Add"},
		},
		{
			name: "with excl regex",
			opts: &Options{AllFuncs: true, ExclFuncs: "Subtract"},
		},
		{
			name: "with print inputs",
			opts: &Options{AllFuncs: true, PrintInputs: true},
		},
		{
			name: "with subtests",
			opts: &Options{AllFuncs: true, Subtests: true},
		},
		{
			name: "with parallel",
			opts: &Options{AllFuncs: true, Parallel: true},
		},
		{
			name: "with named",
			opts: &Options{AllFuncs: true, Named: true},
		},
		{
			name: "with go-cmp",
			opts: &Options{AllFuncs: true, UseGoCmp: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			Run(out, []string{srcPath}, tt.opts)
			output := out.String()
			if output == "" {
				t.Error("Run() produced no output")
			}
		})
	}
}

func TestRun_WithTemplateParams(t *testing.T) {
	tmpDir := t.TempDir()
	srcPath := filepath.Join(tmpDir, "test.go")
	content := `package mypackage

func Greet(name string) string {
	return "Hello, " + name
}
`
	if err := os.WriteFile(srcPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	jsonParams := `{"customParam": "value"}`
	opts := &Options{
		AllFuncs:       true,
		TemplateParams: jsonParams,
	}

	out := &bytes.Buffer{}
	Run(out, []string{srcPath}, opts)
	output := out.String()
	if output == "" {
		t.Error("Run() with template params produced no output")
	}
}

func TestRun_InvalidTemplateParams(t *testing.T) {
	tmpDir := t.TempDir()
	srcPath := filepath.Join(tmpDir, "test.go")
	content := `package mypackage

func Test() {}
`
	if err := os.WriteFile(srcPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	opts := &Options{
		AllFuncs:       true,
		TemplateParams: `{invalid json`,
	}

	out := &bytes.Buffer{}
	Run(out, []string{srcPath}, opts)
	output := out.String()

	// Should produce an error message about invalid JSON
	if output == "" {
		t.Error("Run() should produce error message for invalid JSON")
	}
}
