package ast

import "bee/token"

// Node is implemented by every node
// in the AST. TokenLiteral() is only
// for testing and debugging,
type Node interface {
	TokenLiteral() string
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

// LetStatement represents a `let` variable
// declaration in the AST.
type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral retrieves the literal token (the string representation).
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

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

func (i *Identifier) expressionNode() {}

// TokenLiteral retrieves the literal token (the string representation).
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
