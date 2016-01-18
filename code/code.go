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
	Parameters []*Field
	Results    []*Field
}

type Info struct {
	Package string
	Funcs   []*Function
}

func Read(path string) *Info {
	fmt.Println(path)
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Parsing file: %v", err)
	}
	log.Printf("%#v", f)
	info := &Info{
		Package: f.Name.Name,
	}
	for _, d := range f.Decls {
		fDecl, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		f := &Function{
			Name: fDecl.Name.Name,
		}
		for _, field := range fDecl.Type.Params.List {
			log.Printf("%v", field.Type)
		}
		info.Funcs = append(info.Funcs, f)
	}
	return info
}
