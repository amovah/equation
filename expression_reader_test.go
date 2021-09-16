package equation

import (
	"fmt"
	"testing"
)

func TestStringReader(t *testing.T) {
	//x := splitter("2+2+log(5, 55, 36)")

	err := createGraph(markMe("10*((3+2))+2"))
	if err != nil {
		fmt.Println(err)
	}
}
