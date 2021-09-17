package equation

import (
	"errors"
	"strconv"
)

func applyArgsToPreviousOperator(
	operatorNodeIdStack *stack,
	graph *equationGraph,
	uid uint,
	markedExpressionPart mathPartType,
) error {
	var isMathSurroundStart = markedExpressionPart == mathSurroundStart
	var isMathSurroundEnd = markedExpressionPart == mathSurroundEnd

	previousNodeId, _ := operatorNodeIdStack.peek()
	previousNode, err := graph.getOperatorNode(previousNodeId)
	if err != nil {
		return err
	}

	previousNodeOperatorPlaceType := previousNode.operator.placeType
	if previousNodeOperatorPlaceType == infixOperator {
		previousNode.operatorArgs = append(previousNode.operatorArgs, uid)
		graph.upsert(previousNode.id, previousNode)
		operatorNodeIdStack.pop()

		_, ok := operatorNodeIdStack.peek()
		if ok {
			err := applyArgsToPreviousOperator(operatorNodeIdStack, graph, uid, markedExpressionPart)
			if err != nil {
				return err
			}
		}
	} else if previousNodeOperatorPlaceType == surroundOperator && !isMathSurroundEnd {
		previousNode.operatorArgs = append(previousNode.operatorArgs, uid)
		graph.upsert(previousNode.id, previousNode)
	} else if previousNodeOperatorPlaceType == prefixOperator && !isMathSurroundEnd && !isMathSurroundStart {
		previousNode.operatorArgs = append(previousNode.operatorArgs, uid)
		graph.upsert(previousNode.id, previousNode)
	}

	return nil
}

func findOperatorForMathSymbol(
	symbol string,
	infixOperatorMap map[string]equationOperator,
	prefixOperatorMap map[string]equationOperator,
) (equationOperator, error) {
	foundInfixOperator, found := findOperatorInMap(symbol, infixOperatorMap)
	if found {
		return foundInfixOperator, nil
	}

	foundPrefixOperator, found := findOperatorInMap(symbol, prefixOperatorMap)
	if found {
		return foundPrefixOperator, nil
	}

	return equationOperator{}, errors.New("cannot find operator with symbol " + symbol)
}

func findOperatorForMathSurround(
	symbol string,
	isSymbolSurroundOperator bool,
	surroundOperatorMap map[string]equationOperator,
	prefixOperatorMap map[string]equationOperator,
) (equationOperator, error) {
	if isSymbolSurroundOperator {
		operator, found := findOperatorInMap(symbol, surroundOperatorMap)
		if !found {
			return equationOperator{}, errors.New("cannot find surround operator with symbol " + symbol)
		}

		return operator, nil
	}

	operator, found := findOperatorInMap(symbol, prefixOperatorMap)
	if !found {
		return equationOperator{}, errors.New("cannot find prefix operator with symbol " + symbol)
	}

	return operator, nil
}

func createGraph(
	markedExpressionParts chan markedExpression,
	surroundOperatorMap map[string]equationOperator,
	prefixOperatorMap map[string]equationOperator,
	infixOperatorMap map[string]equationOperator,
) (equationGraph, error) {
	graph := newEquationTree()
	operatorNodeIdStack := newStack()

	var prevNodeId uint
	var currentGraphLevel uint = 0
	var uid uint = 0

	for markedExpressionPart := range markedExpressionParts {
		if markedExpressionPart.contentType == mathSeparator {
			continue
		}

		contentType := markedExpressionPart.contentType

		if operatorNodeIdStack.size() > 0 {
			err := applyArgsToPreviousOperator(&operatorNodeIdStack, &graph, uid, contentType)
			if err != nil {
				return graph, err
			}
		}

		if contentType == mathNumber {
			converted, err := strconv.ParseFloat(markedExpressionPart.content, 64)
			if err != nil {
				converted = 0
			}

			graph.upsert(uid, mathNumberNode{
				id:         uid,
				value:      converted,
				graphLevel: currentGraphLevel,
			})
		}

		if contentType == mathSymbol {
			operator, err := findOperatorForMathSymbol(
				markedExpressionPart.content,
				infixOperatorMap,
				prefixOperatorMap,
			)
			if err != nil {
				return graph, err
			}

			var operatorArgs []uint
			if operator.placeType == infixOperator {
				operatorArgs = append(operatorArgs, prevNodeId)
			}

			graph.upsert(uid, mathOperatorNode{
				id:           uid,
				operator:     operator,
				operatorArgs: operatorArgs,
				graphLevel:   currentGraphLevel,
			})

			operatorNodeIdStack.push(uid)
		}

		if contentType == mathSurroundStart {
			var previousOperatorNode mathOperatorNode
			var isPreviousNodeStickOrSurroundOperator bool

			previousOperatorNodeId, ok := operatorNodeIdStack.peek()
			if ok {
				operatorNode, err := graph.getOperatorNode(previousOperatorNodeId)
				if err != nil {
					return graph, err
				}

				previousOperatorNode = operatorNode
				previousNode, err := graph.getNode(previousOperatorNodeId)
				if err != nil {
					return graph, err
				}

				isPreviousNodeStick := previousNode.equationNodeType() == graphOperatorStickNode
				isPreviousNodeSurroundOperator := previousOperatorNode.operator.placeType == surroundOperator

				isPreviousNodeStickOrSurroundOperator = isPreviousNodeStick || isPreviousNodeSurroundOperator
			}

			isSymbolSurroundOperator := !ok || (ok && isPreviousNodeStickOrSurroundOperator)

			operator, err := findOperatorForMathSurround(
				markedExpressionPart.content,
				isSymbolSurroundOperator,
				surroundOperatorMap,
				prefixOperatorMap,
			)
			if err != nil {
				return graph, err
			}

			if ok && previousOperatorNode.operator.placeType == infixOperator {
				return graph, errors.New("expected previous operator not to be a infix operator")
			}

			if isSymbolSurroundOperator {
				graph.upsert(uid, mathOperatorNode{
					id:           uid,
					operator:     operator,
					operatorArgs: []uint{},
					graphLevel:   currentGraphLevel,
				})
				currentGraphLevel = currentGraphLevel + 1

				operatorNodeIdStack.push(uid)
			} else {
				graph.upsert(uid, mathOperatorStickerNode{
					id: uid,
				})
			}

		}

		if contentType == mathSurroundEnd {
			currentGraphLevel = currentGraphLevel - 1
			lastOperatorId, _ := operatorNodeIdStack.pop()
			lastNode, err := graph.getNode(lastOperatorId)
			if err != nil {
				return graph, err
			}

			if lastNode.equationNodeType() == graphOperatorStickNode {
				lastOperatorId, _ = operatorNodeIdStack.pop()
			}

			prevNodeId = lastOperatorId
		} else {
			prevNodeId = uid
		}

		uid = uid + 1
	}

	return graph, nil
}
