package monkeysareback

import "aoc-2022-go/21_monkeys_are_back/monkey"

const INPUT_PATH = "21_monkeys_are_back/input.txt"
const TEST_INPUT_PATH = "21_monkeys_are_back/input_test.txt"

func GetRootNumber(inputPath string) int {
	monkeys := monkey.ParseMonkeys(inputPath)
	result, _ := monkeys["root"].GetNumber(monkeys)
	return result
}

func GetMissingNumberForResult(inputPath string) int {
	monkeys := monkey.ParseMonkeys(inputPath)
	rootMonkey := monkeys["root"]

	leftMonkey := monkeys[rootMonkey.GetMonkeyLeft()]
	rightMonkey := monkeys[rootMonkey.GetMonkeyRight()]

	leftNumber, isLeftHuman := leftMonkey.GetNumber(monkeys)
	rightNumber, isRightHuman := rightMonkey.GetNumber(monkeys)

	if isLeftHuman {
		return leftMonkey.GetMissingNumberForResult(monkeys, rightNumber)
	}

	if isRightHuman {
		return rightMonkey.GetMissingNumberForResult(monkeys, leftNumber)
	}

	// should not be reached
	return 0
}
