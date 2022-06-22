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
