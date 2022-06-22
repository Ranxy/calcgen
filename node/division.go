package node

type DivisionNode struct {
	BinaryNode
}

func NewDivisionNode(left, right AlgebraNode) *DivisionNode {
	return &DivisionNode{
		NewBinaryNode(left, right),
	}
}

func (n *DivisionNode) GenerateCode() string {
	return n.BinaryNode.GenerateCodeByOpt(BinaryOptDivision)
}
