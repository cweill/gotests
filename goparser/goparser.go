package goparser

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"

	"github.com/cweill/gotests/models"
)

func Parse(srcPath string, files []models.Path) (*models.SourceInfo, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, srcPath, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile: %v", err)
	}
	pkg := f.Name.String()
	var fs []*ast.File
	for _, file := range files {
		ff, err := parser.ParseFile(fset, string(file), nil, 0)
		if err != nil {
			return nil, fmt.Errorf("parser.ParseFile: %v", err)
		}
		if name := f.Name.String(); name != pkg {
			continue
		}
		fs = append(fs, ff)
	}
	conf := &types.Config{Importer: importer.Default()}
	ti := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	if _, err = conf.Check("", fset, fs, ti); err != nil {
		return nil, fmt.Errorf("conf.Check: %v", err)
	}
	ul := make(map[string]string)
	for _, t := range ti.Types {
		ul[t.Type.String()] = t.Type.Underlying().String()
	}
	info := &models.SourceInfo{
		Header: &models.Header{
			Package: pkg,
			Imports: parseImports(f.Imports),
		},
	}
	for _, d := range f.Decls {
		fDecl, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		info.Funcs = append(info.Funcs, parseFunc(fDecl, ul))
	}
	return info, nil
}

func ParseHeader(srcPath, testPath string) (*models.Header, error) {
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
	furthestPos := tf.Name.End()
	for _, node := range tf.Imports {
		if pos := node.End(); pos > furthestPos {
			furthestPos = pos
		}
	}
	tf.Imports = append(tf.Imports, sf.Imports...)
	b, err := ioutil.ReadFile(testPath)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	h := &models.Header{
		Package: tf.Name.String(),
		Imports: parseImports(tf.Imports),
		Code:    b[furthestPos+1:],
	}
	return h, nil
}

func parseFunc(fDecl *ast.FuncDecl, ul map[string]string) *models.Function {
	f := &models.Function{
		Name:       fDecl.Name.String(),
		IsExported: fDecl.Name.IsExported(),
	}
	if fDecl.Recv != nil && fDecl.Recv.List != nil {
		f.Receiver = parseFields(fDecl.Recv.List[0], ul)[0]
	}
	if fDecl.Type.Params != nil {
		for _, fi := range fDecl.Type.Params.List {
			for i, pf := range parseFields(fi, ul) {
				pf.Index = i
				f.Parameters = append(f.Parameters, pf)
			}
		}
	}
	if fDecl.Type.Results != nil {
		i := 0
		for _, fi := range fDecl.Type.Results.List {
			for _, pf := range parseFields(fi, ul) {
				if pf.Type.String() == "error" {
					f.ReturnsError = true
				} else {
					pf.Index = i
					f.Results = append(f.Results, pf)
					i++
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
			n = imp.Name.String()
		}
		is = append(is, &models.Import{
			Name: n,
			Path: imp.Path.Value,
		})
	}
	return is
}

func parseFields(f *ast.Field, ul map[string]string) []*models.Field {
	if f == nil {
		return nil
	}
	t := parseExpr(f.Type, ul)
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

func parseExpr(e ast.Expr, ul map[string]string) *models.Expression {
	switch v := e.(type) {
	case *ast.StarExpr:
		val := types.ExprString(v.X)
		return &models.Expression{
			Value:      val,
			IsStar:     true,
			Underlying: ul[val],
		}
	case *ast.Ellipsis:
		exp := parseExpr(v.Elt, ul)
		return &models.Expression{
			Value:      exp.Value,
			IsStar:     exp.IsStar,
			IsVariadic: true,
			Underlying: ul[exp.Value],
		}
	default:
		val := types.ExprString(e)
		return &models.Expression{
			Value:      val,
			Underlying: ul[val],
		}
	}
}
