package main

import (
	"fmt"
	"log"

	"github.com/lapsang-boys/galvin/typeless/ast"
)

func main() {
	tree, err := ParseFile("foo.tless")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Printf("%#v\n", tree)

	interpret(tree)

	fmt.Println(tree.Text())
}

func interpret(f *ast.File) {
	empty := ClosureEnvironment{
		i2x: make(map[string]ast.Expression),
	}
	for _, expr := range f.Expressions() {
		res := empty.interpretExpression(expr)
		fmt.Printf("%#v\n", res.TypelessNode().Text())

		// val, _ = strconv.ParseInt(e.Text(), 10, 64)
	}
}

// ClosureEnvironment asdf
type ClosureEnvironment struct {
	i2x map[string]ast.Expression
}

func (ce ClosureEnvironment) interpretExpression(expr ast.Expression) ast.Expression {
	switch e := expr.(type) {
	case *ast.FunctionApplication:
		return ce.interpretApplication(e)
	case *ast.FunctionAbstraction:
		return e
	case *ast.Literal:
		return e
	case *ast.Identifier:
		v, ok := ce.i2x[e.Text()]
		if !ok {
			fmt.Println("offset:", e.Offset())
			panic(fmt.Sprintf("Unable to lookup identifer: %v", e.Text()))
		}
		return v
	default:
		panic(fmt.Sprintf("unexpected token. %T,", e))
	}
}

func (ce ClosureEnvironment) interpretApplication(app *ast.FunctionApplication) ast.Expression {
	v := ce.interpretExpression(app.Callee())
	var callee *ast.FunctionAbstraction
	var ok bool
	if callee, ok = v.(*ast.FunctionAbstraction); !ok {
		panic("Callee must be FunctionAbstraction in FunctionApplication.")
	}
	params := callee.Parameters()
	args := app.Arguments()

	if len(params) != len(args) {
		panic("FunctionAbstractions number of parameters not equal to FunctionApplications number of arguments.")
	}

	for i, param := range params {
		// TODO(_): alpha-reduction
		ce.i2x[param.Text()] = ce.interpretExpression(args[i])
	}
	return ce.interpretExpression(callee.Body().Expression())
}
