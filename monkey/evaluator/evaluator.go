package evaluator

import (
	"github.com/istsh/go-writing-an-interpreter/monkey/ast"
	"github.com/istsh/go-writing-an-interpreter/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}