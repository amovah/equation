package equation

import (
	"equation/operators"
	"fmt"
	"strconv"
	"strings"
)

func gather(str []string, from int, to int) string {
	return strings.Join(str[from:to], "")
}

func splitIntoTwo(str []string, index int) (string, string) {
	return gather(str, 0, index), gather(str, index+1, len(str))
}

func commaHandler(str string, userOperators map[string]operators.Operator) []float64 {
	extracted := extractOperators(createReader(str))
	result := make([]float64, 0)
	lastIndex := 0
	for _, v := range extracted {
		if v.symbol == "," {
			result = append(result, calculate(gather(splitter(str), lastIndex, v.index), userOperators))
			lastIndex = v.index + 1
		}
	}

	if len(result) == 0 {
		return []float64{calculate(str, userOperators)}
	}

	return result
}

func calculate(str string, userOperators map[string]operators.Operator) float64 {
	splitted := splitter(str)
	extracted := extractOperators(createReader(str))
	if len(extracted) == 0 {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0.0
		}

		return num
	}

	lowPriority := extracted[0]
	priority := userOperators[lowPriority.symbol].Priority
	for _, v := range extracted {
		if userOperators[v.symbol].Priority < priority {
			lowPriority = v
			priority = userOperators[v.symbol].Priority
		}
	}

	if lowPriority.innerExpression == "" {
		left, right := splitIntoTwo(splitted, lowPriority.index)
		return userOperators[lowPriority.symbol].Operation(
			calculate(left, userOperators),
			calculate(right, userOperators),
		)
	} else {
		fmt.Println(lowPriority.symbol, lowPriority.innerExpression)
		return userOperators[lowPriority.symbol].Operation(
			commaHandler(lowPriority.innerExpression, userOperators)...,
		)
	}
}
