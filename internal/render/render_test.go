package render

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/cweill/gotests/internal/models"
)

func TestNew(t *testing.T) {
	r := New()
	if r == nil {
		t.Fatal("New() returned nil")
	}
	if r.tmpls == nil {
		t.Error("New() did not initialize templates")
	}
}

func TestRender_LoadCustomTemplates(t *testing.T) {
	tests := []struct {
		name    string
		setupFn func(t *testing.T) string
		wantErr bool
	}{
		{
			name: "valid template directory",
			setupFn: func(t *testing.T) string {
				tmpDir := t.TempDir()
				// Create a valid template file
				tmplContent := `{{define "header"}}package {{.Package}}{{end}}`
				tmplPath := filepath.Join(tmpDir, "header.tmpl")
				if err := os.WriteFile(tmplPath, []byte(tmplContent), 0644); err != nil {
					t.Fatal(err)
				}
				return tmpDir
			},
			wantErr: false,
		},
		{
			name: "nonexistent directory",
			setupFn: func(t *testing.T) string {
				return "/nonexistent/path"
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New()
			dir := tt.setupFn(t)
			err := r.LoadCustomTemplates(dir)

			if (err != nil) != tt.wantErr {
				t.Errorf("Render.LoadCustomTemplates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRender_LoadCustomTemplatesName(t *testing.T) {
	tests := []struct {
		name     string
		tmplName string
		wantErr  bool
	}{
		{
			name:     "testify template",
			tmplName: "testify",
			wantErr:  false,
		},
		{
			name:     "nonexistent template",
			tmplName: "nonexistent",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New()
			err := r.LoadCustomTemplatesName(tt.tmplName)

			if (err != nil) != tt.wantErr {
				t.Errorf("Render.LoadCustomTemplatesName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRender_LoadFromData(t *testing.T) {
	r := New()
	templateData := [][]byte{
		[]byte(`{{define "mytemplate"}}test{{end}}`),
	}

	// Should not panic
	r.LoadFromData(templateData)

	if r.tmpls == nil {
		t.Error("LoadFromData() resulted in nil templates")
	}
}

func TestRender_Header(t *testing.T) {
	r := New()
	buf := &bytes.Buffer{}

	header := &models.Header{
		Package: "testpkg",
		Imports: []*models.Import{
			{Path: `"testing"`},
		},
	}

	err := r.Header(buf, header)
	if err != nil {
		t.Errorf("Render.Header() error = %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Error("Render.Header() produced empty output")
	}

	// Check that output contains package declaration
	if !bytes.Contains([]byte(output), []byte("package testpkg")) {
		t.Error("Render.Header() output does not contain package declaration")
	}
}

func TestRender_TestFunction(t *testing.T) {
	r := New()
	buf := &bytes.Buffer{}

	fn := &models.Function{
		Name:       "Add",
		IsExported: true,
		Parameters: []*models.Field{
			{Name: "a", Type: &models.Expression{Value: "int"}, Index: 0},
			{Name: "b", Type: &models.Expression{Value: "int"}, Index: 1},
		},
		Results: []*models.Field{
			{Type: &models.Expression{Value: "int"}, Index: 0},
		},
	}

	err := r.TestFunction(buf, fn, false, false, false, false, false, nil)
	if err != nil {
		t.Errorf("Render.TestFunction() error = %v", err)
	}

	output := buf.String()
	if output == "" {
		t.Error("Render.TestFunction() produced empty output")
	}

	// Check that output contains test function name
	if !bytes.Contains([]byte(output), []byte("TestAdd")) {
		t.Error("Render.TestFunction() output does not contain test function name")
	}
}

func TestRender_TestFunction_WithOptions(t *testing.T) {
	tests := []struct {
		name         string
		printInputs  bool
		subtests     bool
		named        bool
		parallel     bool
		useGoCmp     bool
		checkContent string
	}{
		{
			name:         "with subtests",
			subtests:     true,
			checkContent: "t.Run",
		},
		{
			name:         "with parallel",
			subtests:     true,
			parallel:     true,
			checkContent: "t.Parallel",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New()
			buf := &bytes.Buffer{}

			fn := &models.Function{
				Name:       "TestFunc",
				IsExported: true,
				Parameters: []*models.Field{
					{Name: "x", Type: &models.Expression{Value: "int"}},
				},
				Results: []*models.Field{
					{Type: &models.Expression{Value: "int"}},
				},
			}

			err := r.TestFunction(buf, fn, tt.printInputs, tt.subtests, tt.named, tt.parallel, tt.useGoCmp, nil)
			if err != nil {
				t.Errorf("Render.TestFunction() error = %v", err)
			}

			if tt.checkContent != "" && !bytes.Contains(buf.Bytes(), []byte(tt.checkContent)) {
				t.Errorf("Render.TestFunction() output does not contain %q", tt.checkContent)
			}
		})
	}
}
