package ast

import (
	"inter-median/internal/token"
	"strings"
)


func (bl *BlockStatement) state() {}
func (bl *BlockStatement) TokenLiteral() string {
	return bl.Token.Literal
}
func (bl *BlockStatement) String() string {
	var out strings.Builder

	for _, s := range bl.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}
