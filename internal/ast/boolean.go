package ast

import "inter-median/internal/token"


type Boolean struct {
	Token token.Token
	Value bool
}
func (b *Boolean) express()             {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }
