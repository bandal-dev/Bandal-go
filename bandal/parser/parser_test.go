package parser

import (
	"testing"

	"github.com/Bandal-dev/Bandal-go/bandal/ast"
	"github.com/Bandal-dev/Bandal-go/bandal/lexer"
)

func TestLetStatements(t *testing.T) {
	input = `
	x = 5
	y = 10
	foobar = 9234809
	`

	l := lexer.Lex(input)
	p := Parse(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("nill")
	}
	
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != 
}
