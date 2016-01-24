package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateTestCases(t *testing.T) {
	tests := []struct {
		name      string
		srcPath   string
		testPath  string
		onlyFuncs []string
		exclFuncs []string
		want      string
		wantErr   bool
	}{
		{
			name:    "No funcs",
			srcPath: `testfiles/test000.go`,
			wantErr: true,
		}, {
			name:    "Function w/ neither receiver, parameters, nor results",
			srcPath: `testfiles/test001.go`,
			wantErr: true,
		}, {
			name:    "Function w/ anonymous argument",
			srcPath: `testfiles/test002.go`,
			want: `package testfiles

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
			name:    "Function w/ named argument",
			srcPath: `testfiles/test003.go`,
			want: `package testfiles

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
			name:    "Function w/ return value",
			srcPath: `testfiles/test004.go`,
			want: `package testfiles

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
			name:    "Function returning an error",
			srcPath: `testfiles/test005.go`,
			want: `package testfiles

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
			name:    "Function w/ multiple arguments",
			srcPath: `testfiles/test006.go`,
			want: `package testfiles

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
			name:    "Method on a struct pointer",
			srcPath: `testfiles/test007.go`,
			want: `package testfiles

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
			name:    "Function w/ struct pointer argument and return type",
			srcPath: `testfiles/test008.go`,
			want: `package testfiles

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
			name:    "Struct value method w/ struct value return type",
			srcPath: `testfiles/test009.go`,
			want: `package testfiles

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
			name:    "Function w/ map argument and return type",
			srcPath: `testfiles/test010.go`,
			want: `package testfiles

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
			name:    "Function w/ slice argument and return type",
			srcPath: `testfiles/test011.go`,
			want: `package testfiles

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
			name:    "Function returning only an error",
			srcPath: `testfiles/test012.go`,
			want: `package testfiles

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
			name:    "Function w/ a function parameter",
			srcPath: `testfiles/test013.go`,
			want: `package testfiles

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
			name:    "Function w/ a function parameter w/ its own parameters and result",
			srcPath: `testfiles/test014.go`,
			want: `package testfiles

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
			name:    "Function w/ a function parameter that returns two results",
			srcPath: `testfiles/test015.go`,
			want: `package testfiles

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
			name:    "Function w/ interface parameter and result",
			srcPath: `testfiles/test016.go`,
			want: `package testfiles

import (
	"reflect"
	"testing"
)

func TestFoo16(t *testing.T) {
	tests := []struct {
		name string
		in   Bazzar
		want Bazzar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo16(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo16() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ imported interface receiver, parameter, and result",
			srcPath: `testfiles/test017.go`,
			want: `package testfiles

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
			name:    "Function w/ imported struct receiver, parameter, and result",
			srcPath: `testfiles/test018.go`,
			want: `package testfiles

import (
	"os"
	"reflect"
	"testing"
)

func TestFoo18(t *testing.T) {
	tests := []struct {
		name string
		t    *os.File
		want *os.File
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
			name:    "Function w/ multiple parameters of the same type",
			srcPath: `testfiles/test019.go`,
			want: `package testfiles

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
			name:    "Function w/ a variadic parameter",
			srcPath: `testfiles/test020.go`,
			want: `package testfiles

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
		}, {
			name:    "Function w/ interface{} parameter and result",
			srcPath: `testfiles/test021.go`,
			want: `package testfiles

import (
	"reflect"
	"testing"
)

func TestFoo21(t *testing.T) {
	tests := []struct {
		name string
		i    interface{}
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo21(tt.i); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo21() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ named imports",
			srcPath: `testfiles/test022.go`,
			want: `package testfiles

import (
	ht "html/template"
	"reflect"
	"testing"
)

func TestFoo22(t *testing.T) {
	tests := []struct {
		name string
		t    *ht.Template
		want *ht.Template
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo22(tt.t); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo22() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ channel parameter and result",
			srcPath: `testfiles/test023.go`,
			want: `package testfiles

import (
	"reflect"
	"testing"
)

func TestFoo23(t *testing.T) {
	tests := []struct {
		name string
		ch   chan bool
		want chan string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo23(tt.ch); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo23() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "File with multiple imports",
			srcPath: `testfiles/test024.go`,
			want: `package testfiles

import (
	"go/ast"
	"go/types"
	"io"
	"testing"
)

func TestFoo24(t *testing.T) {
	tests := []struct {
		name    string
		w       io.Writer
		x       ast.Expr
		t       types.Type
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo24(tt.w, tt.x, tt.t); (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo24() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:    "Function returning two results and an error",
			srcPath: `testfiles/test025.go`,
			want: `package testfiles

import (
	"reflect"
	"testing"
)

func TestFoo25(t *testing.T) {
	tests := []struct {
		name    string
		in0     interface{}
		want0   string
		want1   []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got0, got1, err := Foo25(tt.in0)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo25() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if got0 != tt.want0 {
			t.Errorf("%v. Foo25() = %v, want %v", tt.name, got0, tt.want0)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%v. Foo25() = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}
`,
		}, {
			name:    "Multiple named results",
			srcPath: `testfiles/test026.go`,
			want: `package testfiles

import (
	"reflect"
	"testing"
)

func TestFoo26(t *testing.T) {
	tests := []struct {
		name    string
		v       interface{}
		want0   string
		want1   int
		want2   []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got0, got1, got2, err := Foo26(tt.v)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo26() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if got0 != tt.want0 {
			t.Errorf("%v. Foo26() = %v, want %v", tt.name, got0, tt.want0)
		}
		if got1 != tt.want1 {
			t.Errorf("%v. Foo26() = %v, want %v", tt.name, got1, tt.want1)
		}
		if !reflect.DeepEqual(got2, tt.want2) {
			t.Errorf("%v. Foo26() = %v, want %v", tt.name, got2, tt.want2)
		}
	}
}
`,
		}, {
			name:    "Multiple functions",
			srcPath: `testfiles/test100.go`,
			want: `package testfiles

import (
	"reflect"
	"testing"
)

func TestFoo100(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo100(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo100() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBar100(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.b.Bar100(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%v. Bar100() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestBaz100(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := baz100(tt.f); got != tt.want {
			t.Errorf("%v. baz100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:      "Multiple functions w/ onlyFuncs",
			srcPath:   `testfiles/test100.go`,
			onlyFuncs: []string{"Foo100", "baz100"},
			want: `package testfiles

import (
	"reflect"
	"testing"
)

func TestFoo100(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo100(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo100() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBaz100(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := baz100(tt.f); got != tt.want {
			t.Errorf("%v. baz100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:      "Multiple functions w/ onlyFuncs by test name",
			srcPath:   `testfiles/test100.go`,
			onlyFuncs: []string{"TestFoo100", "TestBaz100"},
			want: `package testfiles

import (
	"reflect"
	"testing"
)

func TestFoo100(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo100(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo100() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBaz100(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := baz100(tt.f); got != tt.want {
			t.Errorf("%v. baz100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:      "Multiple functions filtering all out",
			srcPath:   `testfiles/test100.go`,
			onlyFuncs: []string{"foo100"},
			wantErr:   true,
		}, {
			name:      "Multiple functions w/ exclFunc",
			srcPath:   `testfiles/test100.go`,
			exclFuncs: []string{"Foo100", "baz100"},
			want: `package testfiles

import "testing"

func TestBar100(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.b.Bar100(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%v. Bar100() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:      "Multiple functions excluding all",
			srcPath:   `testfiles/test100.go`,
			exclFuncs: []string{"baz100", "Foo100", "Bar100"},
			wantErr:   true,
		}, {
			name:      "Multiple functions excluding all test names",
			srcPath:   `testfiles/test100.go`,
			exclFuncs: []string{"TestBaz100", "TestFoo100", "TestBar100"},
			wantErr:   true,
		}, {
			name:      "Multiple functions w/ both onlyFuncs and exclFunc",
			srcPath:   `testfiles/test100.go`,
			onlyFuncs: []string{"Bar100"},
			exclFuncs: []string{"Foo100"},
			want: `package testfiles

import "testing"

func TestBar100(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.b.Bar100(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%v. Bar100() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:      "Multiple functions w/ onlyFuncs and exclFunc competing",
			srcPath:   `testfiles/test100.go`,
			onlyFuncs: []string{"Foo100", "Bar100"},
			exclFuncs: []string{"Foo100", "baz100"},
			want: `package testfiles

import "testing"

func TestBar100(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.b.Bar100(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%v. Bar100() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:     "Existing test file",
			srcPath:  `testfiles/test100.go`,
			testPath: `testfiles/test100_test.go`,
			want: `package testfiles

import (
	"reflect"
	"testing"
)

func TestBar100(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		i       interface{}
		wantErr bool
	}{
		{
			name:    "Basic test",
			b:       &Bar{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		if err := tt.b.Bar100(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%v. Bar100() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestBaz100(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
		{
			name: "Basic test",
			f:    func() *float64 { var x float64 = 64; return &x }(),
			want: 64,
		},
	}
	// TestBaz100 contains a comment.
	for _, tt := range tests {
		if got := baz100(tt.f); got != tt.want {
			t.Errorf("%v. baz100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFoo100(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo100(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo100() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:     "Existing test file with multiple imports",
			srcPath:  `testfiles/test200.go`,
			testPath: `testfiles/test200_test.go`,
			want: `package testfiles

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestFoo200(t *testing.T) {
	tests := []struct {
		name string
		x    ast.Expr
		t    types.Type
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo200(tt.x, tt.t); got != tt.want {
			t.Errorf("%v. Foo200() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBar200(t *testing.T) {
	tests := []struct {
		name string
		t    types.Type
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Bar200(tt.t); got != tt.want {
			t.Errorf("%v. Bar200() = %v, want %v", tt.name, got, tt.want)
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
		f.Close()
		os.Remove(f.Name())
		_, b, err := generateTests(tt.srcPath, tt.testPath, f.Name(), tt.onlyFuncs, tt.exclFuncs, true)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. generateTests() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if got := string(b); got != tt.want {
			t.Errorf("%v. TestCases(%v) = %v, want %v", tt.name, tt.srcPath, got, tt.want)
		}
		b, err = ioutil.ReadFile(f.Name())
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. ioutil.ReadFile: %v, wantErr: %v", tt.name, err, tt.wantErr)
		}
	}
}
