package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateTestCases(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "Unexported function",
			in:   `testfiles/test000.go`,
			want: ``,
		}, {
			name: "Minimal function",
			in:   `testfiles/test001.go`,
			want: `package test1

import "testing"

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
			in:   `testfiles/test002.go`,
			want: `package test2

import "testing"

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
			in:   `testfiles/test003.go`,
			want: `package test3

import "testing"

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
			in:   `testfiles/test004.go`,
			want: `package test4

import "testing"

func TestFoo4(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo4(); got != tt.want {
			t.Errorf("%v. Foo4() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function returning an error",
			in:   `testfiles/test005.go`,
			want: `package test5

import "testing"

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
			in:   `testfiles/test006.go`,
			want: `package test6

import "testing"

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
			in:   `testfiles/test007.go`,
			want: `package test7

import "testing"

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
			in:   `testfiles/test008.go`,
			want: `package test8

import (
	"reflect"
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
			in:   `testfiles/test009.go`,
			want: `package test9

import (
	"reflect"
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
		if got := tt.b.Foo9(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo9() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ map argument and return type",
			in:   `testfiles/test010.go`,
			want: `package test10

import (
	"reflect"
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
		if got := Foo10(tt.m); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo10() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ slice argument and return type",
			in:   `testfiles/test011.go`,
			want: `package test11

import (
	"reflect"
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
			name: "Function returning only an error",
			in:   `testfiles/test012.go`,
			want: `package test12

import "testing"

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
		}, {
			name: "Function w/ a function parameter",
			in:   `testfiles/test013.go`,
			want: `package test13

import "testing"

func TestFoo13(t *testing.T) {
	tests := []struct {
		name    string
		f       func()
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo13(tt.f); (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo13() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name: "Function w/ a function parameter w/ its own parameters and result",
			in:   `testfiles/test014.go`,
			want: `package test14

import "testing"

func TestFoo14(t *testing.T) {
	tests := []struct {
		name    string
		f       func(string, int) string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo14(tt.f); (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo14() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name: "Function w/ a function parameter that returns two results",
			in:   `testfiles/test015.go`,
			want: `package test15

import "testing"

func TestFoo15(t *testing.T) {
	tests := []struct {
		name    string
		f       func(string) (string, error)
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo15(tt.f); (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo15() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name: "Function w/ interface receiver, parameter, and result",
			in:   `testfiles/test016.go`,
			want: `package test16

import (
	"reflect"
	"testing"
)

func TestBaz(t *testing.T) {
	tests := []struct {
		name string
		b    *baz
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.b.Baz(); got != tt.want {
			t.Errorf("%v. Baz() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFoo16(t *testing.T) {
	tests := []struct {
		name string
		i    Bazzar
		in   Bazzar
		want Bazzar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.i.Foo16(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo16() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ imported interface receiver, parameter, and result",
			in:   `testfiles/test017.go`,
			want: `package test17

import (
	"io"
	"reflect"
	"testing"
)

func TestFoo17(t *testing.T) {
	tests := []struct {
		name string
		w    io.Writer
		want io.Writer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo17(tt.w); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo17() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ imported struct receiver, parameter, and result",
			in:   `testfiles/test018.go`,
			want: `package test18

import (
	"reflect"
	"testing"
	"text/template"
)

func TestFoo18(t *testing.T) {
	tests := []struct {
		name string
		t    *template.Template
		want *template.Template
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo18(tt.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo18() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ multiple parameters of the same type",
			in:   `testfiles/test019.go`,
			want: `package test19

import "testing"

func TestFoo19(t *testing.T) {
	tests := []struct {
		name string
		in1  string
		in2  string
		in3  string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo19(tt.in1, tt.in2, tt.in3); got != tt.want {
			t.Errorf("%v. Foo19() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name: "Function w/ a variadic parameter",
			in:   `testfiles/test020.go`,
			want: `package test20

import "testing"

func TestFoo20(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo20(tt.strs...); got != tt.want {
			t.Errorf("%v. Foo20() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		},
	}
	for _, tt := range tests {
		f, err := ioutil.TempFile("", "")
		if err != nil {
			t.Errorf("%v. ioutil.TempFile: %v", tt.name, err)
			continue
		}
		defer os.Remove(f.Name())
		generateTestCases(f, tt.in)
		b, err := ioutil.ReadFile(f.Name())
		if err != nil {
			t.Errorf("%v. ioutil.ReadFile: %v", tt.name, err)
			continue
		}
		if got := string(b); got != tt.want {
			t.Errorf("%v. TestCases(%v) = %v, want %v", tt.name, tt.in, got, tt.want)
		}
	}
}
