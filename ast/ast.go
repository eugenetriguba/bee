package ast

import (
	"bytes"

	"bee/token"
)

// Node is implemented by every node
// in the AST. TokenLiteral() and String()
// are used for testing and debugging.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement defines whether a node is
// considered a statement in Bee.
type Statement interface {
	Node
	statementNode()
}

// Expression defines whether a node is
// considered a expression in Bee.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST.
type Program struct {
	Statements []Statement
}

// TokenLiteral retrieves the literal token (the string representation)
// of the first statement if we have >= 1 otherwise an empty string.
func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (program *Program) String() string {
	var out bytes.Buffer

	for _, s := range program.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement represents a `let` variable
// declaration in the AST.
type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

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

// ReturnStatement represents a `return <expression>`
// statement.
type ReturnStatement struct {
	Token       token.Token // token.RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()

}

// ExpressionStatement a statement that consists of one expression.
// In that sense, it is not really a statement, more of a wrapper.
// We need it because it is completetly legal to write something like:
// `let x = 5; x + 10;`
type ExpressionStatement struct {
	Token      token.Token // first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// Identifier represents a identifier in Bee
// such as `let x = 5;`, where `x` is the identifier.
//
// Even though an Identifier doesn't produce a value
// in something like a `let x = 5;` statement, we implement
// it as an expression node to keep things simple since it will
// in other parts of a program e.g. `let x = valueProducingIdent`.
type Identifier struct {
	Token token.Token // token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
