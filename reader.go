package equation

import (
	"regexp"
)

func splitter(str string) []string {
	reg := regexp.MustCompile(`\d+\.\d+|\W|\w+`)
	return reg.FindAllString(str, -1)
}

func nextElement(arr []string, index int) string {
	if index+1 >= len(arr) || index+1 < 0 {
		return ""
	}

	return arr[index+1]
}

func prevElement(arr []string, index int) string {
	if index-1 < 0 || index-1 >= len(arr) {
		return ""
	}

	return arr[index-1]
}

type Reader func(step int, peek bool) (string, int)

func createReader(str []string) Reader {
	current := -1

	return func(step int, peek bool) (string, int) {
		if peek == false {
			defer func() {
				current = current + step
			}()
		}

		if step >= 0 {
			return nextElement(str, current+step-1), current + step
		}

		return prevElement(str, current+step+1), current + step
	}
}
