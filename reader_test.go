package equation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitter(t *testing.T) {
	case1 := splitter("23+log(10,5)")
	assert.Equal(t, "23", case1[0])
	assert.Equal(t, "+", case1[1])
	assert.Equal(t, "log", case1[2])
	assert.Equal(t, "(", case1[3])
	assert.Equal(t, "10", case1[4])
	assert.Equal(t, ",", case1[5])
	assert.Equal(t, "5", case1[6])
	assert.Equal(t, ")", case1[7])

	case2 := splitter("+(-4)")
	assert.Equal(t, "+", case2[0])
	assert.Equal(t, "(", case2[1])
	assert.Equal(t, "-", case2[2])
	assert.Equal(t, "4", case2[3])
	assert.Equal(t, ")", case2[4])

	case3 := splitter("2+e")
	assert.Equal(t, "2", case3[0])
	assert.Equal(t, "+", case3[1])
	assert.Equal(t, "e", case3[2])

	case4 := splitter("((+43))")
	assert.Equal(t, "(", case4[0])
	assert.Equal(t, "(", case4[1])
	assert.Equal(t, "+", case4[2])
	assert.Equal(t, "43", case4[3])
	assert.Equal(t, ")", case4[4])
	assert.Equal(t, ")", case4[5])
}

func TestNextElement(t *testing.T) {
	assert.Equal(t, "", nextElement([]string{}, 0))
	assert.Equal(t, "", nextElement([]string{"3", "test"}, 1))
	assert.Equal(t, "foo", nextElement([]string{"bar", "bax", "foo", "feet"}, 1))
	assert.Equal(t, "", nextElement([]string{"bar"}, 1))
	assert.Equal(t, "", nextElement([]string{}, -1))
	assert.Equal(t, "dude", nextElement([]string{"dude"}, -1))
}

func TestPrevElement(t *testing.T) {
	assert.Equal(t, "", prevElement([]string{}, 0))
	assert.Equal(t, "", prevElement([]string{"3", "test"}, 0))
	assert.Equal(t, "bar", prevElement([]string{"bar", "bax", "foo"}, 1))
	assert.Equal(t, "", prevElement([]string{"bar"}, 0))
	assert.Equal(t, "", prevElement([]string{}, 1))
	assert.Equal(t, "dude", prevElement([]string{"dude"}, 1))
}

func TestReaderStream(t *testing.T) {
	reader := readStream("23+(-4)")

	current := reader(1, false)
	assert.Equal(t, "23", current)

	next := reader(1, true)
	current = reader(1, false)
	assert.Equal(t, "+", next)
	assert.Equal(t, "+", current)

	prev := reader(-1, true)
	current = reader(0, false)
	assert.Equal(t, "23", prev)
	assert.Equal(t, "+", current)

	current = reader(2, false)
	assert.Equal(t, "-", current)

	current = reader(-1, false)
	assert.Equal(t, "(", current)
}
