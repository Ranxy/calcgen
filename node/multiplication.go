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
func (n *MultiplicationNode) Closures() Closures {
	return func(e Env) Number {
		return n.Left.Closures()(e).multiplication(n.Right.Closures()(e))
	}
}
