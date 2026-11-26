package goparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/cweill/gotests/internal/models"
)

// QualifyTypes prepares fn for a test in the external <pkg>_test package.
//
// It rewrites the parameter, result and receiver field types so that types
// declared in the package under test (localTypes) are referenced through
// pkgName, e.g. "[]*Bar" becomes "[]*pkg.Bar". The receiver's own type is
// left untouched because it also names the generated test; templates qualify
// it with the ReceiverType helper. Unexported receiver fields are dropped,
// since an external package cannot set them.
//
// It returns a non-empty reason when fn cannot be tested from an external
// package: its receiver type, or a type in its signature, is unexported.
func QualifyTypes(fn *models.Function, pkgName string, localTypes map[string]bool) (reason string) {
	skip := make(map[string]bool)
	for _, tp := range fn.TypeParams {
		skip[tp.Name] = true
	}
	if fn.Receiver != nil {
		recv, err := parser.ParseExpr(fn.Receiver.Type.Value)
		if err != nil {
			return fmt.Sprintf("cannot parse receiver type %s", fn.Receiver.Type.Value)
		}
		base, args := splitTypeArgs(recv)
		if !token.IsExported(base) {
			return fmt.Sprintf("receiver type %s is unexported", base)
		}
		// A method may rename its type's parameters: func (s *Set[E]) Add(v E).
		for _, a := range args {
			skip[a] = true
		}
	}

	q := &qualifier{
		pkg:   pkgName,
		local: localTypes,
		skip:  skip,
		done:  make(map[*models.Expression]bool),
	}
	for _, f := range fn.Parameters {
		q.field(f)
	}
	for _, f := range fn.Results {
		q.field(f)
	}
	if r := fn.Receiver; r != nil {
		var fields []*models.Field
		for _, f := range r.Fields {
			if !token.IsExported(f.Name) {
				continue
			}
			q.field(f)
			fields = append(fields, f)
		}
		r.Fields = fields
	}
	if q.unexported != "" {
		return fmt.Sprintf("uses unexported type %s", q.unexported)
	}
	return ""
}

type qualifier struct {
	pkg   string
	local map[string]bool // Type names declared in the package under test.
	skip  map[string]bool // Type parameter names, which shadow local types.
	// Fields declared together (a, b T) share one Expression; qualify it once.
	done map[*models.Expression]bool
	// The first unexported local type found, which an external test cannot name.
	unexported string
}

func (q *qualifier) field(f *models.Field) {
	if f == nil || f.Type == nil || q.done[f.Type] {
		return
	}
	q.done[f.Type] = true
	e, err := parser.ParseExpr(f.Type.Value)
	if err != nil {
		return
	}
	changed := false
	e = astutil.Apply(e, func(c *astutil.Cursor) bool {
		id, ok := c.Node().(*ast.Ident)
		if !ok || !q.local[id.Name] || q.skip[id.Name] {
			return true
		}
		switch c.Parent().(type) {
		case *ast.SelectorExpr:
			// Already qualified (pkg.Type), or a selector's package name.
			return true
		case *ast.Field:
			// A struct field, method or parameter name rather than a type.
			if c.Name() == "Names" {
				return true
			}
		}
		if !token.IsExported(id.Name) {
			if q.unexported == "" {
				q.unexported = id.Name
			}
			return true
		}
		c.Replace(&ast.SelectorExpr{X: ast.NewIdent(q.pkg), Sel: ast.NewIdent(id.Name)})
		changed = true
		return true
	}, nil).(ast.Expr)
	if changed {
		f.Type.Value = types.ExprString(e)
	}
}

// splitTypeArgs returns the base type name of a receiver expression such as
// Set, Set[T] or Map[K, V], along with the names of its type arguments.
func splitTypeArgs(e ast.Expr) (string, []string) {
	var x ast.Expr
	var idx []ast.Expr
	switch v := e.(type) {
	case *ast.IndexExpr:
		x, idx = v.X, []ast.Expr{v.Index}
	case *ast.IndexListExpr:
		x, idx = v.X, v.Indices
	default:
		x = e
	}
	var args []string
	for _, i := range idx {
		if id, ok := i.(*ast.Ident); ok {
			args = append(args, id.Name)
		}
	}
	if id, ok := x.(*ast.Ident); ok {
		return id.Name, args
	}
	return types.ExprString(x), args
}
