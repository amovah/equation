package equation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	assert.Equal(t, 4.0, Solve("2+2"))
}
