package equation

func selectBlock(reader Reader) string {
	result := ""
	count := 1

	for {
		next := reader(1, true)

		if next == ")" {
			count = count - 1
			if count == 0 {
				break
			}
		} else if next == "(" {
			count = count + 1
		}

		result = result + next
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
