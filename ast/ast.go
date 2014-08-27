package ast

import (
	"fmt"
	"strings"
	"github.com/adammck/m9texp/token"
	"github.com/adammck/m9texp/util"
	//"github.com/adammck/m9texp/parser"
)

type Expr struct {

	// single value
	Atom *Atom

	// or more recursive exprs
	Left     *Expr
	Right    *Atom
	Operator *Operator
}

func MakeUnaryExpr(atom *Atom) (*Expr, error) {
	return &Expr{Atom: atom}, nil
}

func MakeBinaryExpr(left *Expr, right *Atom, op *Operator) (*Expr, error) {
	return &Expr{Operator: op, Left: left, Right: right}, nil
}

func (expr *Expr) String() string {
	if expr.Atom != nil {
		return fmt.Sprintf("expr{%s}", expr.Atom)
	} else {
		return fmt.Sprintf("expr{%s %s %s}", expr.Left, expr.Operator, expr.Right)
	}
}




type Atom struct {
	Variable      *Variable
	Integer       *Integer
	StringLiteral *StringLiteral
	Expr          *Expr
}

// TODO: Maybe move the typecasting in the bnf?
func MakeAtom(untype interface{}) (*Atom, error) {
	switch typed := untype.(type) {
	case *Variable:
		return &Atom{Variable: typed}, nil

	case *Integer:
		return &Atom{Integer: typed}, nil

	case *StringLiteral:
		return &Atom{StringLiteral: typed}, nil

	case *Expr:
		return &Atom{Expr: typed}, nil

	default:
		return nil, fmt.Errorf("invalid atom: %#v", typed)
	}
}

func (atom *Atom) String() string {
	var v interface{}

	if atom.Variable != nil {
		v = atom.Variable

	} else if atom.Integer != nil {
		v = atom.Integer

	} else if atom.StringLiteral != nil {
		v = atom.StringLiteral

	} else if atom.Expr != nil {
		v = atom.Expr

	} else {
		panic("invalid atom")
	}

	return fmt.Sprintf("atom[%s]", v)
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
