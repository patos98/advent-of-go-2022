package monkeysareback

import (
	"aoc-2022-go/21_monkeys_are_back/monkey"
	"aoc-2022-go/utils"
)

const INPUT_PATH = "21_monkeys_are_back/input.txt"
const TEST_INPUT_PATH = "21_monkeys_are_back/input_test.txt"

func GetRootNumber(inputPath string) int {
	monkeys := parseMonkeys(inputPath)
	return monkeys["root"].GetNumber(monkeys)
}

func parseMonkeys(inputPath string) map[string]monkey.Monkey {
	monkeys := map[string]monkey.Monkey{}
	utils.ProcessInputLines(inputPath, func(s string) {
		m := monkey.Parse(s)
		monkeys[m.GetName()] = m
	})
	return monkeys
}
