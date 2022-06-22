package node

type BinaryOpt string

const (
	BinaryOptPlus          = "+"
	BinaryOptMinus         = "-"
	BinaryOptMultiplcation = "*"
	BinaryOptDivision      = "/"
)

func (b BinaryOpt) OptString() string {
	return string(b)
}

type BinaryNode struct {
	dataType DataType
	Left     AlgebraNode
	Right    AlgebraNode
}

func NewBinaryNode(left, right AlgebraNode) BinaryNode {
	res := BinaryNode{
		Left:  left,
		Right: right,
	}

	//res.AbstractNode.AlgebraNode = &res

	return res
}

func (n *BinaryNode) GetType() DataType {
	if n.dataType == 0 {
		n.dataType = n.TypeFromInput()
	}
	return n.dataType
}

func (n *BinaryNode) TypeFromInput() DataType {
	inputs := n.GetInput()
	if len(inputs) == 0 {
		panic("Inputs shoule not be nil")
	}
	//if one of input is float , this node must be float
	for _, input := range inputs {
		if input.GetType() == DataTypeFloat {
			return DataTypeFloat
		}
	}
	return DataTypeInt
}

func (b *BinaryNode) GetInput() []AlgebraNode {
	return []AlgebraNode{b.Left, b.Right}
}

func (n *BinaryNode) GenerateCodeByOpt(opt BinaryOpt) string {
	if n.Left.GetType() == DataTypeFloat && n.Right.GetType() == DataTypeFloat {
		return "(" + n.Left.GenerateCode() + opt.OptString() + n.Right.GenerateCode() + ")"
	} else if n.Left.GetType() == DataTypeFloat && n.Right.GetType() == DataTypeInt {
		return "(" + n.Left.GenerateCode() + opt.OptString() + " float64(" + n.Right.GenerateCode() + "))"
	} else if n.Left.GetType() == DataTypeInt && n.Right.GetType() == DataTypeFloat {
		return "(float64(" + n.Left.GenerateCode() + ")" + opt.OptString() + n.Right.GenerateCode() + ")"
	} else if n.Left.GetType() == DataTypeInt && n.Right.GetType() == DataTypeInt {
		return "(" + n.Left.GenerateCode() + opt.OptString() + n.Right.GenerateCode() + ")"
	} else if n.Left.GetType() == DataTypeVariable && n.Right.GetType() == DataTypeInt {
		return "(int(" + n.Left.GenerateCode() + ")" + opt.OptString() + n.Right.GenerateCode() + ")"
	} else if n.Left.GetType() == DataTypeVariable && n.Right.GetType() == DataTypeFloat {
		return "(float64(" + n.Left.GenerateCode() + ")" + opt.OptString() + n.Right.GenerateCode() + ")"
	} else if n.Left.GetType() == DataTypeInt && n.Right.GetType() == DataTypeVariable {
		return "(" + n.Left.GenerateCode() + opt.OptString() + " int(" + n.Right.GenerateCode() + "))"
	} else if n.Left.GetType() == DataTypeFloat && n.Right.GetType() == DataTypeVariable {
		return "(" + n.Left.GenerateCode() + opt.OptString() + " float64(" + n.Right.GenerateCode() + "))"
	}

	panic("Type Not Support")
}
