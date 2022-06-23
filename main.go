package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Ranxy/calcgen/node"
)

func main() {
	query := Parse(strings.NewReader(os.Args[1]))
	s := query.Expr.GenerateCode()

	fmt.Println("S:", s)

	env := node.Env{
		"a": 2.1,
	}
	value := query.Expr.Closures()(env)
	fmt.Println(value)

}
