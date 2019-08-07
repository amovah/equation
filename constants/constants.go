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

func Add(consts []Constant, elems ...Constant) []Constant {
	result := consts
	for _, v := range elems {
		result = append(result, v)
	}

	return result
}

func Defaults() []Constant {
	all := make([]Constant, 0)
	return Add(
		all,
		Pi(),
		E(),
	)
}
