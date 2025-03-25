package ast

import (
	"github.com/Colossus345/Go-interpreter/internal/token"
	"strings"
)


type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (ie *FunctionLiteral) express()             {}
func (ie *FunctionLiteral) TokenLiteral() string { return ie.Token.Literal }
func (ie *FunctionLiteral) String() string {
	var out strings.Builder

	params := []string{}

	for _, p := range ie.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(ie.TokenLiteral())
	out.WriteString("{")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString("}")
	out.WriteString(ie.Body.String())

	return out.String()
}
