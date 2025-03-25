package ast

import "github.com/Colossus345/Go-interpreter/internal/token"


type IntegerLiteral struct {
	Token token.Token
	Value int64
}
func (id *IntegerLiteral) state()               {}
func (id *IntegerLiteral) express()             {}
func (id *IntegerLiteral) TokenLiteral() string { return id.Token.Literal }
func (id *IntegerLiteral) String() string       { return id.Token.Literal }
