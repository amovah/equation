package equation

func gatherBothSide(str []string, index int) (string, string) {
	left := ""
	right := ""
	direction := false

	for i, v := range str {
		if i == index {
			direction = true
			continue
		}

		if direction {
			right = right + v
		} else {
			left = left + v
		}
	}

	return left, right
}

// func calculate(str string, userOperators map[string]operators.Operator) float64 {
// 	splitted := splitter(str)
// 	extracted := extractOperators(createReader(str))
// 	if len(extracted) == 0 {
// 		num, err := strconv.ParseFloat(str, 64)
// 		if err != nil {
// 			return 0.0
// 		}

// 		return num
// 	}

// 	lowPriority := extracted[0]
// 	priority := userOperators[lowPriority.symbol].Priority
// 	for _, v := range extracted {
// 		if userOperators[v.symbol].Priority < priority {
// 			lowPriority = v
// 			priority = userOperators[v.symbol].Priority
// 		}
// 	}

// 	if lowPriority.innerExpression == "" {

// 	} else {
// 		return userOperators[lowPriority.symbol].Operation(lowPriority.innerExpression)
// 	}
// }
