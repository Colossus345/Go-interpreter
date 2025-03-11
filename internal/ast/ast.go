package ast

import (
	"inter-median/internal/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out strings.Builder
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}
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

func (ie *IfExpression) express()             {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out strings.Builder

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}
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

type Identifier struct {
	Token token.Token
	Value string
}
type IntegerLiteral struct {
	Token token.Token
	Value int64
}
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Right    Expression
	Operator string
}

type Boolean struct {
	Token token.Token
	Value bool
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

func (ls *LetStatement) state()               {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {

	var out strings.Builder

	out.WriteString(ls.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteRune(';')

	return out.String()
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
func (id *Identifier) state()               {}
func (id *Identifier) express()             {}
func (id *Identifier) TokenLiteral() string { return id.Token.Literal }
func (id *Identifier) String() string       { return id.Value }

func (id *IntegerLiteral) state()               {}
func (id *IntegerLiteral) express()             {}
func (id *IntegerLiteral) TokenLiteral() string { return id.Token.Literal }
func (id *IntegerLiteral) String() string       { return id.Token.Literal }

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
func (b *Boolean) express()             {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }
