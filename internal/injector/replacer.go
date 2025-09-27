package injector

import (
	"go/parser"
	"go/token"
	"os"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
)

type replacer struct{}

func NewReplacer() *replacer {
	return &replacer{}
}

func (r *replacer) parse(file *os.File) (*dst.File, error) {
	fset := token.NewFileSet()
	f, err := decorator.ParseFile(fset, file.Name(), file, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (r *replacer) Replace(file *os.File) error {
	f, err := r.parse(file)
	if err != nil {
		return err
	}

	r.addDragoImport(f)
	r.replaceGoStatements(f)

	r.replaceChanelStatements(f)
	outputFile, err := os.Create(file.Name() + "_drago.go")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = decorator.Fprint(outputFile, f)
	if err != nil {
		return err
	}

	return nil
}

func (r *replacer) replaceGoStatements(file *dst.File) {
	dstutil.Apply(file, func(cursor *dstutil.Cursor) bool {
		if goStmt, ok := cursor.Node().(*dst.GoStmt); ok {
			dragoCall := &dst.CallExpr{
				Fun: &dst.SelectorExpr{
					X:   &dst.Ident{Name: "drago"},
					Sel: &dst.Ident{Name: "go"},
				},
				Args: []dst.Expr{goStmt.Call},
			}
			cursor.Replace(&dst.ExprStmt{X: dragoCall})
		}
		return true
	}, nil)
}

// Checks import ? return : append drago to import
func (r *replacer) addDragoImport(file *dst.File) {

	for _, imp := range file.Imports {
		if imp.Path.Value == `"drago"` {
			return
		}
	}
	file.Imports = append(file.Imports, &dst.ImportSpec{
		Path: &dst.BasicLit{Value: `"drago"`},
	})
}

func (r *replacer) replaceChanelStatements(file *dst.File) {
	dstutil.Apply(file, func(cursor *dstutil.Cursor) bool {

		if sendStmt, ok := cursor.Node().(*dst.SendStmt); ok {
			dragoCall := &dst.CallExpr{
				Fun: &dst.SelectorExpr{
					X:   &dst.Ident{Name: "drago"},
					Sel: &dst.Ident{Name: "SendChannel"},
				},
				Args: []dst.Expr{sendStmt.Chan, sendStmt.Value},
			}
			cursor.Replace(&dst.ExprStmt{X: dragoCall})
		}

		switch node := cursor.Node().(type) {
		case *dst.AssignStmt:
			if len(node.Rhs) == 1 {
				if unaryExpr, ok := node.Rhs[0].(*dst.UnaryExpr); ok && unaryExpr.Op == token.ARROW {
					dragoCall := &dst.CallExpr{
						Fun: &dst.SelectorExpr{
							X:   &dst.Ident{Name: "drago"},
							Sel: &dst.Ident{Name: "ReceiveChannel"},
						},
						Args: []dst.Expr{unaryExpr.X, createResultIdent(node.Lhs)},
					}
					cursor.Replace(&dst.ExprStmt{X: dragoCall})
				}
			}

		case *dst.ExprStmt:
			if unaryExpr, ok := node.X.(*dst.UnaryExpr); ok && unaryExpr.Op == token.ARROW {
				dragoCall := &dst.CallExpr{
					Fun: &dst.SelectorExpr{
						X:   &dst.Ident{Name: "drago"},
						Sel: &dst.Ident{Name: "ReceiveChannel"},
					},
					Args: []dst.Expr{unaryExpr.X, dst.NewIdent("_")},
				}
				cursor.Replace(&dst.ExprStmt{X: dragoCall})
			}
		}

		return true
	}, nil)
}

func createResultIdent(lhs []dst.Expr) dst.Expr {
	if len(lhs) == 1 {
		return lhs[0]
	}
	return &dst.CompositeLit{
		Type: &dst.ArrayType{Elt: &dst.Ident{Name: "interface{}"}},
		Elts: lhs,
	}
}
