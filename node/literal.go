package node

import (
	"strconv"
	"strings"
)

type LiteralNode struct {
	ValueRaw any
	valueStr string
	dataType DataType
}

func NewLiteralNode(val any) *LiteralNode {
	res := &LiteralNode{
		ValueRaw: val,
	}

	switch v := val.(type) {
	case int:
		res.dataType = DataTypeInt
		res.valueStr = strconv.Itoa(int(v))
	case float64:
		res.dataType = DataTypeFloat
		res.valueStr = strconv.FormatFloat(v, 'f', 10, 64)
	case string:
		res.valueStr = v
		f, err := strconv.ParseFloat(v, 10)
		if err != nil {
			panic(err)
		}
		res.dataType = DataTypeFloat
		res.ValueRaw = f
		if float64(int(f)) == f && !strings.Contains(v, ".") {
			res.dataType = DataTypeInt
			res.ValueRaw = int(f)
		}
	}

	return res
}

func (n *LiteralNode) GetType() DataType {
	return n.dataType
}
func (n *LiteralNode) GenerateCode() string {
	return n.valueStr
}
func (n *LiteralNode) GetInput() []AlgebraNode {
	return []AlgebraNode{n}
}

func (n *LiteralNode) Closures() Closures {
	return func(e Env) Number {
		if n.dataType == DataTypeInt {
			return NewInt(n.ValueRaw.(int))
		} else {
			return NewFloat(n.ValueRaw.(float64))
		}
	}
}
