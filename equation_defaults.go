package equation

var parenthesisSurroundSign = equationSurroundSign{
	start: "(",
	end:   ")",
}

var bracketSurroundSign = equationSurroundSign{
	start: "[",
	end:   "]",
}

var curlyBracketSurroundSign = equationSurroundSign{
	start: "{",
	end:   "}",
}

var defaultInfixOperatorList = []equationOperator{
	{
		symbol:     "+",
		precedence: 10,
		operation: func(nums ...float64) float64 {
			return nums[0] + nums[1]
		},
		placeType: infixOperator,
	},
	{
		symbol:     "-",
		precedence: 10,
		operation: func(nums ...float64) float64 {
			return nums[0] - nums[1]
		},
		placeType: infixOperator,
	},
	{
		symbol:     "*",
		precedence: 11,
		operation: func(nums ...float64) float64 {
			return nums[0] * nums[1]
		},
		placeType: infixOperator,
	},
}

var defaultPrefixOperatorList = []equationOperator{
	{
		symbol:       "plusTwo",
		surroundSign: parenthesisSurroundSign,
		operation: func(nums ...float64) float64 {
			return nums[0] + 2
		},
		placeType: prefixOperator,
	},
}

var defaultSurroundOperatorList = []equationOperator{
	{
		surroundSign: parenthesisSurroundSign,
		operation: func(nums ...float64) float64 {
			return nums[0]
		},
		placeType:  surroundOperator,
		precedence: 20,
	},
}
