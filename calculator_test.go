package equation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGatherBothSide(t *testing.T) {
	left, right := gatherBothSide(splitter("2+23.1+6"), 1)
	assert.Equal(t, "2", left)
	assert.Equal(t, "23.1+6", right)

	left, right = gatherBothSide((splitter("12+13*(4-(6))")), 3)
	assert.Equal(t, "12+13", left)
	assert.Equal(t, "(4-(6))", right)
}
