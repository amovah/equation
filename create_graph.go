package equation

import (
	"fmt"
	"strconv"
)

func applyArgsToPreviousOperator(
	operatorNodeIdStack *stack,
	tree *equationTree,
	uid uint,
	markedExpressionPart mathPartType,
) error {
	var isMathSurroundStart = markedExpressionPart == mathSurroundStart
	var isMathSurroundEnd = markedExpressionPart == mathSurroundEnd

	previousNodeId, _ := operatorNodeIdStack.peek()
	previousNode, err := tree.getOperatorNode(previousNodeId)
	if err != nil {
		return err
	}

	previousNodeOperatorPlaceType := previousNode.operator.placeType
	if previousNodeOperatorPlaceType == infixOperator {
		previousNode.operatorArgs = append(previousNode.operatorArgs, uid)
		tree.upsert(previousNode.id, previousNode)
		operatorNodeIdStack.pop()

		_, ok := operatorNodeIdStack.peek()
		if ok {
			err := applyArgsToPreviousOperator(operatorNodeIdStack, tree, uid, markedExpressionPart)
			if err != nil {
				return err
			}
		}
	} else if previousNodeOperatorPlaceType == surroundOperator && !isMathSurroundEnd {
		previousNode.operatorArgs = append(previousNode.operatorArgs, uid)
		tree.upsert(previousNode.id, previousNode)
	} else if previousNodeOperatorPlaceType == prefixOperator && !isMathSurroundEnd && !isMathSurroundStart {
		previousNode.operatorArgs = append(previousNode.operatorArgs, uid)
		tree.upsert(previousNode.id, previousNode)
	}

	return nil
}

func createGraph(markedExpressionParts []markedExpression) error {
	tree := newEquationTree()

	var prevNodeId uint
	operatorNodeIdStack := newStack()
	var currentGraphLevel uint = 0

	for id, markedExpressionPart := range markedExpressionParts {
		uid := uint(id)

		contentType := markedExpressionPart.contentType

		if operatorNodeIdStack.size() > 0 {
			err := applyArgsToPreviousOperator(&operatorNodeIdStack, &tree, uid, contentType)
			if err != nil {
				return err
			}
		}

		if contentType == mathNumber {
			converted, err := strconv.ParseFloat(markedExpressionPart.content, 64)
			if err != nil {
				converted = 0
			}

			tree.upsert(uid, mathNumberNode{
				id:         uid,
				value:      converted,
				graphLevel: currentGraphLevel,
			})
		}

		if contentType == mathSymbol {
			var operator equationOperator
			for _, defaultOperator := range defaultOperationList {
				if defaultOperator.symbol == markedExpressionPart.content {
					operator = defaultOperator
					break
				}
			}

			var operatorArgs []uint
			if operator.placeType == infixOperator {
				operatorArgs = append(operatorArgs, prevNodeId)
			}

			tree.upsert(uid, mathOperatorNode{
				id:           uid,
				operator:     operator,
				operatorArgs: operatorArgs,
				graphLevel:   currentGraphLevel,
			})

			operatorNodeIdStack.push(uid)
		}

		if contentType == mathSurroundStart {
			var operator equationOperator
			var previousOperatorNode mathOperatorNode

			previousOperatorNodeId, ok := operatorNodeIdStack.peek()
			var isPreviousNodeStickOrSurround bool
			if ok {
				previousNode, err := tree.getOperatorNode(previousOperatorNodeId)
				if err != nil {
					return err
				}

				previousOperatorNode = previousNode
				isPreviousNodeStick := tree.getNode(previousOperatorNodeId).equationNodeType() == graphOperatorStickNode
				isPreviousNodeSurround := previousOperatorNode.operator.placeType == surroundOperator

				isPreviousNodeStickOrSurround = isPreviousNodeStick || isPreviousNodeSurround
			}

			isSymbolSurround := !ok || (ok && isPreviousNodeStickOrSurround)

			for _, defaultOperator := range defaultOperationList {
				if defaultOperator.surroundSign.start == markedExpressionPart.content {
					if isSymbolSurround && defaultOperator.placeType == surroundOperator {
						operator = defaultOperator
						break
					} else if defaultOperator.placeType == prefixOperator {
						operator = defaultOperator
						break
					}
				}
			}

			if ok && previousOperatorNode.operator.placeType == infixOperator {
				panic("cannot be done")
			}

			if isSymbolSurround {
				tree.upsert(uid, mathOperatorNode{
					id:           uid,
					operator:     operator,
					operatorArgs: []uint{},
					graphLevel:   currentGraphLevel,
				})
				currentGraphLevel = currentGraphLevel + 1

				operatorNodeIdStack.push(uid)
			} else {
				tree.upsert(uid, mathOperatorStickerNode{
					id: uid,
				})
			}

		}

		if contentType == mathSurroundEnd {
			currentGraphLevel = currentGraphLevel - 1
			lastOperatorId, _ := operatorNodeIdStack.pop()
			if tree.getNode(lastOperatorId).equationNodeType() == graphOperatorStickNode {
				lastOperatorId, _ = operatorNodeIdStack.pop()
			}

			prevNodeId = lastOperatorId
		} else {
			prevNodeId = uid
		}
	}

	for x, v := range tree.tree {
		if v.equationNodeType() == graphOperatorNode {
			fmt.Println(x, v.(mathOperatorNode).operatorArgs, v.equationGraphLevel())
		}
	}

	return nil
}
