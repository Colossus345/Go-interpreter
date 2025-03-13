package ast

import (
	"inter-median/internal/token"
	"strings"
)


type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Right    Expression
	Operator string
}

func (pe *InfixExpression) state()               {}
func (pe *InfixExpression) express()             {}
func (pe *InfixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *InfixExpression) String() string {
	var out strings.Builder

	out.WriteByte('(')
	out.WriteString(pe.Left.String())
	out.WriteByte(' ')
	out.WriteString(pe.Operator)
	out.WriteByte(' ')
	out.WriteString(pe.Right.String())
	out.WriteByte(')')

	return out.String()
}
