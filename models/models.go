package models

type Field struct {
	Name string
	Type string
}

func (f *Field) IsError() bool {
	return f.Type == "error"
}

type Function struct {
	Name       string
	Receiver   *Field
	Parameters []*Field
	Results    []*Field
}

func (f *Function) ReturnsError() bool {
	for _, r := range f.Results {
		if r.IsError() {
			return true
		}
	}
	return false
}

func (f *Function) ReturnsMultiple() bool {
	count := 0
	for _, r := range f.Results {
		if !r.IsError() {
			count++
		}
	}
	return count > 1
}

type Info struct {
	Package string
	Funcs   []*Function
}
