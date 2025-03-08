package parser

import (
	"inter-median/internal/ast"
	"inter-median/internal/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
    let x=5;
    let y = 10;
    let foobar = 83;
    `
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram return nil")

	}
	if len(program.Statements) != 3 {
		t.Fatalf("program does not contain 3 statements got=%d",
			len(program.Statements))
	}
	tests := []struct{ expected string }{{"x"}, {"y"}, {"foobar"}}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetState(t, stmt, tt.expected) {
			return
		}
	}
}

func testLetState(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not let got %q", s.TokenLiteral())
		return false
	}

	letState, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement got =%T", s)
		return false
	}
	if letState.Name.Value != name {
		t.Errorf("s not *ast.LetStatement %q got =%T", name, letState.Name.Value)
		return false
	}
	if letState.Name.TokenLiteral() != name {
		t.Errorf("s.name not %q got %s", name, letState.Name)
		return false
	}
	return true
}
