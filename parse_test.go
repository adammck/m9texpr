package test

import (
	"github.com/adammck/m9texp/ast"
	"github.com/adammck/m9texp/lexer"
	"github.com/adammck/m9texp/parser"
	"testing"
)

type TE struct {
	src      string
	expected string
}

var testExprs = []*TE{
	{"1",           "expr{atom[int(1)]}"},
	{"a",           "expr{atom[var(a)]}"},
	{"\"b\"",       "expr{atom[str(b)]}"},
	{"a == 1",      "expr{expr{atom[var(a)]} op(eq) atom[int(1)]}"},
	{"a > b < c",   "expr{expr{expr{atom[var(a)]} op(gt) atom[var(b)]} op(lt) atom[var(c)]}"},
	{"a > (b < c)", "expr{expr{atom[var(a)]} op(gt) atom[expr{expr{atom[var(b)]} op(lt) atom[var(c)]}]}"},
}

func TestParser(t *testing.T) {
	p := parser.NewParser()
	fail := false

	for _, ts := range testExprs {
		s := lexer.NewLexer([]byte(ts.src))
		expr, err := p.Parse(s)

		// parser error?
		if err != nil {
			fail = true
			t.Log(err.Error())
		}

		// invalid output?
		if expr != nil {
			actual := expr.(*ast.Expr).String()

			if actual != ts.expected {
				t.Errorf("\nsrc:    %#v\ngot:    %s\nwanted: %v", ts.src, actual, ts.expected)
				fail = true
			}
		}
	}

	if fail {
		t.Fail()
	}
}
