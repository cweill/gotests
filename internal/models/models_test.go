package models

import (
	"testing"
)

// Test Expression.String method
func TestExpression_String(t *testing.T) {
	tests := []struct {
		name string
		expr *Expression
		want string
	}{
		{
			name: "simple value",
			expr: &Expression{Value: "int"},
			want: "int",
		},
		{
			name: "pointer type",
			expr: &Expression{Value: "string", IsStar: true},
			want: "*string",
		},
		{
			name: "variadic type",
			expr: &Expression{Value: "int", IsVariadic: true},
			want: "[]int",
		},
		{
			name: "variadic pointer (variadic takes precedence)",
			expr: &Expression{Value: "string", IsStar: true, IsVariadic: true},
			want: "[]*string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.expr.String()
			if got != tt.want {
				t.Errorf("Expression.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test Field.IsWriter method
func TestField_IsWriter(t *testing.T) {
	tests := []struct {
		name  string
		field *Field
		want  bool
	}{
		{
			name: "io.Writer field",
			field: &Field{
				Type: &Expression{IsWriter: true},
			},
			want: true,
		},
		{
			name: "non-writer field",
			field: &Field{
				Type: &Expression{IsWriter: false},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.IsWriter()
			if got != tt.want {
				t.Errorf("Field.IsWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test Field.IsStruct method
func TestField_IsStruct(t *testing.T) {
	tests := []struct {
		name  string
		field *Field
		want  bool
	}{
		{
			name: "struct field",
			field: &Field{
				Type: &Expression{Underlying: "struct{Name string}"},
			},
			want: true,
		},
		{
			name: "non-struct field",
			field: &Field{
				Type: &Expression{Underlying: "int"},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.IsStruct()
			if got != tt.want {
				t.Errorf("Field.IsStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test Field.IsBasicType method
func TestField_IsBasicType(t *testing.T) {
	tests := []struct {
		name  string
		field *Field
		want  bool
	}{
		{
			name: "int field",
			field: &Field{
				Type: &Expression{Value: "int"},
			},
			want: true,
		},
		{
			name: "string field",
			field: &Field{
				Type: &Expression{Value: "string"},
			},
			want: true,
		},
		{
			name: "bool field",
			field: &Field{
				Type: &Expression{Value: "bool"},
			},
			want: true,
		},
		{
			name: "struct field",
			field: &Field{
				Type: &Expression{Value: "MyStruct"},
			},
			want: false,
		},
		{
			name: "basic underlying type",
			field: &Field{
				Type: &Expression{Value: "MyInt", Underlying: "int"},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.IsBasicType()
			if got != tt.want {
				t.Errorf("Field.IsBasicType() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test Field.IsNamed method
func TestField_IsNamed(t *testing.T) {
	tests := []struct {
		name  string
		field *Field
		want  bool
	}{
		{
			name: "named field",
			field: &Field{
				Name: "myField",
			},
			want: true,
		},
		{
			name: "unnamed field (empty)",
			field: &Field{
				Name: "",
			},
			want: false,
		},
		{
			name: "unnamed field (underscore)",
			field: &Field{
				Name: "_",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.IsNamed()
			if got != tt.want {
				t.Errorf("Field.IsNamed() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test Field.ShortName method
func TestField_ShortName(t *testing.T) {
	tests := []struct {
		name  string
		field *Field
		want  string
	}{
		{
			name: "String type",
			field: &Field{
				Type: &Expression{Value: "String"},
			},
			want: "s",
		},
		{
			name: "Int type",
			field: &Field{
				Type: &Expression{Value: "Int"},
			},
			want: "i",
		},
		{
			name: "MyStruct type",
			field: &Field{
				Type: &Expression{Value: "MyStruct"},
			},
			want: "m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.ShortName()
			if got != tt.want {
				t.Errorf("Field.ShortName() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test Function.TestParameters method
func TestFunction_TestParameters(t *testing.T) {
	tests := []struct {
		name string
		fn   *Function
		want int
	}{
		{
			name: "function with no parameters",
			fn: &Function{
				Parameters: []*Field{},
			},
			want: 0,
		},
		{
			name: "function with regular parameters",
			fn: &Function{
				Parameters: []*Field{
					{Name: "x", Type: &Expression{Value: "int"}},
					{Name: "y", Type: &Expression{Value: "string"}},
				},
			},
			want: 2,
		},
		{
			name: "function with io.Writer parameter (excluded)",
			fn: &Function{
				Parameters: []*Field{
					{Name: "x", Type: &Expression{Value: "int"}},
					{Name: "w", Type: &Expression{Value: "io.Writer", IsWriter: true}},
				},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn.TestParameters()
			if len(got) != tt.want {
				t.Errorf("Function.TestParameters() returned %d parameters, want %d", len(got), tt.want)
			}
		})
	}
}

// Test Function.TestResults method
func TestFunction_TestResults(t *testing.T) {
	tests := []struct {
		name string
		fn   *Function
		want int
	}{
		{
			name: "function with no results",
			fn: &Function{
				Results:    []*Field{},
				Parameters: []*Field{},
			},
			want: 0,
		},
		{
			name: "function with regular results",
			fn: &Function{
				Results: []*Field{
					{Name: "result", Type: &Expression{Value: "int"}},
				},
				Parameters: []*Field{},
			},
			want: 1,
		},
		{
			name: "function with io.Writer parameter (becomes string result)",
			fn: &Function{
				Results: []*Field{},
				Parameters: []*Field{
					{Name: "w", Type: &Expression{Value: "io.Writer", IsWriter: true}},
				},
			},
			want: 1,
		},
		{
			name: "function with both regular results and io.Writer",
			fn: &Function{
				Results: []*Field{
					{Name: "count", Type: &Expression{Value: "int"}},
				},
				Parameters: []*Field{
					{Name: "w", Type: &Expression{Value: "io.Writer", IsWriter: true}},
				},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn.TestResults()
			if len(got) != tt.want {
				t.Errorf("Function.TestResults() returned %d results, want %d", len(got), tt.want)
			}
		})
	}
}

// Test Function.ReturnsMultiple method
func TestFunction_ReturnsMultiple(t *testing.T) {
	tests := []struct {
		name string
		fn   *Function
		want bool
	}{
		{
			name: "no results",
			fn: &Function{
				Results: []*Field{},
			},
			want: false,
		},
		{
			name: "single result",
			fn: &Function{
				Results: []*Field{
					{Type: &Expression{Value: "int"}},
				},
			},
			want: false,
		},
		{
			name: "multiple results",
			fn: &Function{
				Results: []*Field{
					{Type: &Expression{Value: "int"}},
					{Type: &Expression{Value: "string"}},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn.ReturnsMultiple()
			if got != tt.want {
				t.Errorf("Function.ReturnsMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test Function.OnlyReturnsOneValue method
func TestFunction_OnlyReturnsOneValue(t *testing.T) {
	tests := []struct {
		name string
		fn   *Function
		want bool
	}{
		{
			name: "one value, no error",
			fn: &Function{
				Results:      []*Field{{Type: &Expression{Value: "int"}}},
				ReturnsError: false,
			},
			want: true,
		},
		{
			name: "one value, with error",
			fn: &Function{
				Results:      []*Field{{Type: &Expression{Value: "int"}}},
				ReturnsError: true,
			},
			want: false,
		},
		{
			name: "multiple values",
			fn: &Function{
				Results: []*Field{
					{Type: &Expression{Value: "int"}},
					{Type: &Expression{Value: "string"}},
				},
				ReturnsError: false,
			},
			want: false,
		},
		{
			name: "no values",
			fn: &Function{
				Results:      []*Field{},
				ReturnsError: false,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn.OnlyReturnsOneValue()
			if got != tt.want {
				t.Errorf("Function.OnlyReturnsOneValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test Function.OnlyReturnsError method
func TestFunction_OnlyReturnsError(t *testing.T) {
	tests := []struct {
		name string
		fn   *Function
		want bool
	}{
		{
			name: "only error",
			fn: &Function{
				Results:      []*Field{},
				ReturnsError: true,
			},
			want: true,
		},
		{
			name: "value and error",
			fn: &Function{
				Results:      []*Field{{Type: &Expression{Value: "int"}}},
				ReturnsError: true,
			},
			want: false,
		},
		{
			name: "no error",
			fn: &Function{
				Results:      []*Field{},
				ReturnsError: false,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn.OnlyReturnsError()
			if got != tt.want {
				t.Errorf("Function.OnlyReturnsError() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test Function.FullName method
func TestFunction_FullName(t *testing.T) {
	tests := []struct {
		name string
		fn   *Function
		want string
	}{
		{
			name: "function without receiver",
			fn: &Function{
				Name:     "myFunc",
				Receiver: nil,
			},
			want: "MyFunc",
		},
		{
			name: "method with receiver",
			fn: &Function{
				Name: "process",
				Receiver: &Receiver{
					Field: &Field{
						Type: &Expression{Value: "Handler"},
					},
				},
			},
			want: "HandlerProcess",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn.FullName()
			if got != tt.want {
				t.Errorf("Function.FullName() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test Function.TestName method
func TestFunction_TestName(t *testing.T) {
	tests := []struct {
		name string
		fn   *Function
		want string
	}{
		{
			name: "exported function",
			fn: &Function{
				Name: "MyFunc",
			},
			want: "TestMyFunc",
		},
		{
			name: "unexported function",
			fn: &Function{
				Name: "myFunc",
			},
			want: "Test_myFunc",
		},
		{
			name: "already starts with Test",
			fn: &Function{
				Name: "TestSomething",
			},
			want: "TestSomething",
		},
		{
			name: "method on exported receiver",
			fn: &Function{
				Name: "Process",
				Receiver: &Receiver{
					Field: &Field{
						Type: &Expression{Value: "Handler"},
					},
				},
			},
			want: "TestHandler_Process",
		},
		{
			name: "method on unexported receiver",
			fn: &Function{
				Name: "process",
				Receiver: &Receiver{
					Field: &Field{
						Type: &Expression{Value: "handler"},
					},
				},
			},
			want: "Test_handler_process",
		},
		{
			name: "generic receiver with type params",
			fn: &Function{
				Name: "Add",
				Receiver: &Receiver{
					Field: &Field{
						Type: &Expression{Value: "Set[T]"},
					},
				},
			},
			want: "TestSet_Add",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn.TestName()
			if got != tt.want {
				t.Errorf("Function.TestName() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test Function.IsNaked method
func TestFunction_IsNaked(t *testing.T) {
	tests := []struct {
		name string
		fn   *Function
		want bool
	}{
		{
			name: "naked function (no receiver, params, or results)",
			fn: &Function{
				Receiver:   nil,
				Parameters: []*Field{},
				Results:    []*Field{},
			},
			want: true,
		},
		{
			name: "function with parameters",
			fn: &Function{
				Receiver:   nil,
				Parameters: []*Field{{Type: &Expression{Value: "int"}}},
				Results:    []*Field{},
			},
			want: false,
		},
		{
			name: "function with results",
			fn: &Function{
				Receiver:   nil,
				Parameters: []*Field{},
				Results:    []*Field{{Type: &Expression{Value: "int"}}},
			},
			want: false,
		},
		{
			name: "method (has receiver)",
			fn: &Function{
				Receiver: &Receiver{
					Field: &Field{Type: &Expression{Value: "MyStruct"}},
				},
				Parameters: []*Field{},
				Results:    []*Field{},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn.IsNaked()
			if got != tt.want {
				t.Errorf("Function.IsNaked() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test Path.TestPath method
func TestPath_TestPath(t *testing.T) {
	tests := []struct {
		name string
		path Path
		want string
	}{
		{
			name: "source file",
			path: Path("myfile.go"),
			want: "myfile_test.go",
		},
		{
			name: "source file with path",
			path: Path("/path/to/myfile.go"),
			want: "/path/to/myfile_test.go",
		},
		{
			name: "already test file",
			path: Path("myfile_test.go"),
			want: "myfile_test.go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.path.TestPath()
			if got != tt.want {
				t.Errorf("Path.TestPath() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test Path.IsTestPath method
func TestPath_IsTestPath(t *testing.T) {
	tests := []struct {
		name string
		path Path
		want bool
	}{
		{
			name: "test file",
			path: Path("myfile_test.go"),
			want: true,
		},
		{
			name: "source file",
			path: Path("myfile.go"),
			want: false,
		},
		{
			name: "test file with path",
			path: Path("/path/to/myfile_test.go"),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.path.IsTestPath()
			if got != tt.want {
				t.Errorf("Path.IsTestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
