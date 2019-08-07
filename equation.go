package equation

import (
	"equation/constants"
	"equation/operators"
	"fmt"
	"strings"
)

func removeSpace(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func replaceConsts(str string, consts []constants.Constant) string {
	result := str

	for _, v := range consts {
		result = strings.ReplaceAll(result, v.Symbol, "("+fmt.Sprint(v.Value)+")")
	}

	return result
}

func Solve(expression string, ops map[string]operators.Operator, consts []constants.Constant) float64 {
	return calculate(replaceConsts(removeSpace(expression), consts), ops)
}
