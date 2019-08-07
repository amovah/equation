package equation

import (
	"equation/operators"
	"strconv"
	"strings"
)

func gather(str []string, from int, to int) string {
	return strings.Join(str[from:to], "")
}

func splitIntoTwo(str []string, index int) (string, string) {
	return gather(str, 0, index), gather(str, index+1, len(str))
}

func commaHandler(str string, operators map[string]operators.Operator) []float64 {
	extracted := extractOperators(createReader(str))
	result := make([]float64, 0)
	lastIndex := 0
	for _, v := range extracted {
		if v.symbol == "," {
			result = append(result, calculate(gather(splitter(str), lastIndex, v.index), operators))
			lastIndex = v.index + 1
		}
	}

	if len(result) == 0 {
		return []float64{calculate(str, operators)}
	}

	return result
}

func min(arr []sign, operators map[string]operators.Operator) sign {
	min := arr[0]
	minPriority := operators[min.symbol].Priority

	for _, v := range arr {
		if operators[v.symbol].Priority < minPriority {
			min = v
			minPriority = operators[v.symbol].Priority
		}
	}

	return min
}

func calculate(str string, operators map[string]operators.Operator) float64 {
	splitted := splitter(str)
	extracted := extractOperators(createReader(str))
	if len(extracted) == 0 {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0.0
		}

		return num
	}

	hasInner := make([]sign, 0)
	withoutInner := make([]sign, 0)
	for _, v := range extracted {
		if v.innerExpression == "" {
			withoutInner = append(withoutInner, v)
		} else {
			hasInner = append(hasInner, v)
		}
	}

	var low sign
	if len(withoutInner) > 0 {
		low = min(withoutInner, operators)
	} else {
		low = min(hasInner, operators)
	}

	if low.innerExpression == "" {
		left, right := splitIntoTwo(splitted, low.index)
		return operators[low.symbol].Operation(
			calculate(left, operators),
			calculate(right, operators),
		)
	} else {
		return operators[low.symbol].Operation(
			commaHandler(low.innerExpression, operators)...,
		)
	}
}
