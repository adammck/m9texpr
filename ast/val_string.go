package ast

import (
  "fmt"
  "strings"
  "github.com/adammck/m9texp/gen/token"
)

type Str struct {
  val string
}

func MakeStr(tok *token.Token) (*Str, error) {
  return &Str{strings.Trim(string(tok.Lit), "\"")}, nil
}

func (s *Str) String() string {
  return fmt.Sprintf("%s", s.val)
}

func (s Str) Truthy() bool {
  return s.val != ""
}
