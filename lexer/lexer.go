//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package lexer

import (
	"github.com/cloustone/macaca/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	lineno       int  // curent line no
}

// New create a lexter with repl string input
func New(input string) *Lexer {
	l := &Lexer{input: input, lineno: 1}
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
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.lineno)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch, l.lineno)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.lineno)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch, l.lineno)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch, l.lineno)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.lineno)
	case '<':
		tok = newToken(token.LT, l.ch, l.lineno)
	case '>':
		tok = newToken(token.GT, l.ch, l.lineno)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.lineno)
	case ':':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.DECLARE, Literal: literal, Lineno: l.lineno}
		} else {
			tok = newToken(token.COLON, l.ch, l.lineno)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch, l.lineno)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.lineno)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.lineno)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.lineno)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.lineno)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		tok.Lineno = l.lineno
	case '[':
		tok = newToken(token.LBRACKET, l.ch, l.lineno)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, l.lineno)
	case '\r':
	case '\n':
		l.lineno++
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Lineno = l.lineno
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Lineno = l.lineno
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			tok.Lineno = l.lineno
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.lineno)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
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

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte, lineno int) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
		Lineno:  lineno,
	}
}
