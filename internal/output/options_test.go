package output

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cweill/gotests/internal/models"
)

func TestOptions_Process(t *testing.T) {
	tests := []struct {
		name    string
		options *Options
		head    *models.Header
		funcs   []*models.Function
		wantErr bool
	}{
		{
			name:    "simple function test generation",
			options: &Options{},
			head: &models.Header{
				Package: "mypackage",
				Imports: []*models.Import{},
			},
			funcs: []*models.Function{
				{
					Name:       "Add",
					IsExported: true,
					Parameters: []*models.Field{
						{Name: "a", Type: &models.Expression{Value: "int"}},
						{Name: "b", Type: &models.Expression{Value: "int"}},
					},
					Results: []*models.Field{
						{Type: &models.Expression{Value: "int"}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "with subtests enabled",
			options: &Options{
				Subtests: true,
			},
			head: &models.Header{
				Package: "mypackage",
			},
			funcs: []*models.Function{
				{
					Name:       "Multiply",
					IsExported: true,
					Parameters: []*models.Field{
						{Name: "x", Type: &models.Expression{Value: "int"}},
					},
					Results: []*models.Field{
						{Type: &models.Expression{Value: "int"}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "with print inputs",
			options: &Options{
				PrintInputs: true,
			},
			head: &models.Header{
				Package: "mypackage",
			},
			funcs: []*models.Function{
				{
					Name:       "Divide",
					IsExported: true,
					Parameters: []*models.Field{
						{Name: "a", Type: &models.Expression{Value: "int"}},
						{Name: "b", Type: &models.Expression{Value: "int"}},
					},
					Results: []*models.Field{
						{Type: &models.Expression{Value: "int"}},
					},
					ReturnsError: true,
				},
			},
			wantErr: false,
		},
		{
			name: "with go-cmp",
			options: &Options{
				UseGoCmp: true,
			},
			head: &models.Header{
				Package: "mypackage",
			},
			funcs: []*models.Function{
				{
					Name:       "GetUser",
					IsExported: true,
					Results: []*models.Field{
						{Type: &models.Expression{Value: "User"}},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.options.Process(tt.head, tt.funcs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Options.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) == 0 {
				t.Error("Options.Process() returned empty output")
			}
		})
	}
}

func TestOptions_Process_WithTemplateDir(t *testing.T) {
	tmpDir := t.TempDir()
	// Create a valid template file
	tmplContent := `{{define "header"}}package {{.Package}}{{end}}`
	tmplPath := tmpDir + "/header.tmpl"
	if err := os.WriteFile(tmplPath, []byte(tmplContent), 0644); err != nil {
		t.Fatal(err)
	}

	options := &Options{
		TemplateDir: tmpDir,
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
			Parameters: []*models.Field{
				{Name: "a", Type: &models.Expression{Value: "int"}},
			},
			Results: []*models.Field{
				{Type: &models.Expression{Value: "int"}},
			},
		},
	}

	got, err := options.Process(head, funcs)
	if err != nil {
		t.Errorf("Options.Process() with TemplateDir error = %v", err)
		return
	}
	if len(got) == 0 {
		t.Error("Options.Process() returned empty output")
	}
}

func TestOptions_Process_WithTemplate(t *testing.T) {
	options := &Options{
		Template: "testify",
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
			Parameters: []*models.Field{
				{Name: "a", Type: &models.Expression{Value: "int"}},
			},
			Results: []*models.Field{
				{Type: &models.Expression{Value: "int"}},
			},
		},
	}

	got, err := options.Process(head, funcs)
	if err != nil {
		t.Errorf("Options.Process() with Template error = %v", err)
		return
	}
	if len(got) == 0 {
		t.Error("Options.Process() returned empty output")
	}
}

func TestOptions_Process_WithTemplateData(t *testing.T) {
	options := &Options{
		TemplateData: [][]byte{
			[]byte(`{{define "header"}}package {{.Package}}{{end}}`),
		},
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
			Parameters: []*models.Field{
				{Name: "a", Type: &models.Expression{Value: "int"}},
			},
			Results: []*models.Field{
				{Type: &models.Expression{Value: "int"}},
			},
		},
	}

	got, err := options.Process(head, funcs)
	if err != nil {
		t.Errorf("Options.Process() with TemplateData error = %v", err)
		return
	}
	if len(got) == 0 {
		t.Error("Options.Process() returned empty output")
	}
}

func TestOptions_Process_WithInvalidTemplateDir(t *testing.T) {
	options := &Options{
		TemplateDir: "/nonexistent/path",
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
		},
	}

	_, err := options.Process(head, funcs)
	if err == nil {
		t.Error("Options.Process() with invalid TemplateDir should return error")
	}
}

func TestOptions_Process_WithInvalidTemplate(t *testing.T) {
	options := &Options{
		Template: "nonexistent",
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
		},
	}

	_, err := options.Process(head, funcs)
	if err == nil {
		t.Error("Options.Process() with invalid Template should return error")
	}
}

func TestOptions_Process_WithNamed(t *testing.T) {
	options := &Options{
		Named: true,
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
			Parameters: []*models.Field{
				{Name: "a", Type: &models.Expression{Value: "int"}},
			},
			Results: []*models.Field{
				{Type: &models.Expression{Value: "int"}},
			},
		},
	}

	got, err := options.Process(head, funcs)
	if err != nil {
		t.Errorf("Options.Process() with Named error = %v", err)
		return
	}
	if len(got) == 0 {
		t.Error("Options.Process() returned empty output")
	}
}

func TestOptions_Process_WithParallel(t *testing.T) {
	options := &Options{
		Parallel: true,
		Subtests: true, // Parallel requires subtests
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
			Parameters: []*models.Field{
				{Name: "a", Type: &models.Expression{Value: "int"}},
			},
			Results: []*models.Field{
				{Type: &models.Expression{Value: "int"}},
			},
		},
	}

	got, err := options.Process(head, funcs)
	if err != nil {
		t.Errorf("Options.Process() with Parallel error = %v", err)
		return
	}
	if len(got) == 0 {
		t.Error("Options.Process() returned empty output")
	}
}

func TestOptions_providesTemplateData(t *testing.T) {
	tests := []struct {
		name    string
		otpions *Options
		want    bool
	}{
		{"Opt is nil", nil, false},
		{"Opt is empty", &Options{}, false},
		{"TemplateData is nil", &Options{TemplateData: nil}, false},
		{"TemplateData is empty", &Options{TemplateData: [][]byte{}}, false},
		{"TemplateData is OK", &Options{TemplateData: [][]byte{[]byte("ok")}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.otpions.providesTemplateData(); got != tt.want {
				t.Errorf("Options.isProvidesTemplateData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_providesTemplate(t *testing.T) {
	tests := []struct {
		name    string
		otpions *Options
		want    bool
	}{
		{"Opt is nil", nil, false},
		{"Opt is empty", &Options{}, false},
		{"Template is empty (implicit_zero_val)", &Options{Template: ""}, false},
		{"Template is OK", &Options{Template: "testify"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.otpions.providesTemplate(); got != tt.want {
				t.Errorf("Options.isProvidesTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_providesTemplateDir(t *testing.T) {
	tests := []struct {
		name    string
		otpions *Options
		want    bool
	}{
		{"Opt is nil", nil, false},
		{"Opt is empty", &Options{}, false},
		{"Template is empty", &Options{TemplateDir: ""}, false},
		{"Template is OK", &Options{TemplateDir: "testify"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.otpions.providesTemplateDir(); got != tt.want {
				t.Errorf("Options.isProvidesTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_Process_WithAI_Success(t *testing.T) {
	// Create a mock Ollama server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/tags" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"models": []interface{}{}})
			return
		}
		if r.URL.Path == "/api/generate" {
			response := map[string]interface{}{
				"response": `tests := []struct {
					name string
					args args
					want int
					wantErr bool
				}{
					{
						name: "basic addition",
						args: args{a: 1, b: 2},
						want: 3,
						wantErr: false,
					},
				}`,
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	options := &Options{
		UseAI:      true,
		AIModel:    "test-model",
		AIEndpoint: server.URL,
		AIMinCases: 3,
		AIMaxCases: 3,
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
			Parameters: []*models.Field{
				{Name: "a", Type: &models.Expression{Value: "int"}},
				{Name: "b", Type: &models.Expression{Value: "int"}},
			},
			Results: []*models.Field{
				{Type: &models.Expression{Value: "int"}},
			},
		},
	}

	got, err := options.Process(head, funcs)
	if err != nil {
		t.Errorf("Options.Process() with AI error = %v", err)
		return
	}
	if len(got) == 0 {
		t.Error("Options.Process() with AI returned empty output")
	}
}

func TestOptions_Process_WithAI_ProviderCreationError(t *testing.T) {
	options := &Options{
		UseAI:      true,
		AIModel:    "test-model",
		AIEndpoint: "invalid://endpoint",  // Invalid URL will cause provider creation error
		AIMinCases: 3,
		AIMaxCases: 3,
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
		},
	}

	_, err := options.Process(head, funcs)
	if err == nil {
		t.Error("Options.Process() with invalid AI endpoint should return error")
	}
}

func TestOptions_Process_WithAI_ProviderUnavailable(t *testing.T) {
	// Create a server that returns 404 for /api/tags (unavailable)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	options := &Options{
		UseAI:      true,
		AIModel:    "test-model",
		AIEndpoint: server.URL,
		AIMinCases: 3,
		AIMaxCases: 3,
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
		},
	}

	_, err := options.Process(head, funcs)
	if err == nil {
		t.Error("Options.Process() with unavailable AI provider should return error")
	}
	if err != nil && !contains(err.Error(), "not available") {
		t.Errorf("Error should mention provider not available, got: %v", err)
	}
}

func TestOptions_Process_WithAI_GenerationError(t *testing.T) {
	// Create a mock server that returns an error for generation
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/tags" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{"models": []interface{}{}})
			return
		}
		if r.URL.Path == "/api/generate" {
			// Return invalid response to trigger parsing error
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"response": "invalid go code that will fail parsing",
			})
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	options := &Options{
		UseAI:      true,
		AIModel:    "test-model",
		AIEndpoint: server.URL,
		AIMinCases: 3,
		AIMaxCases: 3,
	}
	head := &models.Header{
		Package: "mypackage",
	}
	funcs := []*models.Function{
		{
			Name:       "Add",
			IsExported: true,
			Parameters: []*models.Field{
				{Name: "a", Type: &models.Expression{Value: "int"}},
			},
			Results: []*models.Field{
				{Type: &models.Expression{Value: "int"}},
			},
		},
	}

	// Should not error - it should fall back to TODO comments
	got, err := options.Process(head, funcs)
	if err != nil {
		t.Errorf("Options.Process() with AI generation error should not fail, should fallback to TODO: %v", err)
		return
	}
	if len(got) == 0 {
		t.Error("Options.Process() should still generate output even with AI error (fallback)")
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 &&
		(s == substr || len(s) >= len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
