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

func markEquationPart(str string) chan markedExpression {
	ch := make(chan markedExpression, 0)
	startSurroundSignMap, endSurroundSignMap := generateSurroundSignMap()

	go func() {
		splatted := splitter(str)
		for _, value := range splatted {
			if isNumber(value) {
				ch <- markedExpression{
					content:     value,
					contentType: mathNumber,
				}
			} else if startSurroundSignMap[value] {
				ch <- markedExpression{
					content:     value,
					contentType: mathSurroundStart,
				}
			} else if endSurroundSignMap[value] {
				ch <- markedExpression{
					content:     value,
					contentType: mathSurroundEnd,
				}
			} else if value == separateOperator {
				ch <- markedExpression{
					content:     value,
					contentType: mathSeparator,
				}
			} else {
				ch <- markedExpression{
					content:     value,
					contentType: mathSymbol,
				}
			}
		}

		close(ch)
	}()

	return ch
}
