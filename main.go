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
		input: "let x = 6",
	}

	lexer.readChar()

	for lexer.readPosition < len(lexer.input) {
		lexer.createToken()
	}

}

func (l *Lexer) createToken() {
	var tok Token
	l.skipWhiteSpace()

	switch l.ch {
	case '+':
		tok = Token{token: "PLUS", literal: string(l.ch)}
	case '=':
		tok = Token{token: "EQUAL", literal: string(l.ch)}
	case 0:
		tok = Token{token: "EOF", literal: ""}
	default:
		if l.isLetter(l.ch) {
			start := l.position
			for l.isLetter(l.peekChar()) {
				l.readChar()
			}
			tok = Token{token: "IDENT", literal: l.input[start:l.readPosition]}
		}

	}

	l.readChar()
	fmt.Println(tok)
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
		l.position = l.readPosition
		l.readPosition++
	}

}

func (Lexer) isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition > len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhiteSpace() {
	if l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
