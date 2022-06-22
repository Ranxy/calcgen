package node

type VariableNode struct {
	name     string
	dataType DataType
}

func NewVariableNode(name string, tp DataType) *VariableNode {
	return &VariableNode{
		name:     name,
		dataType: tp,
	}
}

func NewVariableNodeX(name string) *VariableNode {
	return &VariableNode{
		name:     name,
		dataType: DataTypeVariable,
	}
}

func (n *VariableNode) GetType() DataType {
	return n.dataType
}
func (n *VariableNode) GenerateCode() string {
	return n.name
}
func (n *VariableNode) GetInput() []AlgebraNode {
	return []AlgebraNode{n}
}
