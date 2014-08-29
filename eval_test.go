package test

import (
	"github.com/adammck/m9texp/ast"
	"testing"
)

type ctx map[string]interface{}

type evalExample struct {
	expr     *ast.Expression
	context  ctx
	expected bool
}

var (
	// a
	aExpr = &ast.Expression{Right: &ast.Operand{Variable: &ast.Variable{Name: "a"}}}

	// b == c
	bcExpr = &ast.Expression{
		Left:  &ast.Expression{Right: &ast.Operand{Variable: &ast.Variable{Name: "b"}}},
		Right: &ast.Operand{Variable: &ast.Variable{Name: "c"}},
		Op:    &ast.Equals{},
	}
)

var evalExamples = []*evalExample{
	{aExpr, ctx{"a": true}, true},               // 1
	{aExpr, ctx{"a": false}, false},             // 2
	{bcExpr, ctx{"b": true, "c": true}, true},   // 3
	{bcExpr, ctx{"b": false, "c": false}, true}, // 4
	{bcExpr, ctx{"b": true, "c": false}, false}, // 5
	{bcExpr, ctx{"b": false, "c": true}, false}, // 6
}

func TestEval(t *testing.T) {
	pass := true

	for i, eg := range evalExamples {
		actual, err := eg.expr.Eval(eg.context)

		if err != nil {
			t.Errorf("Example #%d: %s", i, err.Error())
			pass = false

		} else if actual.Truthy() != eg.expected {
			t.Errorf("Example #%d: got %v, expected %v", i, actual, eg.expected)
			pass = false
		}
	}

	if !pass {
		t.Fail()
	}
}
