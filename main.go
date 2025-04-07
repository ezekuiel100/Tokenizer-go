package main

import "fmt"

type Token struct {
	token   string
	literal string
}

type Lexer struct {
	input    string
	position int
	ch       string
}

func main() {
	lexer := Lexer{
		input:    "let x = 6",
		position: 0,
		ch:       "",
	}

	for len(lexer.input)-1 >= lexer.position {
		lexer.readChar()
	}

}

func (l *Lexer) createToken() Token {

	switch l.ch {
	case "+":
		return Token{token: "PLUS", literal: l.ch}
	case "=":
		return Token{token: "EQUAL", literal: l.ch}
	default:
		if l.isLetter(l.ch) {
			start := l.position
			for l.isLetter(l.peekChar()) {
				l.position++
			}

			return Token{token: "IDENT", literal: l.input[start : l.position+1]}
		}

	}
	return Token{}
}

func (l *Lexer) readChar() {
	l.ch = string(l.input[l.position])
	x := l.createToken()
	l.position++

	fmt.Println(x)

}

func (Lexer) isLetter(ch string) bool {
	return (ch >= "a" && ch <= "z") || (ch >= "A" && ch <= "Z") || ch == "_"
}

func (l *Lexer) peekChar() string {
	return string(l.input[l.position+1])
}
