package equation

import (
	"strconv"
	"strings"

	"github.com/amovah/equation/operators"
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

func max(arr []sign, operators map[string]operators.Operator) sign {
	max := arr[0]
	maxPriority := operators[max.symbol].Priority

	for _, v := range arr {
		if operators[v.symbol].Priority < maxPriority {
			max = v
			maxPriority = operators[v.symbol].Priority
		}
	}

	return max
}

func replaceWith(org []string, from, to int, with string) []string {
	result := make([]string, 0)
	for i, v := range org {
		if i >= from && i < to {
			continue
		}

		if i == to {
			result = append(result, with)
			continue
		}

		result = append(result, v)
	}

	return result
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
		low = max(withoutInner, operators)
	} else {
		low = max(hasInner, operators)
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
