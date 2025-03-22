package ast

import (
	"inter-median/internal/token"
	"strings"
)

type AssignExpression struct {
	Token token.Token
	Left  *Identifier
	Right Expression
}

func (ce *AssignExpression) express()             {}
func (ce *AssignExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *AssignExpression) String() string {
	var out strings.Builder


	out.WriteString(ce.Left.String())
	out.WriteString("=")
	out.WriteString(ce.Right.String())

	return out.String()
}
