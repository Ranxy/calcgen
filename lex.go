package main

import (
	"io"
	"text/scanner"

	"github.com/Ranxy/calcgen/node"
)

type Expression node.AlgebraNode
type Lexer struct {
	scanner.Scanner

	result   Expression
	variable []string
}

func (l *Lexer) Lex(lval *yySymType) int {

	token := int(l.Scan())
	if token == scanner.Int || token == scanner.Float {
		token = NUMBER
	} else if token == scanner.Ident {
		l.variable = append(l.variable, l.TokenText())
		token = STR
	}

	lval.token = l.TokenText()
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}

type Query struct {
	Expr  Expression
	Param []string
}

func Parse(file io.Reader) Query {
	l := new(Lexer)
	l.Init(file)

	yyParse(l)
	return Query{
		Expr:  l.result,
		Param: l.variable,
	}
}
