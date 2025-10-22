package render

import (
	"testing"

	"github.com/cweill/gotests/internal/models"
)

// Test typeArguments function with various scenarios
func Test_typeArguments(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
		want string
	}{
		{
			name: "non-generic function",
			data: struct {
				TypeParams []*models.TypeParam
				Receiver   *models.Receiver
			}{
				TypeParams: nil,
				Receiver:   nil,
			},
			want: "",
		},
		{
			name: "generic function with 'any' constraint",
			data: struct {
				TypeParams []*models.TypeParam
				Receiver   *models.Receiver
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "any"},
				},
				Receiver: nil,
			},
			want: "[int]",
		},
		{
			name: "generic function with 'comparable' constraint",
			data: struct {
				TypeParams []*models.TypeParam
				Receiver   *models.Receiver
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "comparable"},
				},
				Receiver: nil,
			},
			want: "[string]",
		},
		{
			name: "generic function with multiple type params",
			data: struct {
				TypeParams []*models.TypeParam
				Receiver   *models.Receiver
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "any"},
					{Name: "U", Constraint: "comparable"},
				},
				Receiver: nil,
			},
			want: "[int, string]",
		},
		{
			name: "generic receiver (should return empty)",
			data: struct {
				TypeParams []*models.TypeParam
				Receiver   *models.Receiver
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "any"},
				},
				Receiver: &models.Receiver{
					Field: &models.Field{
						Type: &models.Expression{Value: "Set[T]"},
					},
				},
			},
			want: "",
		},
		{
			name: "non-generic receiver",
			data: struct {
				TypeParams []*models.TypeParam
				Receiver   *models.Receiver
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "any"},
				},
				Receiver: &models.Receiver{
					Field: &models.Field{
						Type: &models.Expression{Value: "MyStruct"},
					},
				},
			},
			want: "[int]",
		},
		{
			name: "pointer data",
			data: &struct {
				TypeParams []*models.TypeParam
				Receiver   *models.Receiver
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "any"},
				},
				Receiver: nil,
			},
			want: "[int]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := typeArguments(tt.data)
			if got != tt.want {
				t.Errorf("typeArguments() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test mapConstraintToType function
func Test_mapConstraintToType(t *testing.T) {
	tests := []struct {
		name       string
		constraint string
		want       string
	}{
		{
			name:       "any constraint",
			constraint: "any",
			want:       "int",
		},
		{
			name:       "comparable constraint",
			constraint: "comparable",
			want:       "string",
		},
		{
			name:       "union type with basic types",
			constraint: "int | float64",
			want:       "int",
		},
		{
			name:       "union type with spaces",
			constraint: "int64 | float64 | string",
			want:       "int64",
		},
		{
			name:       "approximation constraint",
			constraint: "~int",
			want:       "int",
		},
		{
			name:       "approximation with spaces",
			constraint: "~ string",
			want:       "string",
		},
		{
			name:       "numeric constraint - Integer",
			constraint: "Integer",
			want:       "int",
		},
		{
			name:       "numeric constraint - Signed",
			constraint: "Signed",
			want:       "int",
		},
		{
			name:       "numeric constraint - Unsigned",
			constraint: "Unsigned",
			want:       "int",
		},
		{
			name:       "numeric constraint - Float",
			constraint: "Float",
			want:       "int",
		},
		{
			name:       "numeric constraint - Complex",
			constraint: "Complex",
			want:       "int",
		},
		{
			name:       "numeric constraint - Ordered",
			constraint: "Ordered",
			want:       "int",
		},
		{
			name:       "constraints package - Integer",
			constraint: "constraints.Integer",
			want:       "int",
		},
		{
			name:       "constraints package - Ordered",
			constraint: "constraints.Ordered",
			want:       "int",
		},
		{
			name:       "error type",
			constraint: "error",
			want:       "errors.New(\"test error\")",
		},
		{
			name:       "unknown interface",
			constraint: "CustomInterface",
			want:       "string",
		},
		{
			name:       "constraint with spaces",
			constraint: "  any  ",
			want:       "int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapConstraintToType(tt.constraint)
			if got != tt.want {
				t.Errorf("mapConstraintToType(%q) = %q, want %q", tt.constraint, got, tt.want)
			}
		})
	}
}

// Test isBasicTypeName function
func Test_isBasicTypeName(t *testing.T) {
	tests := []struct {
		name     string
		typeName string
		want     bool
	}{
		// Basic types - should return true
		{name: "bool", typeName: "bool", want: true},
		{name: "string", typeName: "string", want: true},
		{name: "int", typeName: "int", want: true},
		{name: "int8", typeName: "int8", want: true},
		{name: "int16", typeName: "int16", want: true},
		{name: "int32", typeName: "int32", want: true},
		{name: "int64", typeName: "int64", want: true},
		{name: "uint", typeName: "uint", want: true},
		{name: "uint8", typeName: "uint8", want: true},
		{name: "uint16", typeName: "uint16", want: true},
		{name: "uint32", typeName: "uint32", want: true},
		{name: "uint64", typeName: "uint64", want: true},
		{name: "uintptr", typeName: "uintptr", want: true},
		{name: "byte", typeName: "byte", want: true},
		{name: "rune", typeName: "rune", want: true},
		{name: "float32", typeName: "float32", want: true},
		{name: "float64", typeName: "float64", want: true},
		{name: "complex64", typeName: "complex64", want: true},
		{name: "complex128", typeName: "complex128", want: true},
		// Non-basic types - should return false
		{name: "MyStruct", typeName: "MyStruct", want: false},
		{name: "interface{}", typeName: "interface{}", want: false},
		{name: "error", typeName: "error", want: false},
		{name: "[]int", typeName: "[]int", want: false},
		{name: "map[string]int", typeName: "map[string]int", want: false},
		{name: "chan int", typeName: "chan int", want: false},
		{name: "empty string", typeName: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBasicTypeName(tt.typeName)
			if got != tt.want {
				t.Errorf("isBasicTypeName(%q) = %v, want %v", tt.typeName, got, tt.want)
			}
		})
	}
}

