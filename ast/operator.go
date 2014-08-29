package ast

import (
  "fmt"
  "github.com/adammck/m9texp/gen/token"
)

// TMP: This is just a placeholder until expressions are interfaces.
type Expr interface {
  String()
}

type Operator interface {
  Compare(left Expr, right Expr) bool
}

type Equals      struct { }
type NotEquals   struct { }
type GreaterThan struct { }
type LessThan    struct { }

func MakeOperator(tok *token.Token) (Operator, error) {
  v := string(tok.Lit)
  var t Operator

  switch v {
  case "==":
    t = &Equals{}

  case "!=":
    t = &NotEquals{}

  case ">":
    t = &GreaterThan{}

  case "<":
    t = &LessThan{}

  default:
    return nil, fmt.Errorf("invalid operator: %#v", tok.Lit)
  }

  return t, nil
}

// ==

func (o *Equals) Compare(left Expr, right Expr) bool {
  return false
}

func (o *Equals) String() string {
  return "op(eq)"
}

// !=

func (o *NotEquals) Compare(left Expr, right Expr) bool {
  return false
}

func (o *NotEquals) String() string {
  return "op(ne)"
}

// >

func (o *GreaterThan) Compare(left Expr, right Expr) bool {
  return false
}

func (o *GreaterThan) String() string {
  return "op(gt)"
}

// <

func (o *LessThan) Compare(left Expr, right Expr) bool {
  return false
}

func (o *LessThan) String() string {
  return "op(lt)"
}
