package ast

import (
	"fmt"
)

type Expr interface {
	Eval(ctx map[string]interface{}) (Expr, error)
	Truthy() (bool)
}

type Expression struct {
	Left  Expr
	Right Expr
	Op    Operator
}

func MakeUnaryExpression(o *Operand) (*Expression, error) {
	return &Expression{Left: o}, nil
}

func MakeExpression(expression Expr, operand Expr, operator Operator) (*Expression, error) {
	return &Expression{Left: expression, Right: operand, Op: operator}, nil
}

func (e *Expression) String() string {
	if e.Op != nil {
		return fmt.Sprintf("e{%s %s %s}", e.Left, e.Op, e.Left)
	} else {
		return fmt.Sprintf("e{%s}", e.Right)
	}
}

func (e *Expression) Eval(ctx map[string]interface{}) (Expr, error) {
	if e.Op == nil {
		return e.Right.Eval(ctx)

	} else {
		left, err1 := e.Left.Eval(ctx)
		if(err1 != nil) {
			return nil, err1
		}

		right, err2 := e.Right.Eval(ctx)
		if(err2 != nil) {
			return nil, err2
		}

		return e.Op.Compare(left, right)
	}
}

func (e *Expression) Truthy() bool {
	panic("cant truthy an expression")
}
