package models

import (
	"strings"
	"unicode"
)

type Expression struct {
	Value      string
	IsStar     bool
	IsVariadic bool
	IsWriter   bool
	Underlying string
}

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

type Field struct {
	Name  string
	Type  *Expression
	Index int
}

func (f *Field) IsWriter() bool {
	return f.Type.IsWriter
}

func (f *Field) IsStruct() bool {
	return strings.HasPrefix(f.Type.Underlying, "struct")
}

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

func (f *Field) IsNamed() bool {
	return f.Name != "" && f.Name != "_"
}

func (f *Field) ShortName() string {
	return strings.ToLower(string([]rune(f.Type.Value)[0]))
}

type Receiver struct {
	*Field
	Fields []*Field
}

// TypeParam represents a type parameter in a generic function or type
type TypeParam struct {
	Name       string // e.g., "T", "K", "V"
	Constraint string // e.g., "any", "comparable", "int64 | float64"
}

type Function struct {
	Name         string
	IsExported   bool
	Receiver     *Receiver
	Parameters   []*Field
	Results      []*Field
	ReturnsError bool
	TypeParams   []*TypeParam // Type parameters for generic functions
}

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

func (f *Function) ReturnsMultiple() bool {
	return len(f.Results) > 1
}

func (f *Function) OnlyReturnsOneValue() bool {
	return len(f.Results) == 1 && !f.ReturnsError
}

func (f *Function) OnlyReturnsError() bool {
	return len(f.Results) == 0 && f.ReturnsError
}

func (f *Function) FullName() string {
	var r string
	if f.Receiver != nil {
		r = f.Receiver.Type.Value
	}
	return strings.Title(r) + strings.Title(f.Name)
}

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

func (f *Function) IsNaked() bool {
	return f.Receiver == nil && len(f.Parameters) == 0 && len(f.Results) == 0
}

func (f *Function) IsGeneric() bool {
	return len(f.TypeParams) > 0
}

func (f *Function) HasGenericReceiver() bool {
	return f.Receiver != nil && strings.Contains(f.Receiver.Type.Value, "[")
}

// TypeParamMapping returns a map from type parameter names to concrete types for testing
func (f *Function) TypeParamMapping() map[string]string {
	mapping := make(map[string]string)
	for _, tp := range f.TypeParams {
		mapping[tp.Name] = mapConstraintToConcreteType(tp.Constraint)
	}
	return mapping
}

// TypeParamMappings returns multiple mappings for generating tests with different type combinations
// Returns a slice of mappings, each representing a different concrete instantiation
func (f *Function) TypeParamMappings() []map[string]string {
	if !f.IsGeneric() {
		return nil
	}

	// For now, generate 2-3 variants per type parameter
	var allMappings []map[string]string

	// Generate the primary mapping (using defaults)
	primaryMapping := make(map[string]string)
	for _, tp := range f.TypeParams {
		primaryMapping[tp.Name] = mapConstraintToConcreteType(tp.Constraint)
	}
	allMappings = append(allMappings, primaryMapping)

	// Generate alternative mappings
	// For single type parameter, add 1-2 more variants
	if len(f.TypeParams) == 1 {
		tp := f.TypeParams[0]
		alternatives := getAlternativeTypes(tp.Constraint)
		for _, altType := range alternatives {
			altMapping := make(map[string]string)
			altMapping[tp.Name] = altType
			allMappings = append(allMappings, altMapping)
		}
	}

	return allMappings
}

// getAlternativeTypes returns alternative concrete types for a constraint
func getAlternativeTypes(constraint string) []string {
	constraint = strings.TrimSpace(constraint)

	switch {
	case constraint == "any":
		return []string{"string"} // Primary is int, alternative is string
	case constraint == "comparable":
		return []string{"int"} // Primary is string, alternative is int
	case strings.Contains(constraint, "|"):
		// For union types, return other options beyond the first
		parts := strings.Split(constraint, "|")
		var alternatives []string
		for i := 1; i < len(parts) && i < 3; i++ { // Max 2 alternatives
			alternatives = append(alternatives, strings.TrimSpace(parts[i]))
		}
		return alternatives
	default:
		return nil
	}
}

// mapConstraintToConcreteType maps a constraint to a concrete type
func mapConstraintToConcreteType(constraint string) string {
	constraint = strings.TrimSpace(constraint)

	switch {
	case constraint == "any":
		return "int"
	case constraint == "comparable":
		return "string"
	case strings.Contains(constraint, "|"):
		// Union type - pick the first option
		parts := strings.Split(constraint, "|")
		return strings.TrimSpace(parts[0])
	case strings.Contains(constraint, "~"):
		// Approximation constraint like ~int
		constraint = strings.TrimPrefix(constraint, "~")
		return strings.TrimSpace(constraint)
	default:
		return "string"
	}
}

// SubstituteType replaces type parameter names in a type string with concrete types
func (f *Function) SubstituteType(typeStr string) string {
	if !f.IsGeneric() {
		return typeStr
	}

	result := typeStr
	mapping := f.TypeParamMapping()

	// Replace each type parameter with its concrete type
	// We need to be careful about word boundaries to avoid replacing parts of type names
	for paramName, concreteType := range mapping {
		// Simple replacement for now - this handles most cases
		// More sophisticated parsing would be needed for complex nested types
		result = strings.ReplaceAll(result, paramName, concreteType)
	}

	return result
}

type Import struct {
	Name, Path string
}

type Header struct {
	Comments []string
	Package  string
	Imports  []*Import
	Code     []byte
}

type Path string

func (p Path) TestPath() string {
	if !p.IsTestPath() {
		return strings.TrimSuffix(string(p), ".go") + "_test.go"
	}
	return string(p)
}

func (p Path) IsTestPath() bool {
	return strings.HasSuffix(string(p), "_test.go")
}
