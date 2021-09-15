package equation

import (
	"errors"
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
		if markedExpressionPart.contentType == mathSeparator {
			continue
		}

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
			var isPreviousNodeStickOrSurroundOperator bool

			previousOperatorNodeId, ok := operatorNodeIdStack.peek()
			if ok {
				operatorNode, err := tree.getOperatorNode(previousOperatorNodeId)
				if err != nil {
					return err
				}

				previousOperatorNode = operatorNode
				previousNode, err := tree.getNode(previousOperatorNodeId)
				if err != nil {
					return err
				}

				isPreviousNodeStick := previousNode.equationNodeType() == graphOperatorStickNode
				isPreviousNodeSurroundOperator := previousOperatorNode.operator.placeType == surroundOperator

				isPreviousNodeStickOrSurroundOperator = isPreviousNodeStick || isPreviousNodeSurroundOperator
			}

			isSymbolSurroundOperator := !ok || (ok && isPreviousNodeStickOrSurroundOperator)

			for _, defaultOperator := range defaultOperationList {
				if defaultOperator.surroundSign.start == markedExpressionPart.content {
					if isSymbolSurroundOperator && defaultOperator.placeType == surroundOperator {
						operator = defaultOperator
						break
					} else if defaultOperator.placeType == prefixOperator {
						operator = defaultOperator
						break
					}
				}
			}

			if ok && previousOperatorNode.operator.placeType == infixOperator {
				return errors.New("cannot be done")
			}

			if isSymbolSurroundOperator {
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
			lastNode, err := tree.getNode(lastOperatorId)
			if err != nil {
				return err
			}

			if lastNode.equationNodeType() == graphOperatorStickNode {
				lastOperatorId, _ = operatorNodeIdStack.pop()
			}

			prevNodeId = lastOperatorId
		} else {
			prevNodeId = uid
		}
	}

	for x, v := range tree.tree {
		if (*v).equationNodeType() == graphOperatorNode {
			fmt.Println(x, (*v).(mathOperatorNode).operatorArgs, (*v).equationGraphLevel())
		}
	}

	return nil
}
