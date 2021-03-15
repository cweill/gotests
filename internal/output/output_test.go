package output

import "testing"

func TestOptions_isProvidesTemplateData(t *testing.T) {
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
			if got := tt.otpions.isProvidesTemplateData(); got != tt.want {
				t.Errorf("Options.isProvidesTemplateData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_isProvidesTemplate(t *testing.T) {
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
			if got := tt.otpions.isProvidesTemplate(); got != tt.want {
				t.Errorf("Options.isProvidesTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_isProvidesTemplateDir(t *testing.T) {
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
			if got := tt.otpions.isProvidesTemplateDir(); got != tt.want {
				t.Errorf("Options.isProvidesTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
