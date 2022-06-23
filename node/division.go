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
func (n *DivisionNode) Closures() Closures {
	return func(e Env) Number {
		return n.Left.Closures()(e).division(n.Right.Closures()(e))
	}
}
