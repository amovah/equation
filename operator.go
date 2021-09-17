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

var surroundSignList = []equationSurroundSign{
	parenthesisSurroundSign,
	bracketSurroundSign,
	curlyBracketSurroundSign,
}

func generateInfixOperatorMap(infixOperatorList []equationOperator) map[string]equationOperator {
	result := make(map[string]equationOperator)

	for _, operator := range infixOperatorList {
		result[operator.symbol] = operator
	}

	return result
}

var infixOperatorMap = generateInfixOperatorMap(defaultInfixOperatorList)

func generatePrefixOperatorMap(prefixOperatorList []equationOperator) map[string]equationOperator {
	result := make(map[string]equationOperator)

	for _, operator := range prefixOperatorList {
		result[operator.symbol] = operator
	}

	return result
}

var prefixOperatorMap = generatePrefixOperatorMap(defaultPrefixOperatorList)

func generateSurroundOperatorMap(surroundOperatorList []equationOperator) map[string]equationOperator {
	result := make(map[string]equationOperator)

	for _, operator := range surroundOperatorList {
		result[operator.surroundSign.start] = operator
		result[operator.surroundSign.end] = operator
	}

	return result
}

var surroundOperatorMap = generateSurroundOperatorMap(defaultSurroundOperatorList)

func generateSurroundSignMap(surroundSigns []equationSurroundSign) (map[string]bool, map[string]bool) {
	startSignMap := make(map[string]bool)
	endSignMap := make(map[string]bool)
	for _, surroundSign := range surroundSigns {
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
