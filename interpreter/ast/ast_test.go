package ast

import (
	"ape/interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar"},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	expected := "let myVar = anotherVar;"
	result := program.String()

	if result != expected {
		t.Errorf("Bad String() result. Expected '%q', got '%q' instead", expected, result)
	}
}
