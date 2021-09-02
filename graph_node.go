package equation

type mathNode interface {
	equationId() int
	equationValue() float64
	equationPrecedence() uint16
	equationSubNodes() []mathNode
	equationWithSubNodes([]mathNode) mathNode
}

type mathOperatorNode struct {
	id       int
	operator equationOperator
	subNodes []mathNode
}

func (m mathOperatorNode) equationId() int {
	return m.id
}

func (m mathOperatorNode) equationValue() float64 {
	var nums []float64
	for _, subNode := range m.subNodes {
		nums = append(nums, subNode.equationValue())
	}

	return m.operator.operation(nums...)
}

func (m mathOperatorNode) equationPrecedence() uint16 {
	return m.operator.precedence
}

func (m mathOperatorNode) equationSubNodes() []mathNode {
	return m.subNodes
}

func (m mathOperatorNode) equationWithSubNodes(subNodes []mathNode) mathNode {
	m.subNodes = subNodes
	return m
}

type mathNumberNode struct {
	id    int
	value float64
}

func (m mathNumberNode) equationId() int {
	return m.id
}

func (m mathNumberNode) equationValue() float64 {
	return m.value
}

func (m mathNumberNode) equationPrecedence() uint16 {
	return numberPrecedence
}

func (m mathNumberNode) equationSubNodes() []mathNode {
	return nil
}

func (m mathNumberNode) equationWithSubNodes(subNodes []mathNode) mathNode {
	return m
}
