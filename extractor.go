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

type sign struct {
	symbol          string
	innerExpression string
	startIndex      int
	endIndex        int
}

func extractOperators(reader Reader) []sign {
	result := make([]sign, 0)

	for {
		current, index := reader(1, false)
		if current == "" {
			break
		}

		if is := isNumber(current); is {
			continue
		}

		prev, _ := reader(-1, true)
		if isNumber(prev) || prev == ")" {
			result = append(result, sign{
				symbol:          current,
				innerExpression: "",
				startIndex:      index,
				endIndex:        index,
			})
		} else {
			if current == "(" {
				reader(-1, false)
			}

			inner := selectBlock(reader)
			_, end := reader(0, true)
			result = append(result, sign{
				symbol:          current,
				innerExpression: inner,
				startIndex:      index,
				endIndex:        end,
			})
		}
	}

	return result
}
