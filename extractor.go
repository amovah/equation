package equation

import "strconv"

func selectBlock(reader Reader) string {
	result := ""
	count := 0

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

// type operator struct {
// 	sign             string
// 	innerExperession string
// 	index            uint
// }

// func extractOperators(reader Reader) []operator {
// 	result := make([]operator, 0)

// 	for {
// 		next, index := reader(1, false)
// 		if next == "" {
// 			break
// 		}

// 		if _, err := strconv.ParseFloat(next, 64); err == nil {
// 			continue
// 		}

// 	}

// 	return result
// }