package equation

import (
	"errors"
	"fmt"
	"sort"
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

type equationGraph struct {
	tree map[uint]*mathNode
}

func (eg *equationGraph) getNode(id uint) (mathNode, error) {
	node := eg.tree[id]
	if node == nil {
		return mathOperatorNode{}, errors.New(fmt.Sprintf("%d is not a node", id))
	}

	return *node, nil
}

func (eg *equationGraph) upsert(id uint, node mathNode) {
	eg.tree[id] = &node
}

func (eg *equationGraph) getOperatorNode(id uint) (mathOperatorNode, error) {
	node, err := eg.getNode(id)
	if err != nil {
		return mathOperatorNode{}, err
	}

	operatorNode, ok := node.(mathOperatorNode)
	if !ok {
		return mathOperatorNode{}, errors.New(fmt.Sprintf("%d is not a mathOperatorNode", id))
	}

	return operatorNode, nil
}

func (eg *equationGraph) chanOverTree() chan mathNode {
	keys := make([]uint, 0)
	for index, _ := range eg.tree {
		keys = append(keys, index)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	ch := make(chan mathNode)

	go func() {
		for _, index := range keys {
			ch <- *eg.tree[index]
		}

		close(ch)
	}()

	return ch
}

func newEquationTree() equationGraph {
	return equationGraph{tree: map[uint]*mathNode{}}
}
