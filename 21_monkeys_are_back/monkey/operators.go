package monkey

type Operator func(int, int) int

type AntiOperatorParams struct {
	result       int
	knownOperand int
}

type AntiOperator func(AntiOperatorParams) int

type AntiOperation struct {
	ToGetLeftOperand  AntiOperator
	ToGetRightOperand AntiOperator
}

var operators = map[string]Operator{
	"+": func(i1, i2 int) int { return i1 + i2 },
	"-": func(i1, i2 int) int { return i1 - i2 },
	"*": func(i1, i2 int) int { return i1 * i2 },
	"/": func(i1, i2 int) int { return i1 / i2 },
}

var antiOperations = map[string]AntiOperation{
	"+": {
		ToGetLeftOperand:  func(aop AntiOperatorParams) int { return aop.result - aop.knownOperand },
		ToGetRightOperand: func(aop AntiOperatorParams) int { return aop.result - aop.knownOperand },
	},
	"-": {
		ToGetLeftOperand:  func(aop AntiOperatorParams) int { return aop.result + aop.knownOperand },
		ToGetRightOperand: func(aop AntiOperatorParams) int { return aop.knownOperand - aop.result },
	},
	"*": {
		ToGetLeftOperand:  func(aop AntiOperatorParams) int { return aop.result / aop.knownOperand },
		ToGetRightOperand: func(aop AntiOperatorParams) int { return aop.result / aop.knownOperand },
	},
	"/": {
		ToGetLeftOperand:  func(aop AntiOperatorParams) int { return aop.result * aop.knownOperand },
		ToGetRightOperand: func(aop AntiOperatorParams) int { return aop.knownOperand / aop.result },
	},
}
