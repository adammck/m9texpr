package ast

import (
  "fmt"
  "github.com/adammck/m9texp/gen/util"
  "github.com/adammck/m9texp/gen/token"
)

type Int struct {
  val int64
}

func MakeInt(tok *token.Token) (*Int, error) {
  val, err := util.IntValue(tok.Lit)
  if err != nil {
    return nil, err
  }

  return &Int{val}, nil
}

func (i *Int) String() string {
  return fmt.Sprintf("int(%d)", i.val)
}

func (i Int) Truthy() bool {
  return i.val != 0
}
