package equation

import (
	"strconv"
)

func selectBlock(reader Reader) string {
	result := ""
	count := 0

	next, _ := reader(2, true)
	if isNumber(next) {
		return next
	}

	for {
		next, _ := reader(1, false)

		if next == "" {
			break
		}

		if next == ")" {
			count = count - 1
			if count == 0 {
				break
			}
			result = result + next
		} else if next == "(" {
			if count != 0 {
				result = result + next
			}
			count = count + 1
		} else if count != 0 {
			result = result + next
		}
	}

	return result
}

func isNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

type operator struct {
	symbol          string
	innerExpression string
	index           int
}

func extractOperators(reader Reader) []operator {
	result := make([]operator, 0)

	for {
		current, index := reader(1, false)
		if current == "" {
			break
		}

		if _, err := strconv.ParseFloat(current, 64); err == nil {
			continue
		}

		prev, _ := reader(-1, true)
		next, _ := reader(1, true)

		if (isNumber(prev) || prev == ")") && (isNumber(next) || next == "(") {
			result = append(result, operator{
				symbol:          current,
				innerExpression: "",
				index:           index,
			})
		} else {
			inner := selectBlock(reader)
			result = append(result, operator{
				symbol:          current,
				innerExpression: inner,
				index:           index,
			})
		}
	}

	return result
}
