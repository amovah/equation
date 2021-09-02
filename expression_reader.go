package equation

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type mathPartType uint8

const (
	mathSymbol mathPartType = iota
	mathNumber
	mathSurroundStart
	mathSurroundEnd
)

type markedExpression struct {
	content     string
	contentType mathPartType
}

func removeSpace(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func splitter(str string) []string {
	reg := regexp.MustCompile(`\d+\.\d+|\W|\w+`)
	return reg.FindAllString(removeSpace(str), -1)
}

func isNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func markMe(str string) []markedExpression {
	result := make([]markedExpression, 0)

	splatted := splitter(str)
	for _, value := range splatted {
		if isNumber(value) {
			result = append(result, markedExpression{
				content:     value,
				contentType: mathNumber,
			})
		} else if value == "(" {
			result = append(result, markedExpression{
				content:     value,
				contentType: mathSurroundStart,
			})
		} else if value == ")" {
			result = append(result, markedExpression{
				content:     value,
				contentType: mathSurroundEnd,
			})
		} else {
			result = append(result, markedExpression{
				content:     value,
				contentType: mathSymbol,
			})
		}
	}

	return result
}

var defaultOperationList = equationDefaultOperators()

func removeAndReplace(s []mathNode, from int, to int, replaceWith mathNode) []mathNode {
	result := append(s[:from], replaceWith)
	result = append(result, s[to:]...)
	return result
}

func createGraph(parts []markedExpression) {
	tree := make([]mathNode, 0)

	for id, mathPart := range parts {
		if mathPart.contentType == mathNumber {
			converted, err := strconv.ParseFloat(mathPart.content, 64)
			if err != nil {
				converted = 0
			}

			tree = append(tree, mathNumberNode{
				id:    id,
				value: converted,
			})
		}

		if mathPart.contentType == mathSymbol {
			var operator equationOperator
			for _, defaultOperator := range defaultOperationList {
				if defaultOperator.symbol == mathPart.content {
					operator = defaultOperator
					break
				}
			}

			tree = append(tree, mathOperatorNode{
				id:       id,
				operator: operator,
			})
		}
	}

	treeOperators := make([]mathNode, 0)

	for _, node := range tree {
		if node.equationPrecedence() > numberPrecedence {
			treeOperators = append(treeOperators, node)
		}
	}

	sort.SliceStable(treeOperators, func(i, j int) bool {
		return tree[i].equationPrecedence() > tree[j].equationPrecedence()
	})

	for _, operatorNode := range treeOperators {
		var foundIndex int
		for index, node := range tree {
			if node.equationId() == operatorNode.equationId() {
				foundIndex = index
				break
			}
		}

		modified := operatorNode.equationWithSubNodes([]mathNode{tree[foundIndex-1], tree[foundIndex+1]})
		tree = removeAndReplace(tree, foundIndex-1, foundIndex+2, modified)
	}

	readNode(tree)
}

func readNode(nodes []mathNode) {
	for _, i := range nodes {
		fmt.Printf("id %v, value: %f \n", i.equationId(), i.equationValue())
		readNode(i.equationSubNodes())
	}
}
