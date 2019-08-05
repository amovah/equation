package equation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitter1(t *testing.T) {
	all := Splitter("23+log(10,5)")

	assert.Equal(t, "23", all[0])
	assert.Equal(t, "+", all[1])
	assert.Equal(t, "log", all[2])
	assert.Equal(t, "(", all[3])
	assert.Equal(t, "10", all[4])
	assert.Equal(t, ",", all[5])
	assert.Equal(t, "5", all[6])
	assert.Equal(t, ")", all[7])
}

func TestSplitter2(t *testing.T) {
	all := Splitter("+(-4)")

	assert.Equal(t, "+", all[0])
	assert.Equal(t, "(", all[1])
	assert.Equal(t, "-", all[2])
	assert.Equal(t, "4", all[3])
	assert.Equal(t, ")", all[4])
}

func TestSplitter3(t *testing.T) {
	all := Splitter("2+e")

	assert.Equal(t, "2", all[0])
	assert.Equal(t, "+", all[1])
	assert.Equal(t, "e", all[2])
}
