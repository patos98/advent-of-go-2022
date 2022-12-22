package monkey

type MonkeyWithOperation struct {
	name               string
	monkeyLeft         string
	monkeyRight        string
	operatorString     string
	antiOperatorString string
	operator           Operator
	antiOperation      AntiOperation
}

func (m MonkeyWithOperation) GetName() string {
	return m.name
}

func (m MonkeyWithOperation) GetMonkeyLeft() string {
	return m.monkeyLeft
}

func (m MonkeyWithOperation) GetMonkeyRight() string {
	return m.monkeyRight
}

func (m MonkeyWithOperation) GetNumber(monkeys map[string]Monkey) (int, bool) {
	leftNumber, isLeftHuman := monkeys[m.monkeyLeft].GetNumber(monkeys)
	rightNumber, isRightHuman := monkeys[m.monkeyRight].GetNumber(monkeys)
	return m.operator(leftNumber, rightNumber), isLeftHuman || isRightHuman
}

func (m MonkeyWithOperation) GetMissingNumberForResult(monkeys map[string]Monkey, result int) int {
	leftMonkey := monkeys[m.monkeyLeft]
	rightMonkey := monkeys[m.monkeyRight]

	leftNumber, isLeftHuman := leftMonkey.GetNumber(monkeys)
	rightNumber, isRightHuman := rightMonkey.GetNumber(monkeys)

	if isLeftHuman {
		subResult := m.antiOperation.ToGetLeftOperand(AntiOperatorParams{
			result:       result,
			knownOperand: rightNumber,
		})

		if leftMonkey.GetName() == HUMAN_MONKEY_NAME {
			return subResult
		}

		return leftMonkey.GetMissingNumberForResult(monkeys, subResult)
	}

	if isRightHuman {
		subResult := m.antiOperation.ToGetRightOperand(AntiOperatorParams{
			result:       result,
			knownOperand: leftNumber,
		})

		if rightMonkey.GetName() == HUMAN_MONKEY_NAME {
			return subResult
		}

		return rightMonkey.GetMissingNumberForResult(monkeys, subResult)
	}

	// should not be reached
	return 0
}
