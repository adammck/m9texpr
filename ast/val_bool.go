package ast

import (
)

type Bool struct { }

var (
  True = Bool{}
  False = Bool{}
)

func MakeBool(b bool) Bool {
  if b {
    return True
  } else {
    return False
  }
}

func (b Bool) String() string {
  if b == True {
    return "t"
  } else {
    return "f"
  }
}


func (b Bool) Eval(ctx map[string]interface{}) (Expr, error) {
  return MakeBool(b == True), nil
}

func (b Bool) Truthy() bool {
  return b == True
}
