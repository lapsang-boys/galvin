package main

import (
	"flag"
	"fmt"
	goast "go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"

	"github.com/lapsang-boys/galvin/typeless/ast"
	"github.com/lapsang-boys/galvin/typeless/parser"
	"github.com/pkg/errors"
)

func usage() {
	const use = `
Tal FILE.tl`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	// Parse command line arguments.
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	tlPath := flag.Arg(0)

	// Transpile from Typeless to Go.
	t := newTranspiler()
	tlFile, err := parser.ParseFile(tlPath)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	file, err := t.transpile(tlFile)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	//pretty.Println(file)
	if err := printer.Fprint(os.Stdin, token.NewFileSet(), file); err != nil {
		log.Fatalf("%+v", err)
	}
}

type transpiler struct {
	file *goast.File
	body *goast.BlockStmt
}

func newTranspiler() *transpiler {
	body := &goast.BlockStmt{}
	mainFunc := &goast.FuncDecl{
		Name: goast.NewIdent("main"),
		Type: &goast.FuncType{},
		Body: body,
	}
	file := &goast.File{
		Name: goast.NewIdent("main"),
		Imports: []*goast.ImportSpec{
			{Name: goast.NewIdent("fmt")},
		},
		Decls: []goast.Decl{
			mainFunc,
		},
	}
	return &transpiler{
		file: file,
		body: body,
	}
}

func (t *transpiler) transpile(tlFile *ast.File) (*goast.File, error) {
	for _, tlExpr := range tlFile.Expressions() {
		expr, err := t.transpileExpr(tlExpr)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		printf := &goast.SelectorExpr{
			X:   goast.NewIdent("fmt"),
			Sel: goast.NewIdent("Printf"),
		}
		format := &goast.BasicLit{Kind: token.STRING, Value: `"%#v\n"`}
		args := []goast.Expr{format, expr}
		callExpr := &goast.CallExpr{
			Fun:  printf,
			Args: args,
		}
		callStmt := &goast.ExprStmt{
			X: callExpr,
		}
		t.body.List = append(t.body.List, callStmt)
	}
	return t.file, nil
}

func (t *transpiler) transpileExpr(tlExpr ast.Expression) (goast.Expr, error) {
	switch tlExpr := tlExpr.(type) {
	case *ast.Literal:
		// Correct by grammar.
		return &goast.BasicLit{Kind: token.INT, Value: tlExpr.Text()}, nil
	default:
		panic(fmt.Errorf("support for expression %T not yet implemented", tlExpr))
	}
}
