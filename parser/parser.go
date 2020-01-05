package parser

import (
	"bee/ast"
	"bee/lexer"
	"bee/token"
)

// Parser is used to parse the program and
// construct the AST.
type Parser struct {
	lexer *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

// New creates a new Parser by settings its
// lexer and advancing the positions until
// currentToken points to the first token in
// the program and peekToken points to the second.
func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer}

	// Read two tokens so currentToken and peekToken are set.
	parser.nextToken()
	parser.nextToken()

	return parser
}

// nextToken sets the current token to the
// peekToken and advances the peekToken.
func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

// ParseProgram parses the program by running through each
// line until we run into the EOF token. Each statement is added
// to the AST.
func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !parser.currentTokenIs(token.EOF) {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parser.nextToken()
	}

	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
	}
}

// parseLetStatement parses a let statement by making assertions
// on what the next tokens should be (start with 'let' then '=',
// and so on) and uses that information to create a node on the ast.
func (parser *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: parser.currentToken}

	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: parser.currentToken,
		Value: parser.currentToken.Literal,
	}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return stmt
}

// currentTokenIs is a helper method to check the current token type.
func (parser *Parser) currentTokenIs(tok token.TokenType) bool {
	return parser.currentToken.Type == tok
}

// peekTokenIs is a helper method to check the peek token type.
func (parser *Parser) peekTokenIs(tok token.TokenType) bool {
	return parser.peekToken.Type == tok
}

// expectPeek is used to advance to the next token only if we
// have a certain token type.
func (parser *Parser) expectPeek(tok token.TokenType) bool {
	if parser.peekTokenIs(tok) {
		parser.nextToken()
		return true
	} else {
		return false
	}
}
