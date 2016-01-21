package code

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

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
		Name:       fDecl.Name.Name,
		IsExported: fDecl.Name.IsExported(),
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
		return fmt.Sprintf("*%v", parseExpr(v.X))
	case *ast.SelectorExpr:
		return fmt.Sprintf("%v.%v", v.X, v.Sel)
	case *ast.MapType:
		return fmt.Sprintf("map[%v]%v", parseExpr(v.Key), parseExpr(v.Value))
	case *ast.ArrayType:
		return fmt.Sprintf("[]%v", parseExpr(v.Elt))
	case *ast.FuncType:
		var ps, rs []string
		if v.Params != nil {
			for _, p := range v.Params.List {
				ps = append(ps, parseExpr(p.Type))
			}
		}
		if v.Results != nil {
			for _, r := range v.Results.List {
				rs = append(rs, parseExpr(r.Type))
			}
		}
		if len(rs) < 2 {
			return fmt.Sprintf("func(%v) %v", strings.Join(ps, ","), strings.Join(rs, ""))
		}
		return fmt.Sprintf("func(%v) (%v)", strings.Join(ps, ","), strings.Join(rs, ","))
	default:
		return fmt.Sprintf("%v", v)
	}
}
