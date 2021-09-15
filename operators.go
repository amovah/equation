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

var surroundSignList = []equationSurroundSign{
	parenthesisSurroundSign,
	bracketSurroundSign,
	curlyBracketSurroundSign,
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
		surroundSign: parenthesisSurroundSign,
		operation: func(nums ...float64) float64 {
			return nums[0] / nums[1]
		},
		placeType: surroundOperator,
	},
	{
		symbol:       "tavan",
		surroundSign: parenthesisSurroundSign,
		operation: func(nums ...float64) float64 {
			return nums[0] + 2
		},
		placeType: prefixOperator,
	},
}

func surroundSignMap() (map[string]bool, map[string]bool) {
	startSignMap := make(map[string]bool)
	endSignMap := make(map[string]bool)
	for _, surroundSign := range surroundSignList {
		startSignMap[surroundSign.start] = true
		endSignMap[surroundSign.end] = true
	}

	return startSignMap, endSignMap
}
