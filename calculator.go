package equation

import (
	"fmt"
	"strconv"

	"github.com/amovah/equation/operators"
)

func commaHandler(str []string, operators map[string]operators.Operator) []float64 {
	extracted := extractOperators(createReader(str))
	result := make([]float64, 0)
	lastIndex := 0
	for _, v := range extracted {
		if v.symbol == "," {
			result = append(
				result,
				calculate(
					str[lastIndex:v.startIndex],
					operators,
				),
			)
			lastIndex = v.startIndex + 1
		}
	}

	result = append(
		result,
		calculate(
			str[lastIndex:len(str)],
			operators,
		),
	)

	return result
}

func max(arr []sign, operators map[string]operators.Operator) sign {
	max := arr[0]
	maxPriority := operators[max.symbol].Priority

	for _, v := range arr {
		if operators[v.symbol].Priority > maxPriority {
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

func calculate(str []string, operators map[string]operators.Operator) float64 {
	extracted := extractOperators(createReader(str))
	if len(extracted) == 0 {
		num, err := strconv.ParseFloat(str[0], 64)
		if err != nil {
			return 0.0
		}

		return num
	}

	high := max(extracted, operators)

	if high.innerExpression == "" {
		return calculate(
			replaceWith(
				str,
				high.startIndex-1,
				high.startIndex+1,
				fmt.Sprint(
					operators[high.symbol].Operation(
						calculate(str[high.startIndex-1:high.startIndex], operators),
						calculate(str[high.startIndex+1:high.startIndex+2], operators),
					),
				),
			),
			operators,
		)
	} else {
		return calculate(
			replaceWith(
				str,
				high.startIndex,
				high.endIndex,
				fmt.Sprint(
					operators[high.symbol].Operation(
						commaHandler(splitter(high.innerExpression), operators)...,
					),
				),
			),
			operators,
		)
	}
}
