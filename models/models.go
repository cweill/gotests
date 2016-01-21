package models

import (
	"fmt"
	"strings"
)

type Expression interface {
	IsVariadic() bool
	String() string
}

type Identity struct {
	Value string
}

func (s *Identity) String() string { return s.Value }

func (*Identity) IsVariadic() bool { return false }

type InterfaceType struct{}

func (*InterfaceType) String() string { return "interface{}" }

func (*InterfaceType) IsVariadic() bool { return false }

type StarExpr struct {
	X Expression
}

func (s *StarExpr) String() string {
	return fmt.Sprintf("*%v", s.X)
}

func (*StarExpr) IsVariadic() bool { return false }

type SelectorExpr struct {
	X, Sel Expression
}

func (s *SelectorExpr) String() string {
	return fmt.Sprintf("%v.%v", s.X, s.Sel)
}

func (*SelectorExpr) IsVariadic() bool { return false }

type MapExpr struct {
	Key, Value Expression
}

func (m *MapExpr) String() string {
	return fmt.Sprintf("map[%v]%v", m.Key, m.Value)
}

func (*MapExpr) IsVariadic() bool { return false }

type ArrayExpr struct {
	Elt Expression
}

func (a *ArrayExpr) String() string {
	return fmt.Sprintf("[]%v", a.Elt)
}

func (*ArrayExpr) IsVariadic() bool { return false }

type Ellipsis ArrayExpr

func (e *Ellipsis) String() string {
	return fmt.Sprintf("[]%v", e.Elt)
}

func (*Ellipsis) IsVariadic() bool { return true }

type FuncType struct {
	Params, Results []Expression
}

func (f *FuncType) String() string {
	var ps, rs []string
	for _, p := range f.Params {
		ps = append(ps, p.String())
	}
	for _, r := range f.Results {
		rs = append(rs, r.String())
	}
	if len(rs) < 2 {
		return fmt.Sprintf("func(%v) %v", strings.Join(ps, ","), strings.Join(rs, ""))
	}
	return fmt.Sprintf("func(%v) (%v)", strings.Join(ps, ","), strings.Join(rs, ","))
}

func (*FuncType) IsVariadic() bool { return false }

type Field struct {
	Name string
	Type Expression
}

func (f *Field) IsScalar() bool {
	switch f.Type.String() {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int", "int16", "int32", "int64", "float32", "float64", "complex64", "complex128", "byte", "rune", "bool", "string", "error":
		return true
	default:
		return false
	}
}

type Function struct {
	Name         string
	IsExported   bool
	Receiver     *Field
	Parameters   []*Field
	Results      []*Field
	ReturnsError bool
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

func (f *Function) TestName() string {
	return "Test" + f.Name
}

type Info struct {
	Package string
	Funcs   []*Function
}

func (i *Info) ExportedFuncs() []*Function {
	var fs []*Function
	for _, f := range i.Funcs {
		if f.IsExported {
			fs = append(fs, f)
		}
	}
	return fs
}

func (i *Info) UsesReflection() bool {
	for _, f := range i.Funcs {
		for _, fi := range f.Results {
			if !fi.IsScalar() {
				return true
			}
		}
	}
	return false
}
