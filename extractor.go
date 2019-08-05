package equation

func selectBlock(reader Reader) string {
	result := ""
	count := 0

	for {
		next := reader(1, false)

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

// type cookedExpression struct {
// 	sign             string
// 	innerExperession string
// 	index            uint
// }

// func readExp(str string) []cookedExpression {
// 	result := make([]cookedExpression, 0)
// 	splitted := splitter(str)
// 	passed := 0

// 	for i, v := range splitted {
// 		if _, err := strconv.ParseFloat(v, 64); err == nil {
// 			passed += len(v)
// 			continue
// 		}
// 	}

// 	return result
// }