// Test isNumericConstraintName function
func Test_isNumericConstraintName(t *testing.T) {
	tests := []struct {
		name string
		typ  string
		want bool
	}{
		// Numeric constraints - should return true
		{name: "Integer", typ: "Integer", want: true},
		{name: "Signed", typ: "Signed", want: true},
		{name: "Unsigned", typ: "Unsigned", want: true},
		{name: "Float", typ: "Float", want: true},
		{name: "Complex", typ: "Complex", want: true},
		{name: "Ordered", typ: "Ordered", want: true},
		{name: "constraints.Integer", typ: "constraints.Integer", want: true},
		{name: "constraints.Signed", typ: "constraints.Signed", want: true},
		{name: "constraints.Unsigned", typ: "constraints.Unsigned", want: true},
		{name: "constraints.Float", typ: "constraints.Float", want: true},
		{name: "constraints.Complex", typ: "constraints.Complex", want: true},
		{name: "constraints.Ordered", typ: "constraints.Ordered", want: true},
		{name: "pkg.Integer", typ: "pkg.Integer", want: true},
		// Non-numeric constraints - should return false
		{name: "any", typ: "any", want: false},
		{name: "comparable", typ: "comparable", want: false},
		{name: "CustomInterface", typ: "CustomInterface", want: false},
		{name: "int", typ: "int", want: false},
		{name: "string", typ: "string", want: false},
		{name: "empty", typ: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNumericConstraintName(tt.typ)
			if got != tt.want {
				t.Errorf("isNumericConstraintName(%q) = %v, want %v", tt.typ, got, tt.want)
			}
		})
	}
}

// Test fieldType function
func Test_fieldType(t *testing.T) {
	tests := []struct {
		name  string
		data  interface{}
		field *models.Field
		want  string
	}{
		{
			name:  "nil field",
			data:  struct{}{},
			field: nil,
			want:  "",
		},
		{
			name: "field with nil type",
			data: struct{}{},
			field: &models.Field{
				Name: "test",
				Type: nil,
			},
			want: "",
		},
		{
			name: "non-generic field",
			data: struct {
				TypeParams []*models.TypeParam
			}{
				TypeParams: nil,
			},
			field: &models.Field{
				Name: "value",
				Type: &models.Expression{Value: "int"},
			},
			want: "int",
		},
		{
			name: "generic field - type parameter substitution",
			data: struct {
				TypeParams []*models.TypeParam
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "any"},
				},
			},
			field: &models.Field{
				Name: "value",
				Type: &models.Expression{Value: "T"},
			},
			want: "int",
		},
		{
			name: "generic field - slice of type parameter",
			data: struct {
				TypeParams []*models.TypeParam
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "comparable"},
				},
			},
			field: &models.Field{
				Name: "items",
				Type: &models.Expression{Value: "[]T"},
			},
			want: "[]string",
		},
		{
			name: "generic field - multiple type parameters",
			data: struct {
				TypeParams []*models.TypeParam
			}{
				TypeParams: []*models.TypeParam{
					{Name: "K", Constraint: "comparable"},
					{Name: "V", Constraint: "any"},
				},
			},
			field: &models.Field{
				Name: "mapping",
				Type: &models.Expression{Value: "map[K]V"},
			},
			want: "map[string]int",
		},
		{
			name: "pointer data",
			data: &struct {
				TypeParams []*models.TypeParam
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "any"},
				},
			},
			field: &models.Field{
				Name: "value",
				Type: &models.Expression{Value: "T"},
			},
			want: "int",
		},
		{
			name: "field with star (pointer)",
			data: struct {
				TypeParams []*models.TypeParam
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "any"},
				},
			},
			field: &models.Field{
				Name: "ptr",
				Type: &models.Expression{
					Value:  "T",
					IsStar: true,
				},
			},
			want: "*int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fieldType(tt.data, tt.field)
			if got != tt.want {
				t.Errorf("fieldType() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test receiverType function
func Test_receiverType(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		receiver *models.Receiver
		want     string
	}{
		{
			name:     "nil receiver",
			data:     struct{}{},
			receiver: nil,
			want:     "",
		},
		{
			name: "receiver with nil type",
			data: struct{}{},
			receiver: &models.Receiver{
				Field: &models.Field{
					Type: nil,
				},
			},
			want: "",
		},
		{
			name: "non-generic receiver",
			data: struct {
				TypeParams []*models.TypeParam
			}{
				TypeParams: nil,
			},
			receiver: &models.Receiver{
				Field: &models.Field{
					Type: &models.Expression{Value: "MyStruct"},
				},
			},
			want: "MyStruct",
		},
		{
			name: "generic receiver - type parameter substitution",
			data: struct {
				TypeParams []*models.TypeParam
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "any"},
				},
			},
			receiver: &models.Receiver{
				Field: &models.Field{
					Type: &models.Expression{Value: "Set[T]"},
				},
			},
			want: "Set[int]",
		},
		{
			name: "generic receiver - multiple type params",
			data: struct {
				TypeParams []*models.TypeParam
			}{
				TypeParams: []*models.TypeParam{
					{Name: "K", Constraint: "comparable"},
					{Name: "V", Constraint: "any"},
				},
			},
			receiver: &models.Receiver{
				Field: &models.Field{
					Type: &models.Expression{Value: "Map[K, V]"},
				},
			},
			want: "Map[string, int]",
		},
		{
			name: "pointer data",
			data: &struct {
				TypeParams []*models.TypeParam
			}{
				TypeParams: []*models.TypeParam{
					{Name: "T", Constraint: "comparable"},
				},
			},
			receiver: &models.Receiver{
				Field: &models.Field{
					Type: &models.Expression{Value: "Container[T]"},
				},
			},
			want: "Container[string]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := receiverType(tt.data, tt.receiver)
			if got != tt.want {
				t.Errorf("receiverType() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test fieldName function
func Test_fieldName(t *testing.T) {
	tests := []struct {
		name  string
		field *models.Field
		want  string
	}{
		{
			name: "named field",
			field: &models.Field{
				Name: "myField",
				Type: &models.Expression{Value: "int"},
			},
			want: "myField",
		},
		{
			name: "unnamed field",
			field: &models.Field{
				Name: "",
				Type: &models.Expression{Value: "string"},
			},
			want: "string",
		},
		{
			name: "unnamed field with pointer",
			field: &models.Field{
				Name: "",
				Type: &models.Expression{
					Value:  "MyStruct",
					IsStar: true,
				},
			},
			want: "*MyStruct",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fieldName(tt.field)
			if got != tt.want {
				t.Errorf("fieldName() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test receiverName function
func Test_receiverName(t *testing.T) {
	tests := []struct {
		name     string
		receiver *models.Receiver
		want     string
	}{
		{
			name: "named receiver",
			receiver: &models.Receiver{
				Field: &models.Field{
					Name: "s",
					Type: &models.Expression{Value: "MyStruct"},
				},
			},
			want: "s",
		},
		{
			name: "unnamed receiver",
			receiver: &models.Receiver{
				Field: &models.Field{
					Name: "",
					Type: &models.Expression{Value: "MyStruct"},
				},
			},
			want: "m",
		},
		{
			name: "receiver named 'name' (conflict avoidance)",
			receiver: &models.Receiver{
				Field: &models.Field{
					Name: "name",
					Type: &models.Expression{Value: "MyStruct"},
				},
			},
			want: "n",
		},
		{
			name: "receiver named 't' (conflict avoidance)",
			receiver: &models.Receiver{
				Field: &models.Field{
					Name: "t",
					Type: &models.Expression{Value: "MyStruct"},
				},
			},
			want: "tr",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := receiverName(tt.receiver)
			if got != tt.want {
				t.Errorf("receiverName() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test parameterName function
func Test_parameterName(t *testing.T) {
	tests := []struct {
		name  string
		field *models.Field
		want  string
	}{
		{
			name: "named parameter",
			field: &models.Field{
				Name:  "input",
				Index: 0,
			},
			want: "input",
		},
		{
			name: "unnamed parameter index 0",
			field: &models.Field{
				Name:  "",
				Index: 0,
			},
			want: "in0",
		},
		{
			name: "unnamed parameter index 1",
			field: &models.Field{
				Name:  "",
				Index: 1,
			},
			want: "in1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parameterName(tt.field)
			if got != tt.want {
				t.Errorf("parameterName() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test wantName function
func Test_wantName(t *testing.T) {
	tests := []struct {
		name  string
		field *models.Field
		want  string
	}{
		{
			name: "named result",
			field: &models.Field{
				Name:  "result",
				Index: 0,
			},
			want: "wantResult",
		},
		{
			name: "unnamed result index 0",
			field: &models.Field{
				Name:  "",
				Index: 0,
			},
			want: "want",
		},
		{
			name: "unnamed result index 1",
			field: &models.Field{
				Name:  "",
				Index: 1,
			},
			want: "want1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wantName(tt.field)
			if got != tt.want {
				t.Errorf("wantName() = %q, want %q", got, tt.want)
			}
		})
	}
}

// Test gotName function
func Test_gotName(t *testing.T) {
	tests := []struct {
		name  string
		field *models.Field
		want  string
	}{
		{
			name: "named result",
			field: &models.Field{
				Name:  "output",
				Index: 0,
			},
			want: "gotOutput",
		},
		{
			name: "unnamed result index 0",
			field: &models.Field{
				Name:  "",
				Index: 0,
			},
			want: "got",
		},
		{
			name: "unnamed result index 1",
			field: &models.Field{
				Name:  "",
				Index: 1,
			},
			want: "got1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gotName(tt.field)
			if got != tt.want {
				t.Errorf("gotName() = %q, want %q", got, tt.want)
			}
		})
	}
}
