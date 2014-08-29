package ast

import (
	"fmt"
)

type Expression struct {
	Expression *Expression
	Operator   Operator
	Operand    *Operand
}

func MakeUnaryExpression(o *Operand) (*Expression, error) {
	return &Expression{Operand: o}, nil
}

func MakeBinaryExpression(expression *Expression, operand *Operand, operator Operator) (*Expression, error) {
	return &Expression{expression, operator, operand}, nil
}

func (e *Expression) String() string {
	if e.Operator != nil {
		return fmt.Sprintf("e{%s %s %s}", e.Expression, e.Operator, e.Operand)
	} else {
		return fmt.Sprintf("e{%s}", e.Operand)
	}
}

func (e *Expression) Eval(ctx map[string]interface{}) bool {
	if e.Operator == nil {
		return e.Operand.Eval(ctx)

	} else {
		panic("not implemented")
	}
}
