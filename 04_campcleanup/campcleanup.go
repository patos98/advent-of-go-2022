package campcleanup

import (
	"aoc-2022-go/utils"
	"strings"
)

const INPUT_PATH = "04_campcleanup/input.txt"

func CampCleanupFullyContains() int {
	return getOverlappingRangesCount(rangesFullyContain)
}

func CampCleanupOverlap() int {
	return getOverlappingRangesCount(rangesOverlap)
}

func getOverlappingRangesCount(overlapStrategy func(Range, Range) bool) int {
	fullyContainingPairs := 0
	for line := range utils.GetInputLines(INPUT_PATH) {
		rangeStrings := strings.Split(line, ",")
		range1 := createRangeFromString(rangeStrings[0])
		range2 := createRangeFromString(rangeStrings[1])
		if overlapStrategy(range1, range2) {
			fullyContainingPairs++
		}
	}
	return fullyContainingPairs
}
