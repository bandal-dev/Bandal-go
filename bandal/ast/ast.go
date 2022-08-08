package ast

import "github.com/Bandal-dev/Bandal-go/bandal/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type BindingStatement struct {
	Name  *Identifier
	Value Expression
}

func (bs *BindingStatement) statementNode() {}

//func (bs *BindingStatement) TokenLiteral() string { return bs.Token.Literal}

type Identifier struct {
	Token token.Token //token.Identifier
	Value string
}

func (i *Identifier) experssionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
