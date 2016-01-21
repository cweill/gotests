package models

type Field struct {
	Name string
	Type string
}

func (f *Field) IsScalar() bool {
	switch f.Type {
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
