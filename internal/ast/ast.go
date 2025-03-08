package ast

import "inter-median/internal/token"

type Node interface {
	TokenLiteral() string
}
type Statement interface {
	Node
	state()
}
type Expression interface {
	Node
	express()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type Identifier struct {
	Token token.Token
	Value string
}
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) state()               {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (id *Identifier) state()               {}
func (id *Identifier) TokenLiteral() string { return id.Token.Literal }
