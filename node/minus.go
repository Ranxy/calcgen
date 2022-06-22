package node

type MinusNode struct {
	BinaryNode
}

func NewMinusNode(left, right AlgebraNode) *MinusNode {
	return &MinusNode{
		NewBinaryNode(left, right),
	}
}

func (n *MinusNode) GenerateCode() string {
	return n.BinaryNode.GenerateCodeByOpt(BinaryOptMinus)
}
