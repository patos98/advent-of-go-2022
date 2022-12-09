package rucksack

import (
	"aoc-2022-go/utils"
	"strings"
	"unicode/utf8"
)

const INPUT_PATH = "03_rucksack/input.txt"

func Rucksack() int {
	prioritySum := 0

	for line := range utils.GetInputLines(INPUT_PATH) {
		lineLength := len(line)
		firstHalf := strings.Split(line, "")[:lineLength/2]
		secondHalf := strings.Split(line, "")[lineLength/2:]

		duplicatedChars := map[rune]struct{}{}
		for _, char1 := range firstHalf {
			for _, char2 := range secondHalf {
				if char1 == char2 {
					r, _ := utf8.DecodeRuneInString(char1)
					if _, contains := duplicatedChars[r]; contains {
						break
					}
					duplicatedChars[r] = struct{}{}
					prioritySum += getPriority(r)
				}
			}
		}
	}

	return prioritySum
}

func Rucksack2() int {
	prioritySum := 0

	lineCount := 0
	rucksacks := [3]string{
		"",
		"",
		"",
	}
	for line := range utils.GetInputLines(INPUT_PATH) {
		rucksackIndex := lineCount % 3
		rucksacks[rucksackIndex] = line
		if rucksackIndex == 2 {
			for _, char := range strings.Split(line, "") {
				if strings.Contains(rucksacks[0], char) &&
					strings.Contains(rucksacks[1], char) {
					r, _ := utf8.DecodeRuneInString(char)
					prioritySum += getPriority(r)
					break
				}
			}
		}
		lineCount++
	}

	return prioritySum
}
