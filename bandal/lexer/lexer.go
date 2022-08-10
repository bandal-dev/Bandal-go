package lexer

import "github.com/Bandal-dev/Bandal-go/bandal/token"

type Lexer struct {
	source string //input source code
	//start    int
	position int
	current  int

	//line     int
	ch byte
}

func Lex(input string) *Lexer {
	l := &Lexer{source: input}
	l.Advance()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var ret token.Token
	l.Space()

	switch l.ch {
	case '(':
		ret = AddToken(token.LeftParen, l.ch)
	case ')':
		ret = AddToken(token.RightParen, l.ch)
	case '{':
		ret = AddToken(token.LeftBrace, l.ch)
	case '}':
		ret = AddToken(token.RightBrace, l.ch)
	case '[':
		ret = AddToken(token.LeftBracket, l.ch)
	case ']':
		ret = AddToken(token.RightBracket, l.ch)
	case ',':
		ret = AddToken(token.Comma, l.ch)
	case '.':
		ret = AddToken(token.Dot, l.ch)
	case '+':
		ret = AddToken(token.Plus, l.ch)
	case ':':
		ret = AddToken(token.Colon, l.ch)
	case ';':
		ret = AddToken(token.SemiColon, l.ch)
	case '*':
		ret = AddToken(token.Star, l.ch)
	case '^':
		ret = AddToken(token.Caret, l.ch)
	case '~':
		ret = AddToken(token.Tilde, l.ch)
	case '/':
		ret = AddToken(token.Slash, l.ch)
	//case '//':
	//	ret = AddToken(token.SlashSlash, l.ch)
	case '%':
		ret = AddToken(token.Percent, l.ch)
	case '_':
		ret = AddToken(token.Underbar, l.ch)
	case '=':
		if l.Peek() == '=' {
			ch := l.ch
			l.Advance()
			ret = token.Token{Type: token.EqualEqual, Literal: string(ch) + string(l.ch)}
		} else {
			ret = AddToken(token.Equal, l.ch)
		}
	case 0:
		ret.Literal = ""
		ret.Type = token.EOF

	default:
		if IsLetter(l.ch) {
			ret.Literal = l.Identifier()
			ret.Type = token.LookUpIdent(ret.Literal)
			return ret
		} else if IsDigit(l.ch) {
			ret.Literal = l.Number()
			ret.Type = token.INT
			return ret
		} else {
			ret = AddToken(token.Undefined, l.ch)

		}
	}
	l.Advance()
	return ret
}

func AddToken(tokentype token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokentype, Literal: string(ch)}

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

func (l *Lexer) Space() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.Advance()
	}
}
