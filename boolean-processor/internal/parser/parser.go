package parser

import (
	"fmt"

	"boolean-processor/internal/ast"
	"boolean-processor/internal/lexer"
)

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func New(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

func (p *Parser) peek() lexer.Token {
	if p.pos >= len(p.tokens) {
		return lexer.Token{Type: lexer.TokenEOF}
	}
	return p.tokens[p.pos]
}

func (p *Parser) advance() lexer.Token {
	t := p.peek()
	p.pos++
	return t
}

func (p *Parser) expect(t lexer.TokenType) (lexer.Token, error) {
	token := p.peek()
	if token.Type != t {
		return token, fmt.Errorf("expected token %v, got %v at pos %d", t, token, token.Pos)
	}
	p.advance()
	return token, nil
}

func (p *Parser) ParseExpression() (ast.Node, error) {
	tok := p.peek()

	switch tok.Type {

	case lexer.TokenVariable:
		p.advance()
		return &ast.VarNode{Name: tok.Value}, nil

	case lexer.TokenOperator:
		if tok.Value == "-" || tok.Value == "!" {
			p.advance()
			expr, err := p.ParseExpression()
			if err != nil {
				return nil, err
			}
			return &ast.UnaryNode{Op: tok.Value, Expr: expr}, nil
		}
		return nil, fmt.Errorf("unexpected operator '%s' at pos %d", tok.Value, tok.Pos)

	case lexer.TokenLParen:
		return p.parseBinary()

	default:
		return nil, fmt.Errorf("unexpected token '%s' at pos %d", tok.Value, tok.Pos)
	}
}

func (p *Parser) parseBinary() (ast.Node, error) {
	open := p.peek()
	p.advance()

	left, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	opTok := p.peek()
	if opTok.Type != lexer.TokenOperator {
		return nil, fmt.Errorf("expected operator after left expression, got '%s'", opTok.Value)
	}
	p.advance()
	op := opTok.Value

	right, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	closeTok := p.peek()
	if closeTok.Type != lexer.TokenRParen {
		return nil, fmt.Errorf("expected closing bracket, found '%s'", closeTok.Value)
	}

	if !matchingBrackets(open.Value, closeTok.Value) {
		return nil, fmt.Errorf("bracket mismatch: '%s' does not match '%s'", open.Value, closeTok.Value)
	}

	p.advance()
	return &ast.BinaryNode{
		Op:    op,
		Left:  left,
		Right: right,
	}, nil
}

func matchingBrackets(open, close string) bool {
	switch open {
	case "(":
		return close == ")"
	case "[":
		return close == "]"
	case "{":
		return close == "}"
	}
	return false
}
