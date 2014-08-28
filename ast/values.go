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

// IntValue converts a token into an int64.
func IntValue(tok *token.Token) (int64, error) {
  return util.IntValue(tok.Lit)
}

// StringValue converts a token into a string.
func StrValue(tok *token.Token) (string, error) {
  return strings.Trim(string(tok.Lit), "\""), nil
}
