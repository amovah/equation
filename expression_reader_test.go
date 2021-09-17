package equation

import (
	"fmt"
	"testing"
)

func TestStringReader(t *testing.T) {
	//x := splitter("2+2+log(5, 55, 36)")

	_, err := createGraph(
		markEquationPart(splitter("10*((3+2))+2")),
		surroundOperatorMap,
		prefixOperatorMap,
		infixOperatorMap,
	)
	if err != nil {
		fmt.Println(err)
	}
}
