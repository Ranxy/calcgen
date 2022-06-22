package node

type ChangeSignNode struct {
	UnaryNode
}

func NewChangeSignNode(n AlgebraNode) *ChangeSignNode {
	return &ChangeSignNode{
		UnaryNode: UnaryNode{
			n,
		},
	}
}

func (n *ChangeSignNode) GenerateCode() string {
	return "(-(" + n.input.GenerateCode() + "))"
}
