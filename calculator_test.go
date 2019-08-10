package equation

import (
	"testing"

	"github.com/amovah/equation/operators"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	res := replaceWith([]string{"a", "b", "c", "d"}, 0, 2, "f")
	answer := []string{"f", "d"}
	assert.Equal(t, len(answer), len(res))
	for i, v := range res {
		assert.Equal(t, answer[i], v)
	}
}

func TestCalculate(t *testing.T) {
	defaultOps := operators.Defaults()

	answer := calculate(splitter("(2+2*(-3))"), defaultOps)
	assert.Equal(t, -4.0, answer)

	answer = calculate(splitter("+(-(-(2)))"), defaultOps)
	assert.Equal(t, 2.0, answer)

	answer = calculate(splitter("+(5)+(5*8-5)"), defaultOps)
	assert.Equal(t, 40.0, answer)

	answer = calculate(splitter("4^2"), defaultOps)
	assert.Equal(t, 16.0, answer)

	answer = calculate(splitter("2-(-2)"), defaultOps)
	assert.Equal(t, 4.0, answer)

	answer = calculate(splitter("-2"), defaultOps)
	assert.Equal(t, -2.0, answer)
}
