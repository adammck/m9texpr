package ast

import (
  "fmt"
)

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

func (o *Operand) Eval(ctx map[string]interface{}) bool {
  if o.Variable != nil {
    return o.Variable.Eval(ctx)

  } else {
    panic("not implemented!")
  }
}
