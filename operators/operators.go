package operators

import "math"

type Operator struct {
	Symbol    string
	Operation func(...float64) float64
	Priority  int
}

func Plus() Operator {
	return Operator{
		Symbol: "+",
		Operation: func(nums ...float64) float64 {
			if len(nums) == 1 {
				return nums[0]
			}

			return nums[0] + nums[1]
		},
		Priority: 0,
	}
}

func Minus() Operator {
	return Operator{
		Symbol: "-",
		Operation: func(nums ...float64) float64 {
			if len(nums) == 1 {
				return 0 - nums[0]
			}

			return nums[0] - nums[1]
		},
		Priority: 1,
	}
}

func Multiplication() Operator {
	return Operator{
		Symbol: "*",
		Operation: func(nums ...float64) float64 {
			return nums[0] * nums[1]
		},
		Priority: 2,
	}
}

func Division() Operator {
	return Operator{
		Symbol: "/",
		Operation: func(nums ...float64) float64 {
			return nums[0] / nums[1]
		},
		Priority: 2,
	}
}

func Parenthes() Operator {
	return Operator{
		Symbol: "(",
		Operation: func(nums ...float64) float64 {
			return nums[0]
		},
		Priority: 6,
	}
}

func Power() Operator {
	return Operator{
		Symbol: "^",
		Operation: func(nums ...float64) float64 {
			return math.Pow(nums[0], nums[1])
		},
		Priority: 4,
	}
}

func Add(ops map[string]Operator, elems ...Operator) {
	for _, v := range elems {
		ops[v.Symbol] = v
	}
}

func Defaults() map[string]Operator {
	all := make(map[string]Operator)
	Add(
		all,
		Plus(),
		Minus(),
		Multiplication(),
		Division(),
		Parenthes(),
		Power(),
	)

	return all
}
