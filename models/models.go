package models

import (
	"fmt"
	"sort"
	"strings"
)

type Expression interface {
	IsVariadic() bool
	String() string
}

type NonVariadicExpr struct {
	Value string
}

func (e *NonVariadicExpr) String() string { return e.Value }

func (*NonVariadicExpr) IsVariadic() bool { return false }

type VariadicExpr struct {
	Elt string
}

func (e *VariadicExpr) String() string {
	return "[]" + e.Elt
}

func (*VariadicExpr) IsVariadic() bool { return true }

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
	return "Test" + strings.Title(f.Name)
}

type Header struct {
	Package string
	Imports []*Import
	Code    []byte
}

type Import struct {
	Name, Path string
}

type SourceInfo struct {
	Header *Header
	Funcs  []*Function
}

func (i *SourceInfo) TestableFuncs(onlyFuncs, exclFuncs []string) []*Function {
	sort.Strings(onlyFuncs)
	sort.Strings(exclFuncs)
	var fs []*Function
	for _, f := range i.Funcs {
		if f.Receiver == nil && len(f.Parameters) == 0 && len(f.Results) == 0 {
			continue
		}
		if len(exclFuncs) > 0 && (contains(exclFuncs, f.Name) || contains(exclFuncs, f.TestName())) {
			continue
		}
		if len(onlyFuncs) > 0 && !contains(onlyFuncs, f.Name) && !contains(onlyFuncs, f.TestName()) {
			continue
		}
		fs = append(fs, f)
	}
	return fs
}

func (i *SourceInfo) UsesReflection() bool {
	for _, f := range i.Funcs {
		for _, fi := range f.Results {
			if !fi.IsScalar() {
				return true
			}
		}
	}
	return false
}

func contains(ss []string, s string) bool {
	if i := sort.SearchStrings(ss, s); i < len(ss) && ss[i] == s {
		return true
	}
	return false
}

type Path string

func (p Path) TestPath() string {
	if p.IsTestPath() {
		return string(p)
	}
	return strings.TrimSuffix(string(p), ".go") + "_test.go"
}

func (p Path) IsTestPath() bool {
	return strings.HasSuffix(string(p), "_test.go")
}
