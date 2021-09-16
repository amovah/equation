package equation

const separateOperator = ","

type equationPlaceType int8

const (
	infixOperator equationPlaceType = iota
	prefixOperator
	surroundOperator
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
		placeType: surroundOperator,
	},
}

func generateInfixOperatorMap() map[string]equationOperator {
	result := make(map[string]equationOperator)

	for _, operator := range defaultInfixOperatorList {
		result[operator.symbol] = operator
	}

	return result
}

var infixOperatorMap = generateInfixOperatorMap()

func generatePrefixOperatorMap() map[string]equationOperator {
	result := make(map[string]equationOperator)

	for _, operator := range defaultPrefixOperatorList {
		result[operator.symbol] = operator
	}

	return result
}

var prefixOperatorMap = generatePrefixOperatorMap()

func generateSurroundOperatorMap() map[string]equationOperator {
	result := make(map[string]equationOperator)

	for _, operator := range defaultSurroundOperatorList {
		result[operator.surroundSign.start] = operator
		result[operator.surroundSign.end] = operator
	}

	return result
}

var surroundOperatorMap = generateSurroundOperatorMap()

func generateSurroundSignMap() (map[string]bool, map[string]bool) {
	startSignMap := make(map[string]bool)
	endSignMap := make(map[string]bool)
	for _, surroundSign := range surroundSignList {
		startSignMap[surroundSign.start] = true
		endSignMap[surroundSign.end] = true
	}

	return startSignMap, endSignMap
}

func findOperatorInMap(symbol string, operatorMap map[string]equationOperator) (equationOperator, bool) {
	if operatorMap[symbol].symbol == "" && operatorMap[symbol].surroundSign.start == "" {
		return equationOperator{}, false
	}

	return operatorMap[symbol], true
}
