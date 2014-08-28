package test

import (
	"github.com/adammck/m9texp/ast"
	"testing"
)

type evalExample struct {
	expr     *ast.Expression
	context  map[string]interface{}
	expected bool
}

var (
	// a
	aExpr = &ast.Expression{Operand: &ast.Operand{Variable: &ast.Variable{Name: "a"}}}

	// b == c
	bcExpr = &ast.Expression{
		Expression: &ast.Expression{Operand: &ast.Operand{Variable: &ast.Variable{Name: "a"}}},
		Operator:   &ast.Operator{Type: ast.OpEquals},
		Operand:    &ast.Operand{Variable: &ast.Variable{Name: "b"}},
	}
)

var evalExamples = []*evalExample{
	{aExpr, map[string]interface{}{"a": true}, true},
	{aExpr, map[string]interface{}{"a": false}, false},
	{bcExpr, map[string]interface{}{"b": true, "c": true}, true},
	{bcExpr, map[string]interface{}{"b": false, "c": false}, true},
	{bcExpr, map[string]interface{}{"b": true, "c": false}, false},
	{bcExpr, map[string]interface{}{"b": false, "c": true}, false},
}

func TestEval(t *testing.T) {
	pass := true

	for i, eg := range evalExamples {
		actual := eg.expr.Eval(eg.context)
		if actual != eg.expected {
			t.Errorf("Example #%d => %v, expected %v", i, actual, eg.expected)
			pass = false
		}
	}

	if !pass {
		t.Fail()
	}
}
