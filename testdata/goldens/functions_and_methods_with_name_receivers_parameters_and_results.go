package testdata

import "testing"

func Test_name_Name(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		rname name
		// Parameters.
		n string
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.rname.Name(tt.n); got != tt.want {
			t.Errorf("%q. name.Name() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestName_Name1(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rName string
		// Parameters.
		n string
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		rname := &Name{
			Name: tt.rName,
		}
		if got := rname.Name1(tt.n); got != tt.want {
			t.Errorf("%q. Name.Name1() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestName_Name2(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rName string
		// Parameters.
		pname string
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.rName,
		}
		if got := n.Name2(tt.pname); got != tt.want {
			t.Errorf("%q. Name.Name2() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestName_Name3(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rName string
		// Parameters.
		nn string
		// Expected results.
		wantName string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.rName,
		}
		if gotName := n.Name3(tt.nn); gotName != tt.wantName {
			t.Errorf("%q. Name.Name3() = %v, want %v", tt.name, gotName, tt.wantName)
		}
	}
}
