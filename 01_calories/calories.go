package calories

import (
	"aoc-2022-go/utils"
	"strconv"
)

var INPUT_PATH = "01_calories/input.txt"

func GetTop3Calories() int {
	thirdMostCalories := 0
	top3Calories := [3]int{0, 0, 0}
	currentElfCalories := 0
	for line := range utils.GetInputLines(INPUT_PATH) {
		if line == "" {
			for i, maxCalories := range top3Calories {
				if maxCalories == thirdMostCalories && currentElfCalories > maxCalories {
					top3Calories[i] = currentElfCalories
					break
				}
			}
			leastCalories := top3Calories[0]
			for _, maxCalories := range top3Calories {
				if maxCalories < leastCalories {
					leastCalories = maxCalories
				}
			}
			thirdMostCalories = leastCalories
			currentElfCalories = 0
			continue
		}

		calories, _ := strconv.Atoi(line)
		currentElfCalories += calories
	}

	maxCalories := 0
	for _, calories := range top3Calories {
		maxCalories += calories
	}

	return maxCalories
}

func GetMaxCalories() int {
	maxCalories := 0
	currentElfCalories := 0
	for line := range utils.GetInputLines(INPUT_PATH) {
		if line == "" {
			if currentElfCalories > maxCalories {
				maxCalories = currentElfCalories
			}
			currentElfCalories = 0
			continue
		}

		calories, _ := strconv.Atoi(line)
		currentElfCalories += calories
	}

	return maxCalories
}
