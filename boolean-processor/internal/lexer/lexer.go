package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenVariable
	TokenOperator
	TokenLParen
	TokenRParen
)

type Token struct {
	Type  TokenType
	Value string
	Pos   int
}

type Lexer struct {
	input []rune
	pos   int
}

func New(input string) *Lexer {
	return &Lexer{
		input: []rune(input),
		pos:   0,
	}
}

func (l *Lexer) peek() rune {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) advance() rune {
	character := l.peek()
	l.pos++
	return character
}

func (l *Lexer) skipSpaces() {
	for unicode.IsSpace(l.peek()) {
		l.advance()
	}
}

func (l *Lexer) readVariable() string {
	start := l.pos

	if !unicode.IsLetter(l.peek()) {
		return ""
	}
	l.advance()

	if l.peek() == '_' {
		l.advance()
		for unicode.IsLetter(l.peek()) {
			l.advance()
		}
	}

	return string(l.input[start:l.pos])
}

func (l *Lexer) NextToken() (Token, error) {
	l.skipSpaces()
	character := l.peek()

	//variables
	if character == 0 {
		return Token{Type: TokenEOF}, nil
	}

	//operators
	operators := "+&@~>|!-"
	if strings.ContainsRune(operators, character) {
		l.advance()
		return Token{Type: TokenOperator, Value: string(character), Pos: l.pos}, nil
	}

	//parentheses groups
	if character == '(' || character == '[' || character == '{' {
		l.advance()
		return Token{Type: TokenLParen, Value: string(character), Pos: l.pos}, nil
	}
	if character == ')' || character == ']' || character == '}' {
		l.advance()
		return Token{Type: TokenRParen, Value: string(character), Pos: l.pos}, nil
	}

	return Token{}, fmt.Errorf("unknown character '%c' at position %d", character, l.pos)
}
