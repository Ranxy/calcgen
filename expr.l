%{
package main

import (
    "github.com/Ranxy/calcgen/node"
)

type Number node.AlgebraNode
type Variable node.AlgebraNode
%}

%union{
	token string
	expr node.AlgebraNode
    variable node.AlgebraNode
}

%type<expr> program
%type<expr> expr
%type<variable> variable

%token<token> NUMBER
%token<token> STR

%left '+' '-'
%left '*' '/' '%'
%left '$'

%%

program
	: expr
	{
		$$ = $1
		yylex.(*Lexer).result = $$
	}

expr
	: NUMBER
	{
		$$ = node.NewLiteralNode($1)
	}
    | '$' variable
	{
		$$ = $2
	}
    | '(' expr ')'
	{
		$$ = $2
	}
    | '-' expr
	{
		$$ = node.NewChangeSignNode($2)
	}
	| expr '+' expr
	{
		$$ = node.NewPlusNode($1,$3)
	}
	| expr '-' expr
	{
		$$ = node.NewMinusNode($1,$3)
	}
	| expr '*' expr
	{
		$$ = node.NewMultiplicationNode($1,$3)
	}
	| expr '/' expr
	{
		$$ = node.NewDivisionNode($1,$3)
	}
variable
    : STR
   	{
		$$ = node.NewVariableNodeX($1)
	} 
%%
