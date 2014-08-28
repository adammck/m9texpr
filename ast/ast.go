package ast

import (
	"fmt"
	"github.com/adammck/m9texp/token"
	"github.com/adammck/m9texp/util"
	"strings"
)

type Expression struct {

	// single value
	Operand *Operand

	// or more recursive exprs
	Left     *Expression
	Right    *Operand
	Operator *Operator
}

func MakeUnaryExpression(o *Operand) (*Expression, error) {
	return &Expression{Operand: o}, nil
}

func MakeBinaryExpression(left *Expression, right *Operand, op *Operator) (*Expression, error) {
	return &Expression{Operator: op, Left: left, Right: right}, nil
}

func (e *Expression) String() string {
	if e.Operator != nil {
		return fmt.Sprintf("e{%s %s %s}", e.Left, e.Operator, e.Right)
	} else {
		return fmt.Sprintf("e{%s}", e.Operand)
	}
}

type Operand struct {
	Variable      *Variable
	Integer       *Integer
	StringLiteral *StringLiteral
	Expression    *Expression
}

// TODO: Maybe move the typecasting in the bnf?
func MakeOperand(untype interface{}) (*Operand, error) {
	switch typed := untype.(type) {
	case *Variable:
		return &Operand{Variable: typed}, nil

	case *Integer:
		return &Operand{Integer: typed}, nil

	case *StringLiteral:
		return &Operand{StringLiteral: typed}, nil

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

	} else if o.Integer != nil {
		v = o.Integer

	} else if o.StringLiteral != nil {
		v = o.StringLiteral

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

func NewOperator(tok *token.Token) (*Operator, error) {
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

func NewVariable(val string) (*Variable, error) {
	return &Variable{val}, nil
}

func (v *Variable) String() string {
	return fmt.Sprintf("var(%s)", v.Name)
}

type Integer struct {
	Val int64
}

func NewInteger(lit []byte) (*Integer, error) {
	i, err := util.IntValue(lit)
	if err != nil {
		return nil, err
	}

	return &Integer{i}, nil
}

func (i *Integer) String() string {
	return fmt.Sprintf("int(%d)", i.Val)
}

type StringLiteral struct {
	Str string
}

func MakeStringLiteral(tok *token.Token) (*StringLiteral, error) {
	return &StringLiteral{strings.Trim(string(tok.Lit), "\"")}, nil
}

func (s *StringLiteral) String() string {
	return fmt.Sprintf("str(%s)", s.Str)
}
