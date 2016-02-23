package gotests

import (
	"errors"
	"go/types"
	"regexp"
	"testing"
)

func TestGenerateTests(t *testing.T) {
	tests := []struct {
		name              string
		srcPath           string
		only              *regexp.Regexp
		excl              *regexp.Regexp
		exported          bool
		printInputs       bool
		importer          types.Importer
		want              string
		wantNoTests       bool
		wantMultipleTests bool
		wantErr           bool
	}{
		{
			name:        "Hidden file",
			srcPath:     `testdata/.hidden.go`,
			wantNoTests: true,
			wantErr:     true,
		}, {
			name:        "No funcs",
			srcPath:     `testdata/test000.go`,
			wantNoTests: true,
		}, {
			name:        "Function w/ neither receiver, parameters, nor results",
			srcPath:     `testdata/test001.go`,
			wantNoTests: true,
		}, {
			name:    "Function w/ anonymous argument",
			srcPath: `testdata/test002.go`,
			want: `package testdata

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
			srcPath: `testdata/test003.go`,
			want: `package testdata

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
			srcPath: `testdata/test004.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo4() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function returning an error",
			srcPath: `testdata/test005.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo5() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Foo5() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ multiple arguments",
			srcPath: `testdata/test006.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo6() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Foo6() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:        "Print inputs with multiple arguments ",
			srcPath:     `testdata/test006.go`,
			printInputs: true,
			want: `package testdata

import "testing"

func TestFoo6(t *testing.T) {
	tests := []struct {
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
			t.Errorf("Foo6(%v, %v) error = %v, wantErr %v", tt.i, tt.b, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("Foo6(%v, %v) = %v, want %v", tt.i, tt.b, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Method on a struct pointer",
			srcPath: `testdata/test007.go`,
			want: `package testdata

import "testing"

func TestBarFoo7(t *testing.T) {
	tests := []struct {
		name    string
		i       int
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		got, err := b.Foo7(tt.i)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.Foo7() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Bar.Foo7() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:        "Print inputs with single argument",
			srcPath:     `testdata/test007.go`,
			printInputs: true,
			want: `package testdata

import "testing"

func TestBarFoo7(t *testing.T) {
	tests := []struct {
		i       int
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		got, err := b.Foo7(tt.i)
		if (err != nil) != tt.wantErr {
			t.Errorf("Bar.Foo7(%v) error = %v, wantErr %v", tt.i, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("Bar.Foo7(%v) = %v, want %v", tt.i, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ struct pointer argument and return type",
			srcPath: `testdata/test008.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo8() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo8() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Struct value method w/ struct value return type",
			srcPath: `testdata/test009.go`,
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestBarFoo9(t *testing.T) {
	tests := []struct {
		name string
		want Bar
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := Bar{}
		if got := b.Foo9(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Bar.Foo9() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ map argument and return type",
			srcPath: `testdata/test010.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo10() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ slice argument and return type",
			srcPath: `testdata/test011.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo11() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo11() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function returning only an error",
			srcPath: `testdata/test012.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo12() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:    "Function w/ a function parameter",
			srcPath: `testdata/test013.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo13() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:    "Function w/ a function parameter w/ its own parameters and result",
			srcPath: `testdata/test014.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo14() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:    "Function w/ a function parameter that returns two results",
			srcPath: `testdata/test015.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo15() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:    "Function w/ interface parameter and result",
			srcPath: `testdata/test016.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo16() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ imported interface receiver, parameter, and result",
			srcPath: `testdata/test017.go`,
			want: `package testdata

import (
	"io"
	"reflect"
	"testing"
)

func TestFoo17(t *testing.T) {
	tests := []struct {
		name string
		r    io.Reader
		want io.Reader
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo17(tt.r); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo17() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ imported struct receiver, parameter, and result",
			srcPath: `testdata/test018.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo18() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ multiple parameters of the same type",
			srcPath: `testdata/test019.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo19() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ a variadic parameter",
			srcPath: `testdata/test020.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo20() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ interface{} parameter and result",
			srcPath: `testdata/test021.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo21() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ named imports",
			srcPath: `testdata/test022.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo22() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Function w/ channel parameter and result",
			srcPath: `testdata/test023.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo23() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "File with multiple imports",
			srcPath: `testdata/test024.go`,
			want: `package testdata

import (
	"go/ast"
	"go/types"
	"io"
	"testing"
)

func TestFoo24(t *testing.T) {
	tests := []struct {
		name    string
		r       io.Reader
		x       ast.Expr
		t       types.Type
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := Foo24(tt.r, tt.x, tt.t); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo24() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:    "Function returning two results and an error",
			srcPath: `testdata/test025.go`,
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestFoo25(t *testing.T) {
	tests := []struct {
		name    string
		in0     interface{}
		want    string
		want1   []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1, err := Foo25(tt.in0)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo25() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Foo25() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. Foo25() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}
`,
		}, {
			name:    "Multiple named results",
			srcPath: `testdata/test026.go`,
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestFoo26(t *testing.T) {
	tests := []struct {
		name    string
		v       interface{}
		want    string
		want1   int
		want2   []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1, got2, err := Foo26(tt.v)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo26() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Foo26() got = %v, want %v", tt.name, got, tt.want)
		}
		if got1 != tt.want1 {
			t.Errorf("%q. Foo26() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
		if !reflect.DeepEqual(got2, tt.want2) {
			t.Errorf("%q. Foo26() got2 = %v, want %v", tt.name, got2, tt.want2)
		}
	}
}
`,
		}, {
			name:    "Two different structs with same method name",
			srcPath: `testdata/test027.go`,
			want: `package testdata

import "testing"

func TestBookOpen(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Book{}
		if err := b.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. Book.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestDoorOpen(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := &door{}
		if err := d.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. door.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestXmlOpen(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		x := &xml{}
		if err := x.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. xml.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:    "Underlying types",
			srcPath: `testdata/test028.go`,
			want: `package testdata

import (
	"testing"
	"time"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want Fahrenheit
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.c.ToFahrenheit(); got != tt.want {
			t.Errorf("%q. Celsius.ToFahrenheit() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestHourToSecond(t *testing.T) {
	tests := []struct {
		name string
		h    time.Duration
		want time.Duration
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := HourToSecond(tt.h); got != tt.want {
			t.Errorf("%q. HourToSecond() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Struct receiver with multiple fields",
			srcPath: `testdata/test029.go`,
			want: `package testdata

import "testing"

func TestPersonSayHello(t *testing.T) {
	tests := []struct {
		name      string
		firstName string
		lastName  string
		age       int
		gender    string
		siblings  []*Person
		r         *Person
		want      string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		p := &Person{
			FirstName: tt.firstName,
			LastName:  tt.lastName,
			Age:       tt.age,
			Gender:    tt.gender,
			Siblings:  tt.siblings,
		}
		if got := p.SayHello(tt.r); got != tt.want {
			t.Errorf("%q. Person.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Struct receiver with anonymous fields",
			srcPath: `testdata/test030.go`,
			want: `package testdata

import "testing"

func TestDoctorSayHello(t *testing.T) {
	tests := []struct {
		name        string
		person      *Person
		id          string
		numPatients int
		string      string
		r           *Person
		want        string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := &Doctor{
			Person:      tt.person,
			ID:          tt.id,
			numPatients: tt.numPatients,
			string:      tt.string,
		}
		if got := d.SayHello(tt.r); got != tt.want {
			t.Errorf("%q. Doctor.SayHello() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "io.Writer parameters",
			srcPath: `testdata/test031.go`,
			want: `package testdata

import (
	"bytes"
	"testing"
)

func TestBarWrite(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		w := &bytes.Buffer{}
		if err := b.Write(w); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.Write() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got := w.String(); got != tt.want {
			t.Errorf("%q. Bar.Write() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestWrite(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		w := &bytes.Buffer{}
		if err := Write(w, tt.data); (err != nil) != tt.wantErr {
			t.Errorf("%q. Write() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got := w.String(); got != tt.want {
			t.Errorf("%q. Write() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestMultiWrite(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    int
		want1   string
		want2   string
		want3   string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		w1 := &bytes.Buffer{}
		w2 := &bytes.Buffer{}
		got, got1, err := MultiWrite(w1, w2, tt.data)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. MultiWrite() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. MultiWrite() got = %v, want %v", tt.name, got, tt.want)
		}
		if got1 != tt.want1 {
			t.Errorf("%q. MultiWrite() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
		if got2 := w1.String(); got2 != tt.want2 {
			t.Errorf("%q. MultiWrite() got2 = %v, want %v", tt.name, got2, tt.want2)
		}
		if got3 := w2.String(); got3 != tt.want3 {
			t.Errorf("%q. MultiWrite() got3 = %v, want %v", tt.name, got3, tt.want3)
		}
	}
}
`,
		}, {
			name:    "Two structs with same method name",
			srcPath: `testdata/test032.go`,
			want: `package testdata

import "testing"

func TestCelsiusString(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.c.String(); got != tt.want {
			t.Errorf("%q. Celsius.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFahrenheitString(t *testing.T) {
	tests := []struct {
		name string
		f    Fahrenheit
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.f.String(); got != tt.want {
			t.Errorf("%q. Fahrenheit.String() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Functions and methods with 'name' receivers, parameters, and results",
			srcPath: `testdata/test033.go`,
			want: `package testdata

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
		name  string
		fname string
		nn    string
		want  string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		n := &Name{
			Name: tt.fname,
		}
		if got := n.Name3(tt.nn); got != tt.want {
			t.Errorf("%q. Name.Name3() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Multiple functions",
			srcPath: `testdata/test_filter.go`,
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestFooFilter(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		name    string
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestBazFilter(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bazFilter(tt.f); got != tt.want {
			t.Errorf("%q. bazFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Multiple functions w/ only",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("FooFilter|bazFilter"),
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestFooFilter(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBazFilter(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bazFilter(tt.f); got != tt.want {
			t.Errorf("%q. bazFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Multiple functions w/ case-insensitive only",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("(?i)fooFilter|BazFilter"),
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestFooFilter(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBazFilter(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bazFilter(tt.f); got != tt.want {
			t.Errorf("%q. bazFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Multiple functions w/ only filtering on receiver",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("^BarBarFilter$"),
			want: `package testdata

import "testing"

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		name    string
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:    "Multiple functions w/ only filtering on method",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("^(BarFilter)$"),
			want: `package testdata

import "testing"

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		name    string
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:     "Multiple functions filtering exported",
			srcPath:  `testdata/test_filter.go`,
			exported: true,
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestFooFilter(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		name    string
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:     "Multiple functions filtering exported w/ only",
			srcPath:  `testdata/test_filter.go`,
			only:     regexp.MustCompile(`FooFilter`),
			exported: true,
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestFooFilter(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:        "Multiple functions filtering all out",
			srcPath:     `testdata/test_filter.go`,
			only:        regexp.MustCompile("fooFilter"),
			wantNoTests: true,
		}, {
			name:    "Multiple functions w/ excl",
			srcPath: `testdata/test_filter.go`,
			excl:    regexp.MustCompile("FooFilter|bazFilter"),
			want: `package testdata

import "testing"

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		name    string
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:    "Multiple functions w/ case-insensitive excl",
			srcPath: `testdata/test_filter.go`,
			excl:    regexp.MustCompile("(?i)foOFilter|BaZFilter"),
			want: `package testdata

import "testing"

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		name    string
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:     "Multiple functions filtering exported w/ excl",
			srcPath:  `testdata/test_filter.go`,
			excl:     regexp.MustCompile(`FooFilter`),
			exported: true,
			want: `package testdata

import "testing"

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		name    string
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:        "Multiple functions excluding all",
			srcPath:     `testdata/test_filter.go`,
			excl:        regexp.MustCompile("bazFilter|FooFilter|BarFilter"),
			wantNoTests: true,
		}, {
			name:    "Multiple functions excluding on receiver",
			srcPath: `testdata/test_filter.go`,
			excl:    regexp.MustCompile("^BarBarFilter$"),
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestFooFilter(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBazFilter(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bazFilter(tt.f); got != tt.want {
			t.Errorf("%q. bazFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Multiple functions excluding on method",
			srcPath: `testdata/test_filter.go`,
			excl:    regexp.MustCompile("^BarFilter$"),
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestFooFilter(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBazFilter(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bazFilter(tt.f); got != tt.want {
			t.Errorf("%q. bazFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Multiple functions w/ both only and excl",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("BarFilter"),
			excl:    regexp.MustCompile("FooFilter"),
			want: `package testdata

import "testing"

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		name    string
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:    "Multiple functions w/ only and excl competing",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("FooFilter|BarFilter"),
			excl:    regexp.MustCompile("FooFilter|bazFilter"),
			want: `package testdata

import "testing"

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		name    string
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		if err := b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name:     "Custom importer fails",
			srcPath:  `testdata/test_filter.go`,
			importer: &fakeImporter{err: errors.New("error")},
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestFooFilter(t *testing.T) {
	tests := []struct {
		name    string
		strs    []string
		want    []*Bar
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := FooFilter(tt.strs)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FooFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestBarBarFilter(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		i       interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.b.BarFilter(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestBazFilter(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := bazFilter(tt.f); got != tt.want {
			t.Errorf("%q. bazFilter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Existing test file",
			srcPath: `testdata/test100.go`,
			want: `package testdata

import (
	"reflect"
	"testing"
)

func TestBarBar100(t *testing.T) {
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
			t.Errorf("%q. Bar100() error = %v, wantErr %v", tt.name, err, tt.wantErr)
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
			t.Errorf("%q. baz100() = %v, want %v", tt.name, got, tt.want)
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
			t.Errorf("%q. Foo100() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Foo100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Existing test file with just package declaration",
			srcPath: `testdata/test101.go`,
			want: `package testdata

import "testing"

func TestFoo101(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo101(tt.s); got != tt.want {
			t.Errorf("%q. Foo101() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Existing test file with no functions",
			srcPath: `testdata/test102.go`,
			want: `package testdata

import (
	"fmt"
	"testing"
)

var example102 = fmt.Sprintf("test%", 1)

func TestFoo102(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := Foo102(tt.s); got != tt.want {
			t.Errorf("%q. Foo102() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:    "Existing test file with multiple imports",
			srcPath: `testdata/test200.go`,
			want: `package testdata

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
			t.Errorf("%q. Foo200() = %v, want %v", tt.name, got, tt.want)
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
			t.Errorf("%q. Bar200() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
`,
		}, {
			name:              "Entire testdata directory",
			srcPath:           `testdata/`,
			wantMultipleTests: true,
		},
	}
	for _, tt := range tests {
		gts, err := GenerateTests(tt.srcPath, &Options{
			Only:        tt.only,
			Exclude:     tt.excl,
			Exported:    tt.exported,
			PrintInputs: tt.printInputs,
			Importer:    func() types.Importer { return tt.importer },
		})
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. generateTests() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if len(gts) == 0 && !tt.wantNoTests {
			t.Errorf("%q. generateTests() returned no tests", tt.name)
			continue
		}
		if len(gts) > 1 && !tt.wantMultipleTests {
			t.Errorf("%q. generateTests() returned too many tests", tt.name)
			continue
		}
		if tt.wantNoTests || tt.wantMultipleTests {
			continue
		}
		if got := string(gts[0].Output); got != tt.want {
			t.Errorf("%q. TestCases(%v) = \n%v, want \n%v", tt.name, tt.srcPath, got, tt.want)
		}
	}
}

// 249032394 ns/op
func BenchmarkGenerateTests(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateTests("testdata/", &Options{})
	}
}

// A fake importer.
type fakeImporter struct {
	err error
}

func (f *fakeImporter) Import(path string) (*types.Package, error) {
	return nil, f.err
}
