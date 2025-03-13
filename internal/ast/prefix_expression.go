package ast

import (
	"inter-median/internal/token"
	"strings"
)


type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}
func (pe *PrefixExpression) state()               {}
func (pe *PrefixExpression) express()             {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out strings.Builder

	out.WriteByte('(')
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteByte(')')

	return out.String()
}
