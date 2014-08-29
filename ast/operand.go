package ast

import (
  "fmt"
)

type Operand struct {
  Variable      *Variable
  Int           *Int
  Str           *Str
  Expression    *Expression
}

func MakeOperand(untype interface{}) (*Operand, error) {
  switch typed := untype.(type) {
  case *Variable:
    return &Operand{Variable: typed}, nil

  case *Int:
    return &Operand{Int: typed}, nil

  case *Str:
    return &Operand{Str: typed}, nil

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
    v = o.Int

  } else if o.Str != nil {
    v = o.Str

  } else if o.Expression != nil {
    v = o.Expression

  } else {
    panic("invalid operand")
  }

  return fmt.Sprintf("o[%s]", v)
}

func (o *Operand) Eval(ctx map[string]interface{}) (Expr, error) {
  if o.Variable != nil {
    return o.Variable.Eval(ctx)

  } else {
    panic("not implemented!")
  }
}

func (o *Operand) Truthy() bool {
  if o.Variable != nil {
    return o.Variable.Truthy()

  } else if o.Int != nil {
    return o.Int.Truthy()

  } else if o.Str != nil {
    return o.Str.Truthy()

  } else if o.Expression != nil {
    return o.Expression.Truthy()

  } else {
    panic("invalid operand")
  }
}
