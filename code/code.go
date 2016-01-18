package code

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

type Field struct {
	Name string
	Type string
}

type Function struct {
	Name       string
	Receiver   *Field
	Parameters []*Field
	Results    []*Field
}

type Info struct {
	Package string
	Funcs   []*Function
}

func Parse(path string) *Info {
	fmt.Println(path)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Parsing file: %v", err)
	}
	info := &Info{
		Package: f.Name.Name,
	}
	for _, d := range f.Decls {
		fDecl, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		info.Funcs = append(info.Funcs, parseFunc(fDecl))
	}
	return info
}

func parseFunc(fDecl *ast.FuncDecl) *Function {
	f := &Function{
		Name: fDecl.Name.Name,
	}
	if fDecl.Recv != nil && fDecl.Recv.List != nil {
		f.Receiver = parseField(fDecl.Recv.List[0])
	}
	if fDecl.Type.Params != nil {
		for _, fi := range fDecl.Type.Params.List {
			f.Parameters = append(f.Parameters, parseField(fi))
		}
	}
	if fDecl.Type.Results != nil {
		for _, fi := range fDecl.Type.Results.List {
			f.Results = append(f.Results, parseField(fi))
		}
	}
	return f
}

func parseField(f *ast.Field) *Field {
	if f == nil {
		return nil
	}
	var n, t string
	if f.Names != nil {
		n = f.Names[0].Name
	}
	switch v := f.Type.(type) {
	case *ast.StarExpr:
		t = fmt.Sprintf("* %v", v.X)
	}
	return &Field{
		Name: n,
		Type: t,
	}
}
