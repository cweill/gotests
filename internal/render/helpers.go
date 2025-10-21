package render

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cweill/gotests/internal/models"
)

func fieldName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = f.Name
	} else {
		n = f.Type.String()
	}
	return n
}

func receiverName(f *models.Receiver) string {
	var n string
	if f.IsNamed() {
		n = f.Name
	} else {
		n = f.ShortName()
	}
	if n == "name" {
		// Avoid conflict with test struct's "name" field.
		n = "n"
	} else if n == "t" {
		// Avoid conflict with test argument.
		// "tr" is short for t receiver.
		n = "tr"
	}
	return n
}

func parameterName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = f.Name
	} else {
		n = fmt.Sprintf("in%v", f.Index)
	}
	return n
}

func wantName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = "want" + strings.Title(f.Name)
	} else if f.Index == 0 {
		n = "want"
	} else {
		n = fmt.Sprintf("want%v", f.Index)
	}
	return n
}

func gotName(f *models.Field) string {
	var n string
	if f.IsNamed() {
		n = "got" + strings.Title(f.Name)
	} else if f.Index == 0 {
		n = "got"
	} else {
		n = fmt.Sprintf("got%v", f.Index)
	}
	return n
}

// typeArguments generates concrete type arguments for generic functions.
// Maps type parameter constraints to sensible concrete types for testing.
// Accepts interface{} to work with the template wrapper struct.
func typeArguments(data interface{}) string {
	var typeParams []*models.TypeParam
	var isGeneric, hasGenericReceiver bool

	// Use reflection to extract fields from the wrapper struct
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		// Look for TypeParams field
		if tpField := v.FieldByName("TypeParams"); tpField.IsValid() {
			if tps, ok := tpField.Interface().([]*models.TypeParam); ok {
				typeParams = tps
				isGeneric = len(typeParams) > 0
			}
		}

		// Look for Receiver field to check if it's a generic receiver
		if recvField := v.FieldByName("Receiver"); recvField.IsValid() && !recvField.IsNil() {
			if recv, ok := recvField.Interface().(*models.Receiver); ok && recv != nil {
				hasGenericReceiver = recv.Type != nil && strings.Contains(recv.Type.Value, "[")
			}
		}
	}

	if !isGeneric && !hasGenericReceiver {
		return ""
	}

	// For generic receivers, we don't add type args to the method call
	// The receiver already has them
	if hasGenericReceiver {
		return ""
	}

	var args []string
	for _, tp := range typeParams {
		concreteType := mapConstraintToType(tp.Constraint)
		args = append(args, concreteType)
	}

	if len(args) == 0 {
		return ""
	}

	return "[" + strings.Join(args, ", ") + "]"
}

// mapConstraintToType maps a type constraint to a concrete type for testing.
// Uses intelligent defaults based on the constraint type.
func mapConstraintToType(constraint string) string {
	constraint = strings.TrimSpace(constraint)

	switch {
	case constraint == "any":
		// For 'any', use int as a simple value type
		return "int"
	case constraint == "comparable":
		// For 'comparable', use string as it's commonly comparable
		return "string"
	case strings.Contains(constraint, "|"):
		// Union type - pick the first option
		parts := strings.Split(constraint, "|")
		firstType := strings.TrimSpace(parts[0])
		// If it's a basic type, use it directly
		if isBasicTypeName(firstType) {
			return firstType
		}
		return firstType
	case strings.Contains(constraint, "~"):
		// Approximation constraint like ~int or ~string
		// Extract the underlying type
		constraint = strings.TrimPrefix(constraint, "~")
		constraint = strings.TrimSpace(constraint)
		return constraint
	case isNumericConstraintName(constraint):
		// Handle common numeric constraint interfaces
		return "int"
	case constraint == "error":
		return "errors.New(\"test error\")"
	default:
		// For interface constraints or unknown types, use string as safe default
		return "string"
	}
}

// isBasicTypeName checks if a type name is a Go basic type
func isBasicTypeName(typeName string) bool {
	basicTypes := []string{
		"bool", "string",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
		"byte", "rune", "float32", "float64",
		"complex64", "complex128",
	}
	for _, bt := range basicTypes {
		if typeName == bt {
			return true
		}
	}
	return false
}

// isNumericConstraintName checks for common numeric constraint interfaces
func isNumericConstraintName(name string) bool {
	numericConstraints := []string{
		"Integer", "Signed", "Unsigned", "Float", "Complex", "Ordered",
		"constraints.Integer", "constraints.Signed", "constraints.Unsigned",
		"constraints.Float", "constraints.Complex", "constraints.Ordered",
	}
	for _, nc := range numericConstraints {
		if name == nc || strings.HasSuffix(name, "."+nc) {
			return true
		}
	}
	return false
}

// fieldType returns the type of a field, substituting type parameters for generic functions
// This should be called from templates as: {{FieldType $ .}}
// where $ is the full template data context and . is the Field
func fieldType(data interface{}, field *models.Field) string {
	if field == nil || field.Type == nil {
		return ""
	}

	typeStr := field.Type.String()

	// Use reflection to get TypeParams from the wrapper struct
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		// Get TypeParams to build mapping
		if tpField := v.FieldByName("TypeParams"); tpField.IsValid() {
			if tps, ok := tpField.Interface().([]*models.TypeParam); ok && len(tps) > 0 {
				// Build mapping
				mapping := make(map[string]string)
				for _, tp := range tps {
					mapping[tp.Name] = mapConstraintToType(tp.Constraint)
				}

				// Replace each type parameter in the type string
				for paramName, concreteType := range mapping {
					typeStr = strings.ReplaceAll(typeStr, paramName, concreteType)
				}
			}
		}
	}

	return typeStr
}

// receiverType returns the type of a receiver value (not pointer), substituting type parameters
// For generic receivers like *Set[T], this returns Set[string] (concrete instantiation)
func receiverType(data interface{}, receiver *models.Receiver) string {
	if receiver == nil || receiver.Type == nil {
		return ""
	}

	typeStr := receiver.Type.Value // Use .Value not .String() to avoid the * prefix

	// Use reflection to get TypeParams from the wrapper struct
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		// Get TypeParams to build mapping (but use receiver's type params, not function's)
		// For methods on generic types, we need to substitute based on the receiver's type params
		// The receiver type like Set[T] has its own type parameters
		// We'll use a simple heuristic: substitute common type param names

		// Extract type params from receiver type if it's generic (contains [...])
		if strings.Contains(typeStr, "[") {
			// For simplicity, get function's TypeParams and use them
			// This works because method type params shadow the receiver's type params
			if tpField := v.FieldByName("TypeParams"); tpField.IsValid() {
				if tps, ok := tpField.Interface().([]*models.TypeParam); ok && len(tps) > 0 {
					mapping := make(map[string]string)
					for _, tp := range tps {
						mapping[tp.Name] = mapConstraintToType(tp.Constraint)
					}

					// Replace type parameters
					for paramName, concreteType := range mapping {
						typeStr = strings.ReplaceAll(typeStr, paramName, concreteType)
					}
				}
			}
		}
	}

	return typeStr
}
