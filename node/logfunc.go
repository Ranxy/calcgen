package node

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
