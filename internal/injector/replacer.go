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

	r.replaceGoStatements(f)
	
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
