package equation

import (
	"errors"
	"fmt"
)

type graphNodeType uint8

const (
	graphOperatorNode graphNodeType = iota
	graphOperatorStickNode
	graphNumberNode
)

type mathNode interface {
	equationId() uint
	equationNodeType() graphNodeType
	equationNextNode() (uint, bool)
	equationPreviousNode() (uint, bool)
	equationGraphLevel() uint
}

type mathOperatorStickerNode struct {
	id uint
}

func (m mathOperatorStickerNode) equationId() uint {
	return m.id
}

func (m mathOperatorStickerNode) equationNodeType() graphNodeType {
	return graphOperatorStickNode
}

func (m mathOperatorStickerNode) equationNextNode() (uint, bool) {
	return 0, false
}

func (m mathOperatorStickerNode) equationPreviousNode() (uint, bool) {
	return 0, false
}

func (m mathOperatorStickerNode) equationGraphLevel() uint {
	return 0
}

type mathOperatorNode struct {
	id              uint
	operator        equationOperator
	operatorArgs    []uint
	nextNode        uint
	previousNode    uint
	hasNextNode     bool
	hasPreviousNode bool
	graphLevel      uint
}

func (m mathOperatorNode) equationId() uint {
	return m.id
}

func (m mathOperatorNode) equationNodeType() graphNodeType {
	return graphOperatorNode
}

func (m mathOperatorNode) equationNextNode() (uint, bool) {
	return m.nextNode, m.hasNextNode
}

func (m mathOperatorNode) equationPreviousNode() (uint, bool) {
	return m.previousNode, m.hasPreviousNode
}

func (m mathOperatorNode) equationGraphLevel() uint {
	return m.graphLevel
}

type mathNumberNode struct {
	id              uint
	value           float64
	nextNode        uint
	hasNextNode     bool
	previousNode    uint
	hasPreviousNode bool
	graphLevel      uint
}

func (m mathNumberNode) equationId() uint {
	return m.id
}

func (m mathNumberNode) equationNodeType() graphNodeType {
	return graphNumberNode
}

func (m mathNumberNode) equationNextNode() (uint, bool) {
	return m.nextNode, m.hasNextNode
}

func (m mathNumberNode) equationPreviousNode() (uint, bool) {
	return m.previousNode, m.hasPreviousNode
}

func (m mathNumberNode) equationGraphLevel() uint {
	return m.graphLevel
}

type equationTree struct {
	tree map[uint]*mathNode
}

func (et *equationTree) getNode(id uint) (mathNode, error) {
	node := et.tree[id]
	if node == nil {
		return mathOperatorNode{}, errors.New(fmt.Sprintf("%d is not a node", id))
	}

	return *node, nil
}

func (et *equationTree) upsert(id uint, node mathNode) {
	et.tree[id] = &node
}

func (et *equationTree) getOperatorNode(id uint) (mathOperatorNode, error) {
	node, err := et.getNode(id)
	if err != nil {
		return mathOperatorNode{}, err
	}

	operatorNode, ok := node.(mathOperatorNode)
	if !ok {
		return mathOperatorNode{}, errors.New(fmt.Sprintf("%d is not a mathOperatorNode", id))
	}

	return operatorNode, nil
}

func newEquationTree() equationTree {
	return equationTree{tree: map[uint]*mathNode{}}
}
