package ast

import "github.com/Colossus345/Go-interpreter/internal/token"


type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}



func (es *ExpressionStatement) state()               {}
func (es *ExpressionStatement) express()             {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
