package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LeftParen    = "("
	RightParen   = ")"
	LeftBrace    = "{"
	RightBrace   = "}"
	LeftBracket  = "["
	RightBracket = "]"
	SemiColon    = ";"
	Colon        = ":"

	Comma      = ","
	Dot        = "."
	Plus       = "+"
	Minus      = "-"
	Star       = "*"
	Slash      = "/"
	SlashSlash = "//"
	Caret      = "^"
	Percent    = "%"
	Underbar   = "_"
	Tilde      = "~"

	Equal        = "="
	EqualEqual   = "=="
	Greater      = ">"
	GreaterEqual = ">="
	Less         = "<"
	LessEqual    = "<="

	Undefined  = "UNDEFINED"
	Identifier = "IDENTIFIER"
	EOF        = "EOF"
	EOL        = "EOL"

	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	IN       = "IN"
	END      = "END"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	NULL     = "NULL"
	PASS     = "PASS"
	NOT      = "NOT"
	AND      = "AND"
	OR       = "OR"

	INT8  = "INT8"
	INT16 = "INT16"
	INT32 = "INT32" //TODO-------------------------------------add more tokens by doing nogada
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"return": RETURN,
	"in":     IN,
	"end":    END,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"null":   NULL,
	"pass":   PASS,
	"not":    NOT,
	"and":    AND,
	"or":     OR,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Identifier
}
