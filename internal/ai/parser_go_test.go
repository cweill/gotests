package ai

import (
	"reflect"
	"testing"
)

func Test_extractCodeFromMarkdown(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "with_go_code_block",
			text: "```go\nfunc Test() {}\n```",
			want: "func Test() {}",
		},
		{
			name: "with_plain_code_block",
			text: "```\nfunc Test() {}\n```",
			want: "func Test() {}",
		},
		{
			name: "no_markdown",
			text: "func Test() {}",
			want: "func Test() {}",
		},
		{
			name: "multiline_code_block",
			text: "```go\nfunc Test() {\n    return 42\n}\n```",
			want: "func Test() {\n    return 42\n}",
		},
		{
			name: "code_block_with_extra_text",
			text: "Here is some code:\n```go\nfunc Test() {}\n```\nThat was the code",
			want: "func Test() {}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractCodeFromMarkdown(tt.text)
			if got != tt.want {
				t.Errorf("extractCodeFromMarkdown() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_ensureTrailingComma(t *testing.T) {
	tests := []struct {
		name string
		code string
		want string
	}{
		{
			name: "missing_trailing_comma",
			code: "{\n    name: \"test\"\n}",
			want: "{\n    name: \"test\"\n},",
		},
		{
			name: "already_has_comma",
			code: "{\n    name: \"test\"\n},",
			want: "{\n    name: \"test\"\n},",
		},
		{
			name: "complete_function_not_modified",
			code: "func TestFoo(t *testing.T) {\n    tests := []struct{}{}\n}",
			want: "func TestFoo(t *testing.T) {\n    tests := []struct{}{}\n}",
		},
		{
			name: "empty_string",
			code: "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ensureTrailingComma(tt.code)
			if got != tt.want {
				t.Errorf("ensureTrailingComma() = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_parseGoTestCases(t *testing.T) {
	tests := []struct {
		name     string
		goCode   string
		maxCases int
		want     []TestCase
		wantErr  bool
	}{
		{
			name: "complete_function_simple",
			goCode: `package testdata

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
		{
			name: "positive",
			args: args{a: 5, b: 3},
			want: 8,
		},
	}
}`,
			maxCases: 3,
			want: []TestCase{
				{
					Name: "positive",
					Args: map[string]string{"a": "5", "b": "3"},
					Want: map[string]string{"want": "8"},
				},
			},
			wantErr: false,
		},
		{
			name: "complete_function_with_error",
			goCode: `package testdata

import "testing"

type args struct {
	n int
	d int
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "divide_by_zero",
			args: args{n: 10, d: 0},
			want: 0,
			wantErr: true,
		},
		{
			name: "valid_division",
			args: args{n: 10, d: 2},
			want: 5,
			wantErr: false,
		},
	}
}`,
			maxCases: 3,
			want: []TestCase{
				{
					Name:    "divide_by_zero",
					Args:    map[string]string{"n": "10", "d": "0"},
					Want:    map[string]string{"want": "0"},
					WantErr: true,
				},
				{
					Name:    "valid_division",
					Args:    map[string]string{"n": "10", "d": "2"},
					Want:    map[string]string{"want": "5"},
					WantErr: false,
				},
			},
			wantErr: false,
		},
		{
			name: "test_case_array_only",
			goCode: `{
	name: "test1",
	args: args{x: 10},
	want: 20,
},`,
			maxCases: 3,
			want: []TestCase{
				{
					Name: "test1",
					Args: map[string]string{"x": "10"},
					Want: map[string]string{"want": "20"},
				},
			},
			wantErr: false,
		},
		{
			name: "max_cases_limit",
			goCode: `package testdata

import "testing"

func TestMany(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "test1", want: 1},
		{name: "test2", want: 2},
		{name: "test3", want: 3},
		{name: "test4", want: 4},
	}
}`,
			maxCases: 2,
			want: []TestCase{
				{Name: "test1", Args: map[string]string{}, Want: map[string]string{"want": "1"}},
				{Name: "test2", Args: map[string]string{}, Want: map[string]string{"want": "2"}},
			},
			wantErr: false,
		},
		{
			name:     "empty_input",
			goCode:   "",
			maxCases: 3,
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "invalid_go_code",
			goCode:   "this is not valid go code {{{",
			maxCases: 3,
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseGoTestCases(tt.goCode, tt.maxCases)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseGoTestCases() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGoTestCases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTestCase(t *testing.T) {
	// This is primarily tested through parseGoTestCases,
	// but we can add specific edge cases here
	tests := []struct {
		name    string
		goCode  string
		want    *TestCase
		wantNil bool
	}{
		{
			name: "with_receiver_field",
			goCode: `package testdata

import "testing"

type Calculator struct{}
type args struct{ n int }

func Test(t *testing.T) {
	tests := []struct {
		name string
		c    *Calculator
		args args
		want int
	}{
		{
			name: "test",
			c:    &Calculator{},
			args: args{n: 5},
			want: 10,
		},
	}
}`,
			want: &TestCase{
				Name: "test",
				Args: map[string]string{"n": "5"},
				Want: map[string]string{"want": "10"},
			},
			wantNil: false,
		},
		{
			name: "unnamed_test_case_ignored",
			goCode: `package testdata

import "testing"

type args struct{ n int }

func Test(t *testing.T) {
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{n: 5},
			want: 10,
		},
	}
}`,
			want:    nil,
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cases, err := parseGoTestCases(tt.goCode, 10)
			if err != nil {
				if !tt.wantNil {
					t.Errorf("parseGoTestCases() unexpected error = %v", err)
				}
				return
			}

			if tt.wantNil {
				if len(cases) > 0 {
					t.Errorf("parseGoTestCases() expected no cases, got %d", len(cases))
				}
				return
			}

			if len(cases) == 0 {
				t.Errorf("parseGoTestCases() expected cases, got none")
				return
			}

			got := &cases[0]
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGoTestCases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseArgsStruct(t *testing.T) {
	// Tested through parseGoTestCases, this verifies the Args map is correctly populated
	goCode := `package testdata

import "testing"

type args struct {
	firstName string
	lastName  string
	age       int
}

func Test(t *testing.T) {
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "multiple_args",
			args: args{
				firstName: "John",
				lastName:  "Doe",
				age:       30,
			},
			want: "John Doe (30)",
		},
	}
}`

	cases, err := parseGoTestCases(goCode, 1)
	if err != nil {
		t.Fatalf("parseGoTestCases() error = %v", err)
	}

	if len(cases) != 1 {
		t.Fatalf("parseGoTestCases() returned %d cases, want 1", len(cases))
	}

	got := cases[0].Args
	want := map[string]string{
		"firstName": `"John"`,
		"lastName":  `"Doe"`,
		"age":       "30",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Args = %v, want %v", got, want)
	}
}
