package parser

import (
	"inter-median/internal/ast"
	"inter-median/internal/lexer"
	"testing"
)

func TestInteger(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has wrong statement quantity  got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement got %T",
			stmt)
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp is not a integer got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Fatalf("exp is not 5 got %d", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Fatalf("wrong token literal: expected 5 got=%s",
			literal.TokenLiteral())
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough Statements got %d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement got %T",
			stmt)
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	println(ident, ok, "nu")
	if !ok {
		t.Fatalf("exp not *ast.Identifier got %T", ident)
	}
	if ident.Value != "foobar" {
		t.Fatalf("ident.value is not foobar got=%s", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral() is not foobar got=%s", ident.Value)
	}

}

func TestReturnStatement(t *testing.T) {
	input := `
    return 5;
    return  10;
    return  83;
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
	for _, stmt := range program.Statements {
		rS, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got %T", stmt)
			continue
		}
		if rS.TokenLiteral() != "return" {
			t.Errorf("wrong TokenLiteral expected 'return' got %q", rS.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error %s", msg)
	}
	t.FailNow()
}
func TestLetStatement(t *testing.T) {
	input := `
    let x = 5;
    let y = 10;
    let foobar = 83;
    `
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)
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
