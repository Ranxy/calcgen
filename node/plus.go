package node

type PlusNode struct {
	BinaryNode
}

func NewPlusNode(left, right AlgebraNode) *PlusNode {
	return &PlusNode{
		NewBinaryNode(left, right),
	}
}

func (n *PlusNode) GenerateCode() string {
	return n.BinaryNode.GenerateCodeByOpt(BinaryOptPlus)
}
func (n *PlusNode) Closures() Closures {
	return func(e Env) Number {
		return n.Left.Closures()(e).Add(n.Right.Closures()(e))
	}
}
