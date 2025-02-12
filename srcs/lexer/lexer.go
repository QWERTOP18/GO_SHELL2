package lexer

import (
	"shell/token"
	"strings"
)

type Lexer struct {
	input        []rune
	position     int
	readPosition int
	r            rune
}

func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readRune()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.consumeWhitespace()
	switch l.r {
	case '|':
		if l.peekRune() == '|' {
			r := l.r
			l.readRune()
			literal := string(r) + string(l.r)
			tok = token.Token{Type: token.PIPE, Literal: literal}
		} else {
			tok = newToken(token.OR, l.r)
		}
	case '&':
		if l.peekRune() == '&' {
			r := l.r
			l.readRune()
			literal := string(r) + string(l.r)
			tok = token.Token{Type: token.AND, Literal: literal}
		} else {
			tok = newToken(token.ASYNC, l.r)
		}
	case '<':
		if l.peekRune() == '<' {
			r := l.r
			l.readRune()
			literal := string(r) + string(l.r)
			tok = token.Token{Type: token.REDIRECT, Literal: literal}
		} else {
			tok = newToken(token.REDIRECT, l.r)
		}
	case '>':
		if l.peekRune() == '>' {
			r := l.r
			l.readRune()
			literal := string(r) + string(l.r)
			tok = token.Token{Type: token.REDIRECT, Literal: literal}
		} else {
			tok = newToken(token.REDIRECT, l.r)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.r)
	case '{':
		tok = newToken(token.LBRACE, l.r)
	case '}':
		tok = newToken(token.RBRACE, l.r)
	case '(':
		tok = newToken(token.LPAREN, l.r)
	case ')':
		tok = newToken(token.RPAREN, l.r)
	case '"', '\'':
		tok.Type = token.WORD
		tok.Literal = l.readString(l.r)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if !isMetachar(l.r) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.WORD
			return tok
		} else if isDigit(l.r) {
			tok.Type = token.NUMBER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.r)
		}
	}

	l.readRune()
	return tok
}

func (l *Lexer) consumeWhitespace() {
	for l.r == ' ' || l.r == '\t' || l.r == '\n' || l.r == '\r' {
		l.readRune()
	}
}

func (l *Lexer) readRune() {
	if l.readPosition >= len(l.input) {
		l.r = 0
	} else {
		l.r = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekRune() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.r != 0 && !isMetachar(l.r) { // EOFチェックを追加
		l.readRune()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.r) {
		l.readRune()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readString(quote rune) string {
	position := l.position
	for {
		l.readRune()
		if l.r == quote || l.r == 0 {
			break
		}
	}
	if l.r == 0 {
		panic("Unterminated string literal")
	}
	return string(l.input[position : l.position+1])
}

func isLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' || r == '_'
}

const METACHARS = "|&;()<> \t"

func isMetachar(r rune) bool {
	return strings.ContainsRune(METACHARS, r)
}

func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func newToken(tokenType token.TokenType, r rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(r)}
}
