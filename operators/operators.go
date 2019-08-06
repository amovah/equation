package operators

import "math"

type Operator struct {
	symbol    string
	Operation func(...float64) float64
	Priority  int
}

func Plus() Operator {
	return Operator{
		symbol: "+",
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
		symbol: "-",
		Operation: func(nums ...float64) float64 {
			if len(nums) == 1 {
				return 0 - nums[0]
			}

			return nums[0] - nums[1]
		},
		Priority: 0,
	}
}

func Multiplication() Operator {
	return Operator{
		symbol: "*",
		Operation: func(nums ...float64) float64 {
			return nums[0] * nums[1]
		},
		Priority: 1,
	}
}

func Division() Operator {
	return Operator{
		symbol: "/",
		Operation: func(nums ...float64) float64 {
			return nums[0] / nums[1]
		},
		Priority: 1,
	}
}

func Parenthes() Operator {
	return Operator{
		symbol: "(",
		Operation: func(nums ...float64) float64 {
			return nums[0]
		},
		Priority: 4,
	}
}

func Power() Operator {
	return Operator{
		symbol: "^",
		Operation: func(nums ...float64) float64 {
			return math.Pow(nums[0], nums[1])
		},
		Priority: 3,
	}
}

func MergeOps(ops map[string]Operator, elems ...Operator) map[string]Operator {
	for _, v := range elems {
		ops[v.symbol] = v
	}

	return ops
}

func Defaults() map[string]Operator {
	all := make(map[string]Operator)
	return MergeOps(
		all,
		Plus(),
		Minus(),
		Multiplication(),
		Division(),
		Parenthes(),
		Power(),
	)
}
