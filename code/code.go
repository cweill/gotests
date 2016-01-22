package code

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/cweill/gotests/models"
)

func Parse(path string) (*models.SourceInfo, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %v", err)
	}
	info := &models.SourceInfo{
		Header: &models.SourceHeader{
			Pkg:  parseExpr(f.Name).String(),
			Imps: parseImports(f.Imports),
		},
	}
	for _, d := range f.Decls {
		fDecl, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		info.Funcs = append(info.Funcs, parseFunc(fDecl))
	}
	return info, nil
}

func ParseHeader(srcPath, testPath string) (*models.CodeHeader, error) {
	fset := token.NewFileSet()
	sf, err := parser.ParseFile(fset, srcPath, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %v", err)
	}
	fset = token.NewFileSet()
	tf, err := parser.ParseFile(fset, testPath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %v", err)
	}
	tf.Imports = append(tf.Imports, sf.Imports...)

	var furthestPos int
	for _, node := range tf.Imports {
		if pos := int(node.End()); pos > furthestPos {
			furthestPos = pos
		}
	}
	b, err := ioutil.ReadFile(testPath)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	h := &models.CodeHeader{
		Pkg:  parseExpr(tf.Name).String(),
		Imps: parseImports(tf.Imports),
		Code: b[furthestPos:],
	}
	return h, nil
}

func parseFunc(fDecl *ast.FuncDecl) *models.Function {
	f := &models.Function{
		Name:       parseExpr(fDecl.Name).String(),
		IsExported: fDecl.Name.IsExported(),
	}
	if fDecl.Recv != nil && fDecl.Recv.List != nil {
		f.Receiver = parseField(fDecl.Recv.List[0])[0]
	}
	if fDecl.Type.Params != nil {
		for _, fi := range fDecl.Type.Params.List {
			for _, pf := range parseField(fi) {
				f.Parameters = append(f.Parameters, pf)
			}
		}
	}
	if fDecl.Type.Results != nil {
		for _, fi := range fDecl.Type.Results.List {
			for _, mf := range parseField(fi) {
				if mf.Type.String() == "error" {
					f.ReturnsError = true
				} else {
					f.Results = append(f.Results, mf)
				}
			}
		}
	}
	return f
}

func parseImports(imps []*ast.ImportSpec) []*models.Import {
	var is []*models.Import
	for _, imp := range imps {
		var n string
		if imp.Name != nil {
			n = parseExpr(imp.Name).String()
		}
		is = append(is, &models.Import{
			Name: n,
			Path: parseExpr(imp.Path).String(),
		})
	}
	return is
}

func parseField(f *ast.Field) []*models.Field {
	if f == nil {
		return nil
	}
	t := parseExpr(f.Type)
	if len(f.Names) == 0 {
		return []*models.Field{{
			Type: t,
		}}
	}
	var fs []*models.Field
	for _, n := range f.Names {
		fs = append(fs, &models.Field{
			Name: n.Name,
			Type: t,
		})
	}
	return fs
}

func parseExpr(e ast.Expr) models.Expression {
	switch v := e.(type) {
	case *ast.Ident:
		return &models.Identity{Value: v.Name}
	case *ast.BasicLit:
		return &models.BasicLit{Value: v.Value}
	case *ast.InterfaceType:
		return &models.InterfaceType{}
	case *ast.StarExpr:
		return &models.StarExpr{X: parseExpr(v.X)}
	case *ast.SelectorExpr:
		return &models.SelectorExpr{X: parseExpr(v.X), Sel: parseExpr(v.Sel)}
	case *ast.MapType:
		return &models.MapExpr{Key: parseExpr(v.Key), Value: parseExpr(v.Value)}
	case *ast.ArrayType:
		return &models.ArrayExpr{Elt: parseExpr(v.Elt)}
	case *ast.Ellipsis:
		return &models.Ellipsis{Elt: parseExpr(v.Elt)}
	case *ast.FuncType:
		var ps, rs []models.Expression
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
		return &models.FuncType{Params: ps, Results: rs}
	default:
		return &models.Identity{Value: fmt.Sprintf("%v", v)}
	}
}
