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
func (n *MinusNode) Closures() Closures {
	return func(e Env) Number {
		return n.Left.Closures()(e).minus(n.Right.Closures()(e))
	}
}
