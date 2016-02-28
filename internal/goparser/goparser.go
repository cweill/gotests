package goparser

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"

	"github.com/cweill/gotests/internal/models"
)

var ErrEmptyFile = errors.New("file is empty")

type Result struct {
	Header *models.Header
	Funcs  []*models.Function
}

type Parser struct {
	Importer types.Importer
}

func (p *Parser) Parse(srcPath string, files []models.Path) (*Result, error) {
	b, err := p.readFile(srcPath)
	if err != nil {
		return nil, err
	}
	fset := token.NewFileSet()
	f, err := p.parseFile(fset, srcPath)
	if err != nil {
		return nil, err
	}
	fs, err := p.parseFiles(fset, f, files)
	if err != nil {
		return nil, err
	}
	return &Result{
		Header: &models.Header{
			Package: f.Name.String(),
			Imports: parseImports(f.Imports),
			Code:    goCode(b, f),
		},
		Funcs: p.parseFunctions(fset, f, fs),
	}, nil
}

func (p *Parser) readFile(srcPath string) ([]byte, error) {
	b, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	if len(b) == 0 {
		return nil, ErrEmptyFile
	}
	return b, nil
}

func (p *Parser) parseFile(fset *token.FileSet, srcPath string) (*ast.File, error) {
	f, err := parser.ParseFile(fset, srcPath, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("target parser.ParseFile(): %v", err)
	}
	return f, nil
}

func (p *Parser) parseFiles(fset *token.FileSet, f *ast.File, files []models.Path) ([]*ast.File, error) {
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
	return fs, nil
}

func (p *Parser) parseFunctions(fset *token.FileSet, f *ast.File, fs []*ast.File) []*models.Function {
	ul, el := p.parseTypes(fset, fs)
	var funcs []*models.Function
	for _, d := range f.Decls {
		fDecl, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		funcs = append(funcs, parseFunc(fDecl, ul, el))
	}
	return funcs
}

func (p *Parser) parseTypes(fset *token.FileSet, fs []*ast.File) (map[string]types.Type, map[*types.Struct]ast.Expr) {
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
	return ul, el
}

// Returns the Go code below the imports block.
func goCode(b []byte, f *ast.File) []byte {
	furthestPos := f.Name.End()
	for _, node := range f.Imports {
		if pos := node.End(); pos > furthestPos {
			furthestPos = pos
		}
	}
	if furthestPos < token.Pos(len(b)) {
		furthestPos++
	}
	return b[furthestPos:]
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
	switch v := e.(type) {
	case *ast.StarExpr:
		val := types.ExprString(v.X)
		return &models.Expression{
			Value:      val,
			IsStar:     true,
			Underlying: underlying(val, ul),
		}
	case *ast.Ellipsis:
		exp := parseExpr(v.Elt, ul)
		return &models.Expression{
			Value:      exp.Value,
			IsStar:     exp.IsStar,
			IsVariadic: true,
			Underlying: underlying(exp.Value, ul),
		}
	default:
		val := types.ExprString(e)
		return &models.Expression{
			Value:      val,
			Underlying: underlying(val, ul),
			IsWriter:   val == "io.Writer",
		}
	}
}

func underlying(val string, ul map[string]types.Type) string {
	if ul[val] != nil {
		return ul[val].String()
	}
	return ""
}
