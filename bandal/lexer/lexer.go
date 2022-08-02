package lexer

import "bandal/token"

type Lexer struct {
	source string //input source code
	//start    int
	position int
	current  int
	//line     int
	tokens []token.Token
	ch     byte
}

func Lex(input string) []token.Token {
	l := &Lexer{source: input}
	l.Advance()
	for !l.IsEnd() {
		l.ScanToken()
	}
	return l.tokens
}

func (l *Lexer) ScanToken() {

	switch l.ch {
	case '(':
		l.AddToken(token.LeftParen, l.ch)
	case ')':
		l.AddToken(token.RightParen, l.ch)
	case '{':
		l.AddToken(token.LeftBrace, l.ch)
	case '}':
		l.AddToken(token.RightBrace, l.ch)
	case '[':
		l.AddToken(token.LeftBracket, l.ch)
	case ']':
		l.AddToken(token.RightBracket, l.ch)
	case ',':
		l.AddToken(token.Comma, l.ch)
	case '.':
		l.AddToken(token.Dot, l.ch)
	case '+':
		l.AddToken(token.Plus, l.ch)
	case '-':
		l.AddToken(token.Minus, l.ch)

	case '=':
		if l.Peek() == '=' {
			ch := l.ch
			l.Advance()
			l.tokens = append(l.tokens, token.Token{Type: token.EqualEqual, Literal: string(ch) + string(l.ch)})
		} else {
			l.AddToken(token.Equal, l.ch)
		}

	default:
		if IsLetter(l.ch) {
			var tok token.Token
			tok.Literal = l.Identifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			l.tokens = append(l.tokens, tok)
		} else {
			l.AddToken(token.Undefined, l.ch)
		}
	}
	l.Advance()
}

func (l *Lexer) AddToken(tokentype token.TokenType, ch byte) {
	l.tokens = append(l.tokens, token.Token{Type: tokentype, Literal: string(ch)})

}

func (l *Lexer) Advance() {
	if l.current >= len(l.source) {
		l.ch = 0
	} else {
		l.ch = l.source[l.current]
	}
	l.position = l.current
	l.current += 1
}

func (l *Lexer) Identifier() string {
	fix := l.position
	for IsLetter(l.ch) {
		l.Advance()
	}
	return l.source[fix:l.position]
}

func (l *Lexer) Number() string {
	fix := l.position
	for IsDigit(l.ch) {
		l.Advance()
	}
	return l.source[fix:l.position]
}

func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch == 'Z' || ch == '_'
}

func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) IsEnd() bool {
	return l.current > len(l.source)
}

func (l *Lexer) Peek() byte {
	if l.current > len(l.source) {
		return 0
	} else {
		return l.source[l.current]
	}
}
