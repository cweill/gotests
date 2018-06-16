package testdata

import "testing"

func Test_name_Name(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		n    name
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.n.Name(tt.args.n); got != tt.want {
			t.Errorf("%q. name.Name() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Benchmark_name_Name(b *testing.B) {
	type args struct {
		n string
	}
	benchmarks := []struct {
		name string
		n    name
		args args
		want string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		if got := tt.n.Name(tt.args.n); got != bb.want {
			b.Errorf("%q. name.Name() = %v, want %v", tt.name, got, bb.want)
		}
	}
}

func TestName_Name1(t *testing.T) {
	type fields struct {
		Name string
	}
	type args struct {
		n string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.fields.Name,
		}
		if got := n.Name1(tt.args.n); got != tt.want {
			t.Errorf("%q. Name.Name1() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkName_Name1(b *testing.B) {
	type fields struct {
		Name string
	}
	type args struct {
		n string
	}
	benchmarks := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		n := &Name{
			Name: bb.fields.Name,
		}
		if got := n.Name1(tt.args.n); got != bb.want {
			b.Errorf("%q. Name.Name1() = %v, want %v", tt.name, got, bb.want)
		}
	}
}

func TestName_Name2(t *testing.T) {
	type fields struct {
		Name string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.fields.Name,
		}
		if got := n.Name2(tt.args.name); got != tt.want {
			t.Errorf("%q. Name.Name2() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkName_Name2(b *testing.B) {
	type fields struct {
		Name string
	}
	type args struct {
		name string
	}
	benchmarks := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		n := &Name{
			Name: bb.fields.Name,
		}
		if got := n.Name2(tt.args.name); got != bb.want {
			b.Errorf("%q. Name.Name2() = %v, want %v", tt.name, got, bb.want)
		}
	}
}

func TestName_Name3(t *testing.T) {
	type fields struct {
		Name string
	}
	type args struct {
		nn string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantName string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.fields.Name,
		}
		if gotName := n.Name3(tt.args.nn); gotName != tt.wantName {
			t.Errorf("%q. Name.Name3() = %v, want %v", tt.name, gotName, tt.wantName)
		}
	}
}

func BenchmarkName_Name3(b *testing.B) {
	type fields struct {
		Name string
	}
	type args struct {
		nn string
	}
	benchmarks := []struct {
		name     string
		fields   fields
		args     args
		wantName string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		n := &Name{
			Name: bb.fields.Name,
		}
		if gotName := n.Name3(tt.args.nn); gotName != bb.wantName {
			b.Errorf("%q. Name.Name3() = %v, want %v", tt.name, gotName, bb.wantName)
		}
	}
}
