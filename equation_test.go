package equation

import (
	"fmt"
	"testing"

	"github.com/amovah/equation/constants"
	"github.com/amovah/equation/operators"

	"github.com/stretchr/testify/assert"
)

func TestRemoveSpace(t *testing.T) {
	assert.Equal(t, "foobar()", removeSpace(" foo bar( )"))
	assert.Equal(t, "ab.c-(5+)", removeSpace("a b.c - (5+)"))
}

func TestRepalceConsts(t *testing.T) {
	consts := constants.Defaults()
	pi := constants.Pi()

	assert.Equal(t, "("+fmt.Sprint(pi.Value)+")", replaceConsts("pi", consts))
	assert.Equal(
		t,
		"c+3+log((5+6),("+fmt.Sprint(pi.Value)+"))",
		replaceConsts("c+3+log((5+6),pi)", consts),
	)
}

func TestFirst(t *testing.T) {
	ops := operators.Defaults()
	consts := constants.Defaults()

	assert.Equal(t, 4.0, Solve("2+2", ops, consts))
	assert.Equal(t, 22.718281828459045, Solve("e+20", ops, consts))
	assert.Equal(t, -6.0, Solve("+(-(+(-(-6))))", ops, consts))
	assert.Equal(t, 12.24, Solve("+12+0.24", ops, consts))
	assert.Equal(t, 8.0, Solve("8-2^3+8", ops, consts))
}
