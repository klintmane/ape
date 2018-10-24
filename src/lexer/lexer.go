package lexer

import (
	"ape/src/token"
)

type Lexer struct {
	input        string // The input string
	char         byte   // The current character
	charPosition int    // The position of the current character
	position     int    // The position of the cursor
}

// New Lexer creates and returns a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.advanceChar()
	return l
}

// NextToken scans the input and returns a new token from it
func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.char {
	case '+':
		t = token.New(token.PLUS, l.char)
	case '-':
		t = token.New(token.MINUS, l.char)
	case '*':
		t = token.New(token.ASTERISK, l.char)
	case '/':
		t = token.New(token.SLASH, l.char)

	case '<':
		t = token.New(token.LT, l.char)
	case '>':
		t = token.New(token.GT, l.char)

	case '=':
		if l.peekChar() == '=' {
			char := l.char
			l.advanceChar()
			t = token.Token{Type: token.EQ, Literal: string(char) + string(l.char)}
		} else {
			t = token.New(token.ASSIGN, l.char)
		}

	case '!':
		if l.peekChar() == '=' {
			char := l.char
			l.advanceChar()
			t = token.Token{Type: token.NEQ, Literal: string(char) + string(l.char)}
		} else {
			t = token.New(token.BANG, l.char)
		}

	case '(':
		t = token.New(token.PARENL, l.char)
	case ')':
		t = token.New(token.PARENR, l.char)
	case '{':
		t = token.New(token.BRACEL, l.char)
	case '}':
		t = token.New(token.BRACER, l.char)
	case '[':
		t = token.New(token.BRACKETL, l.char)
	case ']':
		t = token.New(token.BRACKETR, l.char)
	case ',':
		t = token.New(token.COMMA, l.char)
	case ':':
		t = token.New(token.COLON, l.char)
	case ';':
		t = token.New(token.SEMICOLON, l.char)

	case '"':
		t.Type = token.STRING
		t.Literal = l.readString()

	case 0:
		t.Literal = ""
		t.Type = token.EOF

	default:
		if isLetter(l.char) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if isDigit(l.char) {
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		} else {
			t = token.New(token.ILLEGAL, l.char)
		}
	}

	l.advanceChar()
	return t
}

// Sets the character and advances the positions
func (l *Lexer) advanceChar() {
	l.char = l.peekChar()
	l.charPosition = l.position
	l.position++
}

func (l *Lexer) peekChar() byte {
	if l.position < len(l.input) {
		return l.input[l.position]
	}
	return 0
}

func (l *Lexer) readWhile(predicate func(byte) bool) string {
	pos := l.charPosition
	for predicate(l.char) {
		l.advanceChar()
	}
	return l.input[pos:l.charPosition]
}

func (l *Lexer) readIdentifier() string {
	return l.readWhile(isLetter)
}

func (l *Lexer) readNumber() string {
	return l.readWhile(isDigit)
}

func (l *Lexer) readString() string {
	l.advanceChar() // Skips the first double quotes
	return l.readWhile(whileString)
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.advanceChar()
	}
}
