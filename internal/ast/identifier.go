package ast

import "github.com/Colossus345/Go-interpreter/internal/token"

type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) state()               {}
func (id *Identifier) express()             {}
func (id *Identifier) TokenLiteral() string { return id.Token.Literal }
func (id *Identifier) String() string       { return id.Value }
