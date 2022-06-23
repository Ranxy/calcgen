package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"plugin"
	"strings"

	"github.com/Ranxy/calcgen/node"
)

func main() {
	query := Parse(strings.NewReader(os.Args[1]))

	code := generateCode(query)
	fmt.Println("S:", code)

	buildPlugin(code)
	callPlugin()

	env := node.Env{
		"a": 2.1,
	}
	value := query.Expr.Closures()(env)
	fmt.Println(value)

}

func buildPlugin(code string) {
	ioutil.WriteFile("/tmp/calc.go", []byte(code), 0644)
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", "/tmp/a.so", "/tmp/calc.go")
	cmd.Run()
}

func callPlugin() {
	p, err := plugin.Open("/tmp/a.so")
	if err != nil {
		panic(err)
	}
	sym, err := p.Lookup("Calc")
	if err != nil {
		panic(err)
	}
	sym.(func())()
}

func generateCode(q Query) string {
	temp := `package main
import "fmt"
func Calc() {
	res:=%s
	fmt.Println(res)
}`
	return fmt.Sprintf(temp, q.Expr.GenerateCode())
}
