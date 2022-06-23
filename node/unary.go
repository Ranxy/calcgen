package node

type UnaryNode struct {
	input AlgebraNode
}

func (n *UnaryNode) GetType() DataType {
	return n.input.GetType()
}
func (n *UnaryNode) GenerateCode() string {
	return n.input.GenerateCode()
}
func (n *UnaryNode) GetInput() []AlgebraNode {
	return []AlgebraNode{n}
}
func (n *UnaryNode) Closures() Closures {
	return func(e Env) Number {
		return n.input.Closures()(e)
	}
}
