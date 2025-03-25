package ast

import "github.com/Colossus345/Go-interpreter/internal/token"

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) express()              {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }
