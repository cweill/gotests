package models

import (
	"strings"
	"unicode"
)

// Expression represents a type expression in Go code, including metadata about pointers, variadic parameters, and writers.
type Expression struct {
	Value      string
	IsStar     bool
	IsVariadic bool
	IsWriter   bool
	Underlying string
}

// String returns the string representation of the expression, including pointer and variadic prefixes.
func (e *Expression) String() string {
	value := e.Value
	if e.IsStar {
		value = "*" + value
	}
	if e.IsVariadic {
		return "[]" + value
	}
	return value
}

// Field represents a parameter, result, or struct field in a function or method signature.
type Field struct {
	Name  string
	Type  *Expression
	Index int
}

// IsWriter returns true if the field is an io.Writer.
func (f *Field) IsWriter() bool {
	return f.Type.IsWriter
}

// IsStruct returns true if the field's underlying type is a struct.
func (f *Field) IsStruct() bool {
	return strings.HasPrefix(f.Type.Underlying, "struct")
}

// IsBasicType returns true if the field is a Go basic type (bool, string, int, etc.).
func (f *Field) IsBasicType() bool {
	return isBasicType(f.Type.String()) || isBasicType(f.Type.Underlying)
}

func isBasicType(t string) bool {
	switch t {
	case "bool", "string", "int", "int8", "int16", "int32", "int64", "uint",
		"uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune",
		"float32", "float64", "complex64", "complex128":
		return true
	default:
		return false
	}
}

// IsNamed returns true if the field has a non-blank name.
func (f *Field) IsNamed() bool {
	return f.Name != "" && f.Name != "_"
}

// ShortName returns a short single-letter name based on the field's type.
func (f *Field) ShortName() string {
	return strings.ToLower(string([]rune(f.Type.Value)[0]))
}

// Receiver represents a method receiver, including its type and struct fields.
type Receiver struct {
	*Field
	Fields []*Field
}

// TypeParam represents a type parameter in a generic function or type.
type TypeParam struct {
	Name       string // e.g., "T", "K", "V"
	Constraint string // e.g., "any", "comparable", "int64 | float64"
}

// Function represents a function or method signature with its parameters, results, and metadata.
type Function struct {
	Name         string
	IsExported   bool
	Receiver     *Receiver
	Parameters   []*Field
	Results      []*Field
	ReturnsError bool
	TypeParams   []*TypeParam // Type parameters for generic functions
	Body         string       // Source code of the function body for AI context
}

// TestParameters returns the function's parameters excluding io.Writer parameters.
func (f *Function) TestParameters() []*Field {
	var ps []*Field
	for _, p := range f.Parameters {
		if p.IsWriter() {
			continue
		}
		ps = append(ps, p)
	}
	return ps
}

// TestResults returns the function's results plus any io.Writer parameters converted to string results.
func (f *Function) TestResults() []*Field {
	var ps []*Field
	ps = append(ps, f.Results...)
	for _, p := range f.Parameters {
		if !p.IsWriter() {
			continue
		}
		ps = append(ps, &Field{
			Name: p.Name,
			Type: &Expression{
				Value:      "string",
				IsWriter:   true,
				Underlying: "string",
			},
			Index: len(ps),
		})
	}
	return ps
}

// ReturnsMultiple returns true if the function returns more than one value.
func (f *Function) ReturnsMultiple() bool {
	return len(f.Results) > 1
}

// OnlyReturnsOneValue returns true if the function returns exactly one non-error value.
func (f *Function) OnlyReturnsOneValue() bool {
	return len(f.Results) == 1 && !f.ReturnsError
}

// OnlyReturnsError returns true if the function returns only an error.
func (f *Function) OnlyReturnsError() bool {
	return len(f.Results) == 0 && f.ReturnsError
}

// FullName returns the full name of the function, including the receiver type if it's a method.
func (f *Function) FullName() string {
	var r string
	if f.Receiver != nil {
		r = f.Receiver.Type.Value
	}
	return strings.Title(r) + strings.Title(f.Name)
}

// TestName returns the name to use for the generated test function.
func (f *Function) TestName() string {
	if strings.HasPrefix(f.Name, "Test") {
		return f.Name
	}
	if f.Receiver != nil {
		receiverType := f.Receiver.Type.Value
		// Strip type parameters from receiver type for test name
		// e.g., "Set[T]" -> "Set", "Map[K, V]" -> "Map"
		if idx := strings.Index(receiverType, "["); idx != -1 {
			receiverType = receiverType[:idx]
		}
		if unicode.IsLower([]rune(receiverType)[0]) {
			receiverType = "_" + receiverType
		}
		return "Test" + receiverType + "_" + f.Name
	}
	if unicode.IsLower([]rune(f.Name)[0]) {
		return "Test_" + f.Name
	}
	return "Test" + f.Name
}

// IsNaked returns true if the function has no receiver, parameters, or results.
func (f *Function) IsNaked() bool {
	return f.Receiver == nil && len(f.Parameters) == 0 && len(f.Results) == 0
}

// Import represents an import statement with an optional name and the import path.
type Import struct {
	Name, Path string
}

// Header represents the header of a Go file, including package name, imports, and any code between imports and declarations.
type Header struct {
	Comments []string
	Package  string
	Imports  []*Import
	Code     []byte
}

// Path represents a file system path.
type Path string

// TestPath returns the test file path for the given source file path.
func (p Path) TestPath() string {
	if !p.IsTestPath() {
		return strings.TrimSuffix(string(p), ".go") + "_test.go"
	}
	return string(p)
}

// IsTestPath returns true if the path is a test file path (ends with _test.go).
func (p Path) IsTestPath() bool {
	return strings.HasSuffix(string(p), "_test.go")
}
