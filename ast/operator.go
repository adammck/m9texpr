package ast

import (
  "fmt"
  "github.com/adammck/m9texp/gen/token"
)

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
