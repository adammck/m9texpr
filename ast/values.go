package ast

import (
  "fmt"
  "strings"
  "github.com/adammck/m9texp/token"
  "github.com/adammck/m9texp/util"
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

func (v *Variable) Eval(ctx map[string]interface{}) bool {
  vv, ok := ctx[v.Name]
  if !ok {
    return false
  }

  switch vvv := vv.(type) {
    case bool:
      return vvv
    default:
      panic("only bool vars are implemented")
  }
}

// IntValue converts a token into an int64.
func IntValue(tok *token.Token) (int64, error) {
  return util.IntValue(tok.Lit)
}

// StringValue converts a token into a string.
func StrValue(tok *token.Token) (string, error) {
  return strings.Trim(string(tok.Lit), "\""), nil
}
