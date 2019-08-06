package equation

import (
	"strconv"
)

func selectBlock(reader Reader) string {
	result := ""
	count := 0

	current, _ := reader(1, true)
	if current != "(" {
		reader(1, false)
		return current
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

		if is := isNumber(current); is {
			continue
		}

		prev, _ := reader(-1, true)
		// next, _ := reader(1, true)

		// if (isNumber(prev) || prev == ")") && (isNumber(next) || next == "(") {
		if isNumber(prev) || prev == ")" {
			result = append(result, operator{
				symbol:          current,
				innerExpression: "",
				index:           index,
			})
		} else {
			if current == "(" {
				reader(-1, false)
			}

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
