package distresssignal

import (
	"aoc-2022-go/13_distresssignal/comparator"
	intorlist "aoc-2022-go/13_distresssignal/types/intorlist"
	"aoc-2022-go/utils"
	"sort"
)

const INPUT_PATH = "13_distresssignal/input.txt"
const TEST_INPUT_PATH = "13_distresssignal/input_test.txt"

func GetSumOfCorrectPairIndexes(inputPath string) int {
	var left *intorlist.IntOrList
	var right *intorlist.IntOrList

	sumOfCorrectPairIndexes := 0

	for i, line := range utils.GetAllInputLines(inputPath) {
		if i%3 == 0 {
			left = parse(line)
		} else if i%3 == 1 {
			right = parse(line)
			if comparator.IsPairInCorrectOrder(left, right) {
				sumOfCorrectPairIndexes += i/3 + 1
			}
		}
	}

	return sumOfCorrectPairIndexes
}

type LineWithItem struct {
	line string
	item *intorlist.IntOrList
}

func GetDecoderKey(inputPath string) int {
	packets := []LineWithItem{}
	utils.ProcessInputLines(inputPath, func(line string) {
		if line == "" {
			return
		}

		packets = append(packets, LineWithItem{
			line: line,
			item: parse(line),
		})
	})

	signal2 := parse("[[2]]")
	signal6 := parse("[[6]]")

	packets = append(packets, LineWithItem{
		item: signal2,
		line: "[[2]]",
	})
	packets = append(packets, LineWithItem{
		item: signal6,
		line: "[[6]]",
	})

	sort.Slice(packets, func(i, j int) bool {
		return comparator.IsPairInCorrectOrder(packets[i].item, packets[j].item)
	})

	product := 1
	for i, packet := range packets {
		if packet.item == signal2 || packet.item == signal6 {
			product *= i + 1
		}
	}

	return product
}
