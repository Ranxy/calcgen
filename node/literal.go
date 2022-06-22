package node

import (
	"strconv"
)

type LiteralNode struct {
	ValueRaw string
	dataType DataType
}

func NewLiteralNode(val string) *LiteralNode {
	res := &LiteralNode{
		ValueRaw: val,
	}

	f, err := strconv.ParseFloat(val, 10)
	if err != nil {
		panic(err)
	}
	res.dataType = DataTypeFloat
	if float64(int(f)) == f {
		res.dataType = DataTypeInt
	}

	return res
}

func (n *LiteralNode) GetType() DataType {
	return n.dataType
}
func (n *LiteralNode) GenerateCode() string {
	return n.ValueRaw
}
func (n *LiteralNode) GetInput() []AlgebraNode {
	return []AlgebraNode{n}
}
