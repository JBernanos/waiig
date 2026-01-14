package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch, false)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch, false)
	case '-':
		tok = newToken(token.MINUS, l.ch, false)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch, false)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch, false)
	case '/':
		tok = newToken(token.SLASH, l.ch, false)
	case '<':
		tok = newToken(token.LT, l.ch, false)
	case '>':
		tok = newToken(token.GT, l.ch, false)
	case ',':
		tok = newToken(token.COMMA, l.ch, false)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, false)
	case '(':
		tok = newToken(token.LPAREN, l.ch, false)
	case ')':
		tok = newToken(token.RPAREN, l.ch, false)
	case '{':
		tok = newToken(token.LBRACE, l.ch, false)
	case '}':
		tok = newToken(token.RBRACE, l.ch, false)
	case 0:
		tok = newToken(token.EOF, l.ch, true)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, false)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte, isEOF bool) token.Token {
	if isEOF {
		return token.Token{Type: tokenType, Literal: ""}
	} else {
		return token.Token{Type: tokenType, Literal: string(ch)}
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
