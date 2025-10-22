package output

import (
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
