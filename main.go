package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	query := Parse(strings.NewReader(os.Args[1]))
	s := query.Expr.GenerateCode()

	fmt.Println("S:", s)
}
