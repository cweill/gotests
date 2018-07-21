package testdata

import "testing"

func Test_name_Name(t *testing.T) {
	tests := []struct {
		name string
		n    name
		arg  string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.Name(tt.arg); got != tt.want {
			t.Errorf("%q. name.Name() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestName_Name1(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		arg    string
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.fields.Name,
		}
		if got := n.Name1(tt.arg); got != tt.want {
			t.Errorf("%q. Name.Name1() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestName_Name2(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		arg    string
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.fields.Name,
		}
		if got := n.Name2(tt.arg); got != tt.want {
			t.Errorf("%q. Name.Name2() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestName_Name3(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name     string
		fields   fields
		arg      string
		wantName string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.fields.Name,
		}
		if gotName := n.Name3(tt.arg); gotName != tt.wantName {
			t.Errorf("%q. Name.Name3() = %v, want %v", tt.name, gotName, tt.wantName)
		}
	}
}
