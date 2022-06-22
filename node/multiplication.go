package node

type MultiplicationNode struct {
	BinaryNode
}

func NewMultiplicationNode(left, right AlgebraNode) *MultiplicationNode {
	return &MultiplicationNode{
		NewBinaryNode(left, right),
	}
}

func (n *MultiplicationNode) GenerateCode() string {
	return n.BinaryNode.GenerateCodeByOpt(BinaryOptMultiplcation)
}
