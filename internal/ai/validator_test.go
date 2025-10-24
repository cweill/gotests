package ai

import (
	"strings"
	"testing"
)

func TestValidateGeneratedTest(t *testing.T) {
	tests := []struct {
		name     string
		testCode string
		pkgName  string
		wantErr  bool
		errMsg   string
	}{
		{
			name: "valid_test_code",
			testCode: `package testpkg

import "testing"

func TestFoo(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "test1", want: 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := 42; got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
`,
			pkgName: "testpkg",
			wantErr: false,
		},
		{
			name: "syntax_error_missing_brace",
			testCode: `package testpkg

import "testing"

func TestBar(t *testing.T) {
	// Missing closing brace
`,
			pkgName: "testpkg",
			wantErr: true,
			errMsg:  "parse error",
		},
		{
			name: "syntax_error_invalid_token",
			testCode: `package testpkg

import "testing"

func TestBaz(t *testing.T) {
	var x int = "string" // Type mismatch
}
`,
			pkgName: "testpkg",
			wantErr: true,
			errMsg:  "type check error",
		},
		{
			name: "missing_package_declaration",
			testCode: `import "testing"

func TestMissingPkg(t *testing.T) {
}
`,
			pkgName: "testpkg",
			wantErr: true,
			errMsg:  "parse error",
		},
		{
			name: "undefined_variable",
			testCode: `package testpkg

import "testing"

func TestUndefined(t *testing.T) {
	x := undefinedVariable + 1
}
`,
			pkgName: "testpkg",
			wantErr: true,
			errMsg:  "type check error",
		},
		{
			name: "valid_with_imports",
			testCode: `package testpkg

import (
	"testing"
	"fmt"
	"strings"
)

func TestWithImports(t *testing.T) {
	s := strings.ToUpper("hello")
	fmt.Println(s)
}
`,
			pkgName: "testpkg",
			wantErr: false,
		},
		{
			name: "valid_table_driven_test",
			testCode: `package testpkg

import "testing"

type args struct {
	a int
	b int
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive", args: args{a: 2, b: 3}, want: 5},
		{name: "negative", args: args{a: -1, b: -1}, want: -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.a + tt.args.b
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
`,
			pkgName: "testpkg",
			wantErr: false,
		},
		{
			name: "type_error_wrong_operation",
			testCode: `package testpkg

import "testing"

func TestTypeError(t *testing.T) {
	var s string = "hello"
	var x int = s + 10  // Can't add string and int
}
`,
			pkgName: "testpkg",
			wantErr: true,
			errMsg:  "type check error",
		},
		{
			name:     "empty_test_code",
			testCode: ``,
			pkgName:  "testpkg",
			wantErr:  true,
			errMsg:   "parse error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGeneratedTest(tt.testCode, tt.pkgName)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGeneratedTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.errMsg != "" {
				if err == nil {
					t.Errorf("ValidateGeneratedTest() expected error containing %q, got nil", tt.errMsg)
				} else if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("ValidateGeneratedTest() error = %v, want error containing %q", err, tt.errMsg)
				}
			}
		})
	}
}

func TestValidateGeneratedTest_ComplexTypes(t *testing.T) {
	tests := []struct {
		name     string
		testCode string
		pkgName  string
		wantErr  bool
	}{
		{
			name: "with_struct_types",
			testCode: `package testpkg

import "testing"

type User struct {
	Name string
	Age  int
}

func TestUser(t *testing.T) {
	u := User{Name: "Alice", Age: 30}
	if u.Name != "Alice" {
		t.Errorf("Name = %v, want Alice", u.Name)
	}
}
`,
			pkgName: "testpkg",
			wantErr: false,
		},
		{
			name: "with_slice_types",
			testCode: `package testpkg

import "testing"

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3}
	if len(s) != 3 {
		t.Errorf("len = %v, want 3", len(s))
	}
}
`,
			pkgName: "testpkg",
			wantErr: false,
		},
		{
			name: "with_map_types",
			testCode: `package testpkg

import "testing"

func TestMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	if m["a"] != 1 {
		t.Errorf("m[a] = %v, want 1", m["a"])
	}
}
`,
			pkgName: "testpkg",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGeneratedTest(tt.testCode, tt.pkgName)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGeneratedTest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
