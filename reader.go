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

type Reader func(step int, peek bool) string

func createReader(str string) Reader {
	current := -1
	splitted := splitter(str)

	return func(step int, peek bool) string {
		if peek == false {
			defer func() {
				current = current + step
			}()
		}

		if step >= 0 {
			return nextElement(splitted, current+step-1)
		}

		return prevElement(splitted, current+step+1)
	}
}
