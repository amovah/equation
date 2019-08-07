package constants

import "math"

type Constant struct {
	Symbol string
	Value  float64
}

func Pi() Constant {
	return Constant{
		Symbol: "pi",
		Value:  math.Pi,
	}
}

func E() Constant {
	return Constant{
		Symbol: "e",
		Value:  math.E,
	}
}

func Add(consts map[string]Constant, elems ...Constant) {
	for _, v := range elems {
		consts[v.Symbol] = v
	}
}

func Defaults() map[string]Constant {
	all := make(map[string]Constant)
	Add(
		all,
		Pi(),
		E(),
	)

	return all
}
