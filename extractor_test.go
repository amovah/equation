package equation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectBlock(t *testing.T) {
	reader := createReader("((2+3),2^3+(2))")
	assert.Equal(t, "(2+3),2^3+(2)", selectBlock(reader))
	_, index := reader(0, true)
	assert.Equal(t, 14, index)

	reader = createReader("(5,10)+6")
	assert.Equal(t, "5,10", selectBlock(reader))
	_, index = reader(0, true)
	assert.Equal(t, 4, index)

	reader = createReader("5+sigma(10,20,30)")
	assert.Equal(t, "5", selectBlock(reader))
}

func TestIsNumber(t *testing.T) {
	assert.Equal(t, false, isNumber("+()"))
	assert.Equal(t, false, isNumber("sin"))
	assert.Equal(t, true, isNumber("4.0"))
	assert.Equal(t, false, isNumber("(5"))
	assert.Equal(t, true, isNumber("20031"))
	assert.Equal(t, false, isNumber(""))
}

func TestExractOperator(t *testing.T) {
	operators := extractOperators(createReader("(2.1234)+((5))"))
	assert.Equal(t, 3, len(operators))
	assert.Equal(t, "(", operators[0].symbol)
	assert.Equal(t, "2.1234", operators[0].innerExpression)
	assert.Equal(t, 0, operators[0].index)
	assert.Equal(t, "+", operators[1].symbol)
	assert.Equal(t, "", operators[1].innerExpression)
	assert.Equal(t, 3, operators[1].index)
	assert.Equal(t, "(", operators[2].symbol)
	assert.Equal(t, "(5)", operators[2].innerExpression)
	assert.Equal(t, 4, operators[2].index)

	operators = extractOperators(createReader("+5-log(2,10)"))
	assert.Equal(t, 3, len(operators))
	assert.Equal(t, "+", operators[0].symbol)
	assert.Equal(t, "5", operators[0].innerExpression)
	assert.Equal(t, 0, operators[0].index)
	assert.Equal(t, "-", operators[1].symbol)
	assert.Equal(t, "", operators[1].innerExpression)
	assert.Equal(t, 2, operators[1].index)
	assert.Equal(t, "log", operators[2].symbol)
	assert.Equal(t, "2,10", operators[2].innerExpression)
	assert.Equal(t, 3, operators[2].index)

	operators = extractOperators(createReader("+(-5)"))
	assert.Equal(t, 1, len(operators))
	assert.Equal(t, "+", operators[0].symbol)
	assert.Equal(t, "-5", operators[0].innerExpression)
	assert.Equal(t, 0, operators[0].index)

	operators = extractOperators(createReader("12+13*4-6"))
	assert.Equal(t, 3, len(operators))
	assert.Equal(t, "+", operators[0].symbol)
	assert.Equal(t, "", operators[0].innerExpression)
	assert.Equal(t, 1, operators[0].index)
	assert.Equal(t, "*", operators[1].symbol)
	assert.Equal(t, "", operators[1].innerExpression)
	assert.Equal(t, 3, operators[1].index)
	assert.Equal(t, "-", operators[2].symbol)
	assert.Equal(t, "", operators[2].innerExpression)
	assert.Equal(t, 5, operators[2].index)

	operators = extractOperators(createReader("12+13*(4-(6))"))
	assert.Equal(t, 3, len(operators))
	assert.Equal(t, "+", operators[0].symbol)
	assert.Equal(t, "", operators[0].innerExpression)
	assert.Equal(t, 1, operators[0].index)
	assert.Equal(t, "*", operators[1].symbol)
	assert.Equal(t, "", operators[1].innerExpression)
	assert.Equal(t, 3, operators[1].index)
	assert.Equal(t, "(", operators[2].symbol)
	assert.Equal(t, "4-(6)", operators[2].innerExpression)
	assert.Equal(t, 4, operators[2].index)
}
