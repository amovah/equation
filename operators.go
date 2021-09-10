package equation

const noPrecedence = 0

type equationPlaceType int8

const (
	infixOperator equationPlaceType = iota
	prefixOperator
	surroundOperator
	argSplitterOperator
)

type equationOperation func(nums ...float64) float64

type equationOperator struct {
	symbol       string
	surroundSign equationSurroundSign
	precedence   uint16
	operation    equationOperation
	placeType    equationPlaceType
}

type equationSurroundSign struct {
	start string
	end   string
}

func parenthesisSurround() equationSurroundSign {
	return equationSurroundSign{
		start: "(",
		end:   ")",
	}
}

var defaultOperationList = []equationOperator{
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
	{
		surroundSign: parenthesisSurround(),
		operation: func(nums ...float64) float64 {
			return nums[0] / nums[1]
		},
		placeType: surroundOperator,
	},
	{
		symbol:       "tavan",
		surroundSign: parenthesisSurround(),
		operation: func(nums ...float64) float64 {
			return nums[0] + 2
		},
		placeType: prefixOperator,
	},
}
