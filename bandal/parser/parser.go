package parser

import (
	"github.com/Bandal-dev/Bandal-go/bandal/ast"
	"github.com/Bandal-dev/Bandal-go/bandal/lexer"
	"github.com/Bandal-dev/Bandal-go/bandal/token"
)

type Parser struct {
	l       *lexer.Lexer
	current token.Token
	peek    token.Token
}

func Parse(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.current = p.peek
	p.peek = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
