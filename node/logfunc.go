package node

import "math"

type LogFuncNode struct {
	UnaryNode
}

func NewLogFuncNodeNode(n AlgebraNode) *LogFuncNode {
	return &LogFuncNode{
		UnaryNode: UnaryNode{
			n,
		},
	}
}

func (n *LogFuncNode) GenerateCode() string {

	return "math.Log(float64(" + n.input.GenerateCode() + "))"
}
func (n *LogFuncNode) Closures() Closures {
	return func(e Env) Number {
		return NewFloat(math.Log(n.input.Closures()(e).GetFloat()))
	}
}
