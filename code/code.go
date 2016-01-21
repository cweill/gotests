package code

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	"github.com/cweill/gotests/models"
)

func Parse(path string) *models.Info {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Parsing file: %v", err)
	}
	info := &models.Info{
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

func parseFunc(fDecl *ast.FuncDecl) *models.Function {
	f := &models.Function{
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
			mf := parseField(fi)
			if mf.Type == "error" {
				f.ReturnsError = true
			} else {
				f.Results = append(f.Results, mf)
			}
		}
	}
	return f
}

func parseField(f *ast.Field) *models.Field {
	if f == nil {
		return nil
	}
	var n string
	if f.Names != nil {
		n = f.Names[0].Name
	}
	return &models.Field{
		Name: n,
		Type: parseExpr(f.Type),
	}
}

func parseExpr(e ast.Expr) string {
	switch v := e.(type) {
	case *ast.StarExpr:
		return fmt.Sprintf("*%v", v.X)
	case *ast.MapType:
		return fmt.Sprintf("map[%v]%v", parseExpr(v.Key), parseExpr(v.Value))
	case *ast.ArrayType:
		return fmt.Sprintf("[]%v", parseExpr(v.Elt))
	default:
		return fmt.Sprintf("%v", v)
	}
}
