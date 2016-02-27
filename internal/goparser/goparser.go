package goparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"

	"github.com/cweill/gotests/internal/models"
)

type Parser struct {
	Importer types.Importer
}

func (p *Parser) Parse(srcPath string, files []models.Path) (*models.SourceInfo, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, srcPath, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("target parser.ParseFile(): %v", err)
	}
	pkg := f.Name.String()
	var fs []*ast.File
	for _, file := range files {
		ff, err := parser.ParseFile(fset, string(file), nil, 0)
		if err != nil {
			return nil, fmt.Errorf("other file parser.ParseFile: %v", err)
		}
		if name := ff.Name.String(); name != pkg {
			continue
		}
		fs = append(fs, ff)
	}
	conf := &types.Config{
		Importer: p.Importer,
		// Adding a NO-OP error function ignores errors and performs best-effort
		// type checking. https://godoc.org/golang.org/x/tools/go/types#Config
		Error: func(error) {},
	}
	ti := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	// Note: conf.Check can fail, but since Info is not required data, it's ok.
	conf.Check("", fset, fs, ti)
	ul := make(map[string]types.Type)
	el := make(map[*types.Struct]ast.Expr)
	for e, t := range ti.Types {
		// Collect the underlying types.
		ul[t.Type.String()] = t.Type.Underlying()
		// Collect structs to determine the fields of a receiver.
		if v, ok := t.Type.(*types.Struct); ok {
			el[v] = e
		}
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
		info.Funcs = append(info.Funcs, parseFunc(fDecl, ul, el))
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
	if furthestPos < token.Pos(len(b)) {
		furthestPos++
	}
	h := &models.Header{
		Package: tf.Name.String(),
		Imports: parseImports(tf.Imports),
		Code:    b[furthestPos:],
	}
	return h, nil
}

func parseFunc(fDecl *ast.FuncDecl, ul map[string]types.Type, el map[*types.Struct]ast.Expr) *models.Function {
	f := &models.Function{
		Name:       fDecl.Name.String(),
		IsExported: fDecl.Name.IsExported(),
	}
	if fDecl.Recv != nil && fDecl.Recv.List != nil {
		f.Receiver = parseReceiver(fDecl.Recv.List[0], ul, el)
	}
	if fDecl.Type.Params != nil {
		i := 0
		for _, fi := range fDecl.Type.Params.List {
			for _, pf := range parseFields(fi, ul) {
				pf.Index = i
				f.Parameters = append(f.Parameters, pf)
				i++
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

func parseReceiver(f *ast.Field, ul map[string]types.Type, el map[*types.Struct]ast.Expr) *models.Receiver {
	r := &models.Receiver{
		Field: parseFields(f, ul)[0],
	}
	t, ok := ul[r.Type.Value]
	if !ok {
		return r
	}
	s, ok := t.(*types.Struct)
	if !ok {
		return r
	}
	st := el[s].(*ast.StructType)
	if st.Fields == nil {
		return r
	}
	for _, f := range st.Fields.List {
		r.Fields = append(r.Fields, parseFields(f, ul)...)
	}
	for i, f := range r.Fields {
		f.Name = s.Field(i).Name()
	}
	return r

}

func parseFields(f *ast.Field, ul map[string]types.Type) []*models.Field {
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

func parseExpr(e ast.Expr, ul map[string]types.Type) *models.Expression {
	var u string
	switch v := e.(type) {
	case *ast.StarExpr:
		val := types.ExprString(v.X)
		if ul[val] != nil {
			u = ul[val].String()
		}
		return &models.Expression{
			Value:      val,
			IsStar:     true,
			Underlying: u,
		}
	case *ast.Ellipsis:
		exp := parseExpr(v.Elt, ul)
		if ul[exp.Value] != nil {
			u = ul[exp.Value].String()
		}
		return &models.Expression{
			Value:      exp.Value,
			IsStar:     exp.IsStar,
			IsVariadic: true,
			Underlying: u,
		}
	default:
		val := types.ExprString(e)
		if ul[val] != nil {
			u = ul[val].String()
		}
		return &models.Expression{
			Value:      val,
			Underlying: u,
			IsWriter:   val == "io.Writer",
		}
	}
}
