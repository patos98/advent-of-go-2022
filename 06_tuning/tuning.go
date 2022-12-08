package tuning

import (
	"aoc-2022-go/utils"
	"strings"
)

const INPUT_PATH = "06_tuning/input.txt"

func Marker() int {
	return marker(4)
}

func Message() int {
	return marker(14)
}

func marker(uniqueCount int) int {
	line := utils.GetFirstInputLine(INPUT_PATH)
	chars := strings.Split(line, "")
	charCount := len(chars)
	for i := 0; i < charCount; i++ {
		uniqueChars := map[string]struct{}{}
		for j := 0; j < uniqueCount; j++ {
			uniqueChars[chars[i+j]] = struct{}{}
		}
		if len(uniqueChars) == uniqueCount {
			return i + uniqueCount
		}
	}

	return 0
}
