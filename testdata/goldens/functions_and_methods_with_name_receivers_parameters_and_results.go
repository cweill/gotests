package testdata

import "testing"

func TestNameName(t *testing.T) {
	tests := []struct {
		name  string
		rname name
		n     string
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.rname.Name(tt.n); got != tt.want {
			t.Errorf("%q. name.Name() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNameName1(t *testing.T) {
	tests := []struct {
		name  string
		fname string
		n     string
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		rname := &Name{
			Name: tt.fname,
		}
		if got := rname.Name1(tt.n); got != tt.want {
			t.Errorf("%q. Name.Name1() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNameName2(t *testing.T) {
	tests := []struct {
		name  string
		fname string
		pname string
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.fname,
		}
		if got := n.Name2(tt.pname); got != tt.want {
			t.Errorf("%q. Name.Name2() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNameName3(t *testing.T) {
	tests := []struct {
		name     string
		fname    string
		nn       string
		wantName string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.fname,
		}
		if gotName := n.Name3(tt.nn); gotName != tt.wantName {
			t.Errorf("%q. Name.Name3() = %v, want %v", tt.name, gotName, tt.wantName)
		}
	}
}
