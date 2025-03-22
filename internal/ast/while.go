package ast

import (
	"inter-median/internal/token"
	"strings"
)

type WhileExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}
func (ie *WhileExpression) express()             {}
func (ie *WhileExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *WhileExpression) String() string {
	var out strings.Builder

	out.WriteString("while")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	return out.String()
}
