package equation

import (
	"equation/operators"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitIntoTwo(t *testing.T) {
	left, right := splitIntoTwo(splitter("2+23.1+6"), 1)
	assert.Equal(t, "2", left)
	assert.Equal(t, "23.1+6", right)

	left, right = splitIntoTwo((splitter("12+13*(4-(6))")), 3)
	assert.Equal(t, "12+13", left)
	assert.Equal(t, "(4-(6))", right)

	left, right = splitIntoTwo(splitter("+((5))+3.14-1"), 6)
	assert.Equal(t, "+((5))", left)
	assert.Equal(t, "3.14-1", right)
}

func TestGather(t *testing.T) {
	gathered := gather(splitter("(5)+3^3-log(5,6.5)"), 0, 5)
	assert.Equal(t, "(5)+3", gathered)

	gathered = gather(splitter("(5)+3^3-log(5,6.5)"), 6, 10)
	assert.Equal(t, "3-log(", gathered)
}

func TestCalculate(t *testing.T) {
	defaultOps := operators.Defaults()

	// answer := calculate("(2+2*(-3))", defaultOps)
	// assert.Equal(t, -4.0, answer)

	// answer = calculate("+(-(-(2)))", defaultOps)
	// assert.Equal(t, 2.0, answer)

	answer := calculate("(5)+(5)", defaultOps)
	assert.Equal(t, 40.0, answer)
}
