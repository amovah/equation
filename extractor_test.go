package equation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectBlock(t *testing.T) {
	reader := createReader("((2+3),2^3+(2))")
	assert.Equal(t, "(2+3),2^3+(2)", selectBlock(reader))

	reader = createReader("2+log(5,10)+6")
	assert.Equal(t, "5,10", selectBlock(reader))

	reader = createReader("sin(4)")
	assert.Equal(t, "4", selectBlock(reader))
}
