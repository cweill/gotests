package render

import "testing"

func TestUnexport(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"Field", "field"},
		{"FieldName", "fieldName"},
		{"XML", "xml"},
		{"XMLDocument", "xmlDocument"},
		{"XMLDocumentXML", "xmlDocumentXML"},
	}
	for _, tt := range tests {
		if got := unexport(tt.s); got != tt.want {
			t.Errorf("unexport(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
