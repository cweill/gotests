package main

import (
	"io/ioutil"
	"testing"
)

func TestGenerateTestCases(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "Minimal function",
			in:   `testfiles/test1.go`,
			want: `package test1

import (
	"testing"
)

func TestFoo1(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		Foo1()
	}
}
`,
		}, {
			name: "Function w/ anonymous argument",
			in:   `testfiles/test2.go`,
			want: `package test2

import (
	"testing"
)

func TestFoo2(t *testing.T) {
	tests := []struct {
		name string
		in0  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		Foo2(tt.in0)
	}
}
`,
		}, {
			name: "Function w/ named argument",
			in:   `testfiles/test3.go`,
			want: `package test3

import (
	"testing"
)

func TestFoo3(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		Foo3(tt.s)
	}
}
`,
		}, {
			name: "Function w/ return value",
			in:   `testfiles/test4.go`,
			want: `package test4

import (
	"testing"
)

func TestFoo4(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got := Foo4()
		if got != tt.want {
			t.Errorf("%v. Foo4() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function returning an error",
			in:   `testfiles/test5.go`,
			want: `package test5

import (
	"testing"
)

func TestFoo5(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo5()
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo5() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%v. Foo5() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ multiple arguments",
			in:   `testfiles/test6.go`,
			want: `package test6

import (
	"testing"
)

func TestFoo6(t *testing.T) {
	tests := []struct {
		name    string
		i       int
		b       bool
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo6(tt.i, tt.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo6() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%v. Foo6() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Method on a struct pointer",
			in:   `testfiles/test7.go`,
			want: `package test7

import (
	"testing"
)

func TestFoo7(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.b.Foo7()
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo7() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%v. Foo7() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ struct pointer argument and return type",
			in:   `testfiles/test8.go`,
			want: `package test8

import (
	"testing"
)

func TestFoo8(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		want    *Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo8(tt.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo8() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo8() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Struct value method w/ struct value return type",
			in:   `testfiles/test9.go`,
			want: `package test9

import (
	"testing"
)

func TestFoo9(t *testing.T) {
	tests := []struct {
		name string
		b    Bar
		want Bar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got := tt.b.Foo9()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo9() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ map argument and return type",
			in:   `testfiles/test10.go`,
			want: `package test10

import (
	"testing"
)

func TestFoo10(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int32
		want map[string]*Bar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got := Foo10(tt.m)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo10() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ slice argument and return type",
			in:   `testfiles/test11.go`,
			want: `package test11

import (
	"testing"
)

func TestFoo11(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo11(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo11() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo11() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ slice argument and return type",
			in:   `testfiles/test12.go`,
			want: `package test12

import (
	"testing"
)

func TestFoo12(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo12(tt.str); (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo12() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		},
	}
	for _, tt := range tests {
		f, err := ioutil.TempFile("", "")
		if err != nil {
			t.Errorf("%v. Creating temp file: %v", tt.name, err)
			continue
		}
		defer f.Close()
		generateTestCases(f, tt.in)
		b, err := ioutil.ReadFile(f.Name())
		if err != nil {
			t.Errorf("%v. Reading temp file: %v", tt.name, err)
			continue
		}
		if got := string(b); got != tt.want {
			t.Errorf("%v. TestCases(%v) = %v, want %v", tt.name, tt.in, got, tt.want)
		}
	}
}
