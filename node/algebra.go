package node

type AlgebraNode interface {
	GetType() DataType
	GenerateCode() string
	GetInput() []AlgebraNode
}

var _ AlgebraNode = &PlusNode{}
var _ AlgebraNode = &MinusNode{}
var _ AlgebraNode = &MultiplicationNode{}
var _ AlgebraNode = &DivisionNode{}
var _ AlgebraNode = &LiteralNode{}
