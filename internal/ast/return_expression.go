package ast

import (
	"inter-median/internal/token"
	"strings"
)

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (ls *ReturnStatement) state()               {}
func (ls *ReturnStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *ReturnStatement) String() string {
	var out strings.Builder
	out.WriteString(ls.TokenLiteral())
	out.WriteString(" ")
	if ls.ReturnValue != nil {
		out.WriteString(ls.ReturnValue.String())
	}
	out.WriteRune(';')

	return out.String()
}
