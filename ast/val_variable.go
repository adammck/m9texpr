package ast

import (
  "fmt"
  "github.com/adammck/m9texp/gen/token"
)

type Variable struct {
  Name string
}

func MakeVariable(tok *token.Token) (*Variable, error) {
  return &Variable{Name: string(tok.Lit)}, nil
}

func (v *Variable) String() string {
  return fmt.Sprintf("var(%s)", v.Name)
}

func (v *Variable) Eval(ctx map[string]interface{}) (Expr, error) {
  vv, ok := ctx[v.Name]
  if !ok {
    return nil, fmt.Errorf("undefined: %s", v.Name)
  }

  switch vvv := vv.(type) {
    case bool:
      return MakeBool(vvv), nil
    default:
      panic(fmt.Sprintf("only bool vars are implemented, not %T", vv))
  }
}

func (v Variable) Truthy() bool {
  return v.Truthy()
}
