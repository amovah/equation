package operators

type Operator struct {
	symbol    string
	operation func(...float64) float64
	priority  int
}

func Plus() Operator {
	return Operator{
		symbol: "+",
		operation: func(nums ...float64) float64 {
			if len(nums) == 1 {
				return nums[0]
			}

			return nums[0] + nums[1]
		},
		priority: 0,
	}
}

func Minus() Operator {
	return Operator{
		symbol: "-",
		operation: func(nums ...float64) float64 {
			if len(nums) == 1 {
				return -1 * nums[0]
			}

			return nums[0] - nums[1]
		},
		priority: 0,
	}
}

func Multiplication() Operator {
	return Operator{
		symbol: "*",
		operation: func(nums ...float64) float64 {
			return nums[0] * nums[1]
		},
		priority: 1,
	}
}

func Division() Operator {
	return Operator{
		symbol: "/",
		operation: func(nums ...float64) float64 {
			return nums[0] / nums[1]
		},
		priority: 1,
	}
}

func Parenthes() Operator {
	return Operator{
		symbol: "(",
		operation: func(nums ...float64) float64 {
			return nums[0]
		},
		priority: 2,
	}
}

func AllDefaults() []Operator {
	return []Operator{
		Plus(),
		Minus(),
		Multiplication(),
		Division(),
		Parenthes(),
	}
}
