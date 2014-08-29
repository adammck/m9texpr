package ast

import (
	"fmt"
)

type Expression struct {
	*Expression
	Operator
	*Operand
}

func MakeUnaryExpression(o *Operand) (*Expression, error) {
	return &Expression{Operand: o}, nil
}

func MakeExpression(expression *Expression, operand *Operand, operator Operator) (*Expression, error) {
	return &Expression{expression, operator, operand}, nil
}

func (e *Expression) String() string {
	if e.Operator != nil {
		return fmt.Sprintf("e{%s %s %s}", e.Expression, e.Operator, e.Operand)
	} else {
		return fmt.Sprintf("e{%s}", e.Operand)
	}
}

func (e *Expression) Eval(ctx map[string]interface{}) (interface{}, error) {
	if e.Operator == nil {
		return e.Operand.Eval(ctx)

	} else {
		left, err1 := e.Expression.Eval(ctx)
		if(err1 != nil) {
			return nil, err1
		}

		right, err2 := e.Operand.Eval(ctx)
		if(err2 != nil) {
			return nil, err2
		}

		return e.Operator.Compare(left, right)
	}
}
