package test

import (
	"github.com/adammck/m9texp/ast"
	"github.com/adammck/m9texp/gen/lexer"
	"github.com/adammck/m9texp/gen/parser"
	"testing"
)

type TE struct {
	src      string
	expected string
}

var testExprs = []*TE{
	{"1", "e{o[int(1)]}"},
	{"a", "e{o[var(a)]}"},
	{"\"b\"", "e{o[str(b)]}"},
	{"a == 1", "e{e{o[var(a)]} op(eq) o[int(1)]}"},
	{"a > b < c", "e{e{e{o[var(a)]} op(gt) o[var(b)]} op(lt) o[var(c)]}"},
	{"(a > b) < c", "e{e{o[e{e{o[var(a)]} op(gt) o[var(b)]}]} op(lt) o[var(c)]}"},
	{"a > (b < c)", "e{e{o[var(a)]} op(gt) o[e{e{o[var(b)]} op(lt) o[var(c)]}]}"},
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
			actual := expr.(*ast.Expression).String()

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
