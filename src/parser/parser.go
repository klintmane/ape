package parser

import (
	"ape/src/ast"
	"ape/src/lexer"
	"ape/src/token"
	"fmt"
)

type (
	prefixParser func() ast.Expression
	infixParser  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer *lexer.Lexer

	current token.Token
	next    token.Token

	errors []string

	prefixParsers map[token.TokenType]prefixParser
	infixParsers  map[token.TokenType]infixParser
}

// New creates and returns a new Parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{lexer: l, errors: []string{}}

	// Create the prefix parsers map
	p.prefixParsers = make(map[token.TokenType]prefixParser)

	// Append the prefix parsers to the map
	p.addPrefixParser(token.IDENT, p.parseIdentifier)
	p.addPrefixParser(token.INT, p.parseIntegerLiteral)
	p.addPrefixParser(token.BANG, p.parsePrefixExpression)
	p.addPrefixParser(token.MINUS, p.parsePrefixExpression)
	p.addPrefixParser(token.TRUE, p.parseBoolean)
	p.addPrefixParser(token.FALSE, p.parseBoolean)
	p.addPrefixParser(token.PARENL, p.parseGroupedExpression)
	p.addPrefixParser(token.IF, p.parseIfExpression)
	p.addPrefixParser(token.FUNCTION, p.parseFunctionLiteral)
	p.addPrefixParser(token.STRING, p.parseStringLiteral)
	p.addPrefixParser(token.BRACKETL, p.parseArrayLiteral)
	p.addPrefixParser(token.BRACEL, p.parseHashLiteral)

	// Create the infix parsers map
	p.infixParsers = make(map[token.TokenType]infixParser)

	// Append the infix parsers to the map
	p.addInfixParser(token.PLUS, p.parseInfixExpression)
	p.addInfixParser(token.MINUS, p.parseInfixExpression)
	p.addInfixParser(token.SLASH, p.parseInfixExpression)
	p.addInfixParser(token.ASTERISK, p.parseInfixExpression)
	p.addInfixParser(token.EQ, p.parseInfixExpression)
	p.addInfixParser(token.NEQ, p.parseInfixExpression)
	p.addInfixParser(token.LT, p.parseInfixExpression)
	p.addInfixParser(token.GT, p.parseInfixExpression)
	p.addInfixParser(token.PARENL, p.parseCallExpression)
	p.addInfixParser(token.BRACKETL, p.parseIndexExpression)

	// Reads twice so current and next are not nil
	p.advance()
	p.advance()

	return p
}

// Errors returns the parser errors
func (p *Parser) Errors() []string {
	return p.errors
}

// ParseProgram parses a program and returns the AST root
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.current.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.advance()
	}
	return program
}

// Advances the two tokens
func (p *Parser) advance() {
	p.current = p.next
	p.next = p.lexer.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.current.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// If the next token is what is expected, advances the tokens
func (p *Parser) advanceIfNext(t token.TokenType) bool {
	if p.isNext(t) {
		p.advance()
		return true
	}
	p.errorNext(t)
	return false
}

func (p *Parser) isCurrent(t token.TokenType) bool {
	return p.current.Type == t
}

func (p *Parser) isNext(t token.TokenType) bool {
	return p.next.Type == t
}

func (p *Parser) errorNext(t token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be '%s', got '%s' instead", t, p.next.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) addPrefixParser(tokenType token.TokenType, fn prefixParser) {
	p.prefixParsers[tokenType] = fn
}

func (p *Parser) addInfixParser(tokenType token.TokenType, fn infixParser) {
	p.infixParsers[tokenType] = fn
}
