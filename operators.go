package equation

const numberPrecedence = 0

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

func equationDefaultOperators() []equationOperator {
	return []equationOperator{
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
			precedence:   12,
			operation: func(nums ...float64) float64 {
				return nums[0] / nums[1]
			},
			placeType: surroundOperator,
		},
	}
}

func defaultSurroundSigns() []equationSurroundSign {
	return []equationSurroundSign{parenthesisSurround()}
}
