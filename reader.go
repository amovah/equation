package equation

import "regexp"

func splitter(str string) []string {
	reg := regexp.MustCompile(`\d+|\W|\w+`)

	return reg.FindAllString(str, -1)
}
