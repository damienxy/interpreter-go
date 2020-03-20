package ast

import (
	"bytes"

	"github.com/damienxy/interpreter-go/token"
)

// Node interface
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement interface
type Statement interface {
	Node
	statementNode()
}

// Expression interface
type Expression interface {
	Node
	expressionNode()
}

// Program is a slice of statements
type Program struct {
	Statements []Statement
}

// Identifier has the structure `<identifier>`
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// IntegerLiteral has the structure `<integer>`
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// LetStatement has the structure `let <identifier> = <expression>;`
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// ReturnStatement has the structure `return <expression>;`
type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

// ExpressionStatement has the structure `<expression>`
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// PrefixExpression has the structure `<prefix operator><expression>;`
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// InfixExpression has the structure `<expression><infix operator><expression>;`
type InfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
	Right    Expression
}

// Program

// TokenLiteral returns the token literal of a program's first statement
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String returns a program's statements as strings
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Identifier

func (i *Identifier) expressionNode() {}

// TokenLiteral returns an identifier's token literal
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String returns an identifier's value as a string
func (i *Identifier) String() string { return i.Value }

// Integer Literal

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns an integer literal's token literal
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String returns an integer literal as a string
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// Let Statement

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns a let statement's token literal
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String returns a let statement's token as a statement string
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Return Statement

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns a return statement's token literal
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns a return statement's token as a statement string
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// Expression Statement

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns an expression statement's token literal
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String returns an expression statement's expression as a string
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Prefix Expression

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns a prefix expression's token literal
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String returns a prefix expression as a parenthesized string
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// Infix Expression

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral returns an infix expression's token literal
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String returns an infix expression as a parenthesized string
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
