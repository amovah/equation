package equation

import (
	"equation/constants"
	"equation/operators"
	"fmt"
	"testing"

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
}
