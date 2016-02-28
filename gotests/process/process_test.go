package process

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		args []string
		opts *Options
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Nil options and nil args",
			args: nil,
			opts: nil,
			want: "Please specify either the -only, -excl, -export, or -all flag\n",
		}, {
			name: "Nil options",
			args: []string{"testdata/foobar.go"},
			opts: nil,
			want: "Please specify either the -only, -excl, -export, or -all flag\n",
		}, {
			name: "Empty options",
			args: []string{"testdata/foobar.go"},
			opts: &Options{},
			want: "Please specify either the -only, -excl, -export, or -all flag\n",
		}, {
			name: "Non-empty options with no args",
			args: []string{},
			opts: &Options{AllFuncs: true},
			want: "Please specify a file or directory containing the source\n",
		}, {
			name: "AllFuncs option",
			args: []string{"testdata/foobar.go"},
			opts: &Options{AllFuncs: true},
			want: `Generated TestFooFoo
Generated TestBarBar
package foobar

import "testing"

func TestFooFoo(t *testing.T) {
	tests := []struct {
		name    string
		bar     string
		s       string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		f := &Foo{
			Bar: tt.bar,
		}
		if err := f.Foo(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo.Foo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestBarBar(t *testing.T) {
	tests := []struct {
		name    string
		foo     string
		s       string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{
			Foo: tt.foo,
		}
		if err := b.bar(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.bar() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name: "AllFuncs and PrintInputs option",
			args: []string{"testdata/foobar.go"},
			opts: &Options{AllFuncs: true, PrintInputs: true},
			want: `Generated TestFooFoo
Generated TestBarBar
package foobar

import "testing"

func TestFooFoo(t *testing.T) {
	tests := []struct {
		bar     string
		s       string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		f := &Foo{
			Bar: tt.bar,
		}
		if err := f.Foo(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("Foo.Foo(%v) error = %v, wantErr %v", tt.s, err, tt.wantErr)
		}
	}
}

func TestBarBar(t *testing.T) {
	tests := []struct {
		foo     string
		s       string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{
			Foo: tt.foo,
		}
		if err := b.bar(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("Bar.bar(%v) error = %v, wantErr %v", tt.s, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name: "OnlyFuncs option",
			args: []string{"testdata/foobar.go"},
			opts: &Options{OnlyFuncs: "Foo"},
			want: `Generated TestFooFoo
package foobar

import "testing"

func TestFooFoo(t *testing.T) {
	tests := []struct {
		name    string
		bar     string
		s       string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		f := &Foo{
			Bar: tt.bar,
		}
		if err := f.Foo(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo.Foo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name: "OnlyFuncs option w/ no matches",
			args: []string{"testdata/foobar.go"},
			opts: &Options{OnlyFuncs: "FooBar"},
			want: "No tests generated for testdata/foobar.go\n",
		}, {
			name: "Invalid OnlyFuncs option",
			args: []string{"testdata/foobar.go"},
			opts: &Options{OnlyFuncs: "??"},
			want: "Invalid -only regex: error parsing regexp: missing argument to repetition operator: `??`\n",
		}, {
			name: "ExclFuncs option",
			args: []string{"testdata/foobar.go"},
			opts: &Options{ExclFuncs: "Foo"},
			want: `Generated TestBarBar
package foobar

import "testing"

func TestBarBar(t *testing.T) {
	tests := []struct {
		name    string
		foo     string
		s       string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{
			Foo: tt.foo,
		}
		if err := b.bar(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.bar() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		}, {
			name: "Invalid ExclFuncs option",
			args: []string{"testdata/foobar.go"},
			opts: &Options{ExclFuncs: "??"},
			want: "Invalid -excl regex: error parsing regexp: missing argument to repetition operator: `??`\n",
		}, {
			name: "Exported option",
			args: []string{"testdata/foobar.go"},
			opts: &Options{ExportedFuncs: true},
			want: `Generated TestFooFoo
package foobar

import "testing"

func TestFooFoo(t *testing.T) {
	tests := []struct {
		name    string
		bar     string
		s       string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		f := &Foo{
			Bar: tt.bar,
		}
		if err := f.Foo(tt.s); (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo.Foo() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
`,
		},
	}
	for _, tt := range tests {
		out := &bytes.Buffer{}
		Run(out, tt.args, tt.opts)
		if got := out.String(); got != tt.want {
			t.Errorf("%q. Run() =\n%v, want %v", tt.name, got, tt.want)
		}
	}
}
