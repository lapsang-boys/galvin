package main

import (
	"flag"
	"fmt"
	goast "go/ast"
	"go/printer"
	"go/token"
	"log"
	"os"

	tlast "github.com/lapsang-boys/galvin/typeless/ast"
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
	if err := printer.Fprint(os.Stdout, token.NewFileSet(), file); err != nil {
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
	fmtImport := &goast.ImportSpec{
		Path: &goast.BasicLit{Kind: token.STRING, Value: `"fmt"`},
	}
	importDecl := &goast.GenDecl{
		Tok: token.IMPORT,
		Specs: []goast.Spec{
			fmtImport,
		},
	}
	file := &goast.File{
		Name: goast.NewIdent("main"),
		Imports: []*goast.ImportSpec{
			fmtImport,
		},
		Decls: []goast.Decl{
			importDecl,
			mainFunc,
		},
	}
	return &transpiler{
		file: file,
		body: body,
	}
}

func (t *transpiler) transpile(tlFile *tlast.File) (*goast.File, error) {
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

func (t *transpiler) transpileExpr(tlExpr tlast.Expression) (goast.Expr, error) {
	switch tlExpr := tlExpr.(type) {
	case *tlast.Literal:
		// Correct by grammar.
		return &goast.BasicLit{Kind: token.INT, Value: tlExpr.Text()}, nil
	case *tlast.Identifier:
		return goast.NewIdent(tlExpr.Text()), nil
	case *tlast.FunctionApplication:
		v, err := t.transpileExpr(tlExpr.Callee())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		callee, ok := v.(*goast.FuncLit)
		if !ok {
			return nil, errors.Errorf("invalid callee type; expected *ast.FuncLit, got %T", v)
		}
		var args []goast.Expr
		for _, tlExpr := range tlExpr.Arguments() {
			arg, err := t.transpileExpr(tlExpr)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			args = append(args, arg)
		}
		callExpr := &goast.CallExpr{
			Fun:  callee,
			Args: args,
		}
		return callExpr, nil
	case *tlast.FunctionAbstraction:
		params := tlExpr.Parameters()
		body, err := t.transpileExpr(tlExpr.Body().Expression())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		paramNames := make([]*goast.Ident, len(params))
		for i, param := range params {
			paramNames[i] = goast.NewIdent(param.Text())
		}
		emptyInterfaceType := &goast.InterfaceType{
			Methods: &goast.FieldList{},
		}
		goParams := &goast.FieldList{
			List: []*goast.Field{
				{
					Names: paramNames,
					Type:  emptyInterfaceType,
				},
			},
		}
		goResult := &goast.FieldList{
			List: []*goast.Field{
				{
					Type: emptyInterfaceType,
				},
			},
		}
		return &goast.FuncLit{
			Type: &goast.FuncType{
				Params:  goParams,
				Results: goResult,
			},
			Body: &goast.BlockStmt{
				List: []goast.Stmt{&goast.ReturnStmt{Results: []goast.Expr{body}}},
			},
		}, nil
	default:
		panic(fmt.Errorf("support for expression %T not yet implemented", tlExpr))
	}
}
