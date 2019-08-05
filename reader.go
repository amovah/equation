package equation

import (
	"regexp"
)

func splitter(str string) []string {
	reg := regexp.MustCompile(`\d+|\W|\w+`)

	return reg.FindAllString(str, -1)
}

func nextElement(arr []string, index int) string {
	if index+1 >= len(arr) {
		return ""
	}

	return arr[index+1]
}

func prevElement(arr []string, index int) string {
	if index-1 < 0 || len(arr) == 0 {
		return ""
	}

	return arr[index-1]
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
