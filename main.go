package main

import "fmt"

type Token struct {
	token   string
	literal string
}

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func main() {
	lexer := Lexer{
		input: "let x = 6  + (5)",
	}

	lexer.readChar()

	for {
		tok := lexer.createToken()
		fmt.Println(tok)

		if tok.token == "EOF" {
			break
		}
	}

}

func (l *Lexer) createToken() Token {
	var tok Token
	l.skipWhiteSpace()

	switch l.ch {
	case '+':
		tok = Token{token: "PLUS", literal: string(l.ch)}
	case '-':
		tok = Token{token: "MiNUS", literal: string(l.ch)}
	case '=':
		tok = Token{token: "EQUAL", literal: string(l.ch)}
	case 0:
		tok = Token{token: "EOF", literal: ""}
	default:
		if isLetter(l.ch) {
			start := l.position
			for isLetter(l.ch) {
				l.readChar()
			}
			tok = Token{token: "IDENT", literal: l.input[start:l.position]}
			return tok
		} else if isDigit(l.ch) {
			tok = Token{token: "INT", literal: string(l.ch)}
		} else {
			tok = Token{token: "ILLEGAL", literal: string(l.ch)}

		}

	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
