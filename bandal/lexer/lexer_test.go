package lexer

import (
	"fmt"
	"testing"

	"github.com/Bandal-dev/Bandal-go/bandal/token"
)

func TestNextToken(t *testing.T) {
	input := `
	false
	=
	`

	tests := []struct {
		epectedType     token.TokenType
		expectedLiteral string
	}{
		{token.FALSE, "FALSE"},
		{token.Equal, "="},
	}

	l := Lex(input)

	for _, i := range tests {
		tok := l.NextToken()
		fmt.Println(i)
		fmt.Println(tok.Type, tok.Literal)
	}
}
