package ast

import (
	"fmt"
	"strings"
	"github.com/adammck/m9texp/token"
	"github.com/adammck/m9texp/util"
)

type Expression struct {
	Expression *Expression
	Operator   *Operator
	Operand    *Operand
}

func MakeUnaryExpression(o *Operand) (*Expression, error) {
	return &Expression{Operand: o}, nil
}

func MakeBinaryExpression(expression *Expression, operand *Operand, operator *Operator) (*Expression, error) {
	return &Expression{expression, operator, operand}, nil
}

func (e *Expression) String() string {
	if e.Operator != nil {
		return fmt.Sprintf("e{%s %s %s}", e.Expression, e.Operator, e.Operand)
	} else {
		return fmt.Sprintf("e{%s}", e.Operand)
	}
}

type Operand struct {
	Variable      *Variable
	Int           *int64
	Str           *string
	Expression    *Expression
}

func MakeOperand(untype interface{}) (*Operand, error) {
	switch typed := untype.(type) {
	case *Variable:
		return &Operand{Variable: typed}, nil

	case int64:
		return &Operand{Int: &typed}, nil

	case string:
		return &Operand{Str: &typed}, nil

	case *Expression:
		return &Operand{Expression: typed}, nil

	default:
		return nil, fmt.Errorf("invalid operand: %#v", typed)
	}
}

func (o *Operand) String() string {
	var v interface{}

	if o.Variable != nil {
		v = o.Variable

	} else if o.Int != nil {
		v = fmt.Sprintf("int(%d)", *o.Int)

	} else if o.Str != nil {
		v = fmt.Sprintf("str(%s)", *o.Str)

	} else if o.Expression != nil {
		v = o.Expression

	} else {
		panic("invalid operand")
	}

	return fmt.Sprintf("o[%s]", v)
}

type opType uint8

const (
	OpEquals opType = iota
	OpNotEquals
	OpLessThan
	OpGreaterThan
)

type Operator struct {
	Type opType
}

func MakeOperator(tok *token.Token) (*Operator, error) {
	var t opType

	switch string(tok.Lit) {
	case "==":
		t = OpEquals

	case "!=":
		t = OpNotEquals

	case ">":
		t = OpGreaterThan

	case "<":
		t = OpLessThan

	default:
		return nil, fmt.Errorf("invalid operator: %#v", tok.Lit)
	}

	return &Operator{Type: t}, nil
}

func (op *Operator) String() string {
	switch op.Type {
	case OpEquals:
		return "op(eq)"

	case OpNotEquals:
		return "op(ne)"

	case OpGreaterThan:
		return "op(gt)"

	case OpLessThan:
		return "op(lt)"

	default:
		panic("invalid operator")
	}
}

type Variable struct {
	Name string
}

func MakeVariable(tok *token.Token) (*Variable, error) {
	return &Variable{Name: string(tok.Lit)}, nil
}

func (v *Variable) String() string {
	return fmt.Sprintf("var(%s)", v.Name)
}

// IntValue converts a token into an int64.
func IntValue(tok *token.Token) (int64, error) {
	return util.IntValue(tok.Lit)
}

// StringValue converts a token into a string.
func StrValue(tok *token.Token) (string, error) {
	return strings.Trim(string(tok.Lit), "\""), nil
}
