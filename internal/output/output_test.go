package output

import "testing"

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
