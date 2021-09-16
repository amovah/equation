package equation

import (
	"regexp"
	"strconv"
	"strings"
)

type mathPartType uint8

const (
	mathSymbol mathPartType = iota
	mathNumber
	mathSurroundStart
	mathSurroundEnd
	mathSeparator
)

type markedExpression struct {
	content     string
	contentType mathPartType
}

func removeSpace(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func splitter(str string) []string {
	reg := regexp.MustCompile(`\d+\.\d+|\W|\w+`)
	return reg.FindAllString(removeSpace(str), -1)
}

func isNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func markMe(str string) []markedExpression {
	result := make([]markedExpression, 0)
	startSurroundSignMap, endSurroundSignMap := generateSurroundSignMap()

	splatted := splitter(str)
	for _, value := range splatted {
		if isNumber(value) {
			result = append(result, markedExpression{
				content:     value,
				contentType: mathNumber,
			})
		} else if startSurroundSignMap[value] {
			result = append(result, markedExpression{
				content:     value,
				contentType: mathSurroundStart,
			})
		} else if endSurroundSignMap[value] {
			result = append(result, markedExpression{
				content:     value,
				contentType: mathSurroundEnd,
			})
		} else if value == separateOperator {
			result = append(result, markedExpression{
				content:     value,
				contentType: mathSeparator,
			})
		} else {
			result = append(result, markedExpression{
				content:     value,
				contentType: mathSymbol,
			})
		}
	}

	return result
}
