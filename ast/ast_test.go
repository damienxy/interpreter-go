package ast

import (
	"testing"

	"github.com/damienxy/interpreter-go/token"
)

func TestString(t *testing.T) {
	expectedOutput := "let myVar = anotherVar;"
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != expectedOutput {
		t.Errorf("program.String() wrong. Got %q instead of %q", program.String(), expectedOutput)
	}
}
