package regolithreservoir

const INPUT_PATH = "14_regolithreservoir/input.txt"
const TEST_INPUT_PATH = "14_regolithreservoir/input_test.txt"
const RESULT_FILE_PATH_1 = "14_regolithreservoir/result_part1.txt"
const RESULT_FILE_PATH_2 = "14_regolithreservoir/result_part2.txt"

func GetNumberOfRestingSandUnits() int {
	rockPositions := parseRockPositions(INPUT_PATH)
	lowestLevel, _ := getLowestAndRightMostPoint(rockPositions)
	sandCount, sandPositions, foreverFallingPositions := getSandCount(rockPositions, lowestLevel, createSimpleSand, func(sand Sand) bool { return sand.GetPosition().Y > lowestLevel })
	printMapToFile(rockPositions, sandPositions, foreverFallingPositions, lowestLevel, false, RESULT_FILE_PATH_1)
	return sandCount
}

func GetNumberOfRestingSandUnitsWithFloor() int {
	rockPositions := parseRockPositions(TEST_INPUT_PATH)
	lowestLevel, _ := getLowestAndRightMostPoint(rockPositions)
	lowestLevel += 2
	sandFactory := func() Sand { return createFloorAwareSand(lowestLevel) }
	sandCount, sandPositions, foreverFallingPositions := getSandCount(rockPositions, lowestLevel, sandFactory, func(sand Sand) bool { return sand.GetPosition().Y == 0 })
	printMapToFile(rockPositions, sandPositions, foreverFallingPositions, lowestLevel, true, RESULT_FILE_PATH_2)
	return sandCount + 1 // the last sand stays on top
}

func getSandCount(rockPositions map[string]Position, lowestLevel int, sandFactory func() Sand, exitCondition func(sand Sand) bool) (int, map[string]Position, map[string]Position) {
	sandPositions := map[string]Position{}
	exit := false
	foreverFallingPositions := map[string]Position{}

	sandCount := -1
	for !exit {
		printMapToFile(rockPositions, sandPositions, foreverFallingPositions, lowestLevel, true, RESULT_FILE_PATH_2)

		sandCount++
		sand := sandFactory()
		sandPath := []Position{}
		for {
			sandPath = append(sandPath, sand.GetPosition())
			previousPosition := sand.GetPosition()
			sand.Fall(mergePosititonMaps(rockPositions, sandPositions))
			if sand.GetPosition().ToString() == previousPosition.ToString() && !exitCondition(sand) {
				sandPositions[sand.GetPosition().ToString()] = sand.GetPosition()
				break
			}
			if exitCondition(sand) {
				exit = true
				if len(sandPath) > 1 {
					for _, position := range sandPath {
						foreverFallingPositions[position.ToString()] = position
					}
				} else {
					sandPositions[sand.GetPosition().ToString()] = sand.GetPosition()
				}
				break
			}
		}
	}

	return sandCount, sandPositions, foreverFallingPositions
}

func getLowestAndRightMostPoint(rockPositions map[string]Position) (int, int) {
	lowest := 0
	rightMost := 0
	for _, position := range rockPositions {
		if position.Y > lowest {
			lowest = position.Y
		}
		if position.X > rightMost {
			rightMost = position.X
		}
	}

	return lowest, rightMost
}

func mergePosititonMaps(rockPositions map[string]Position, sandPositions map[string]Position) map[string]Position {
	mergedPositions := map[string]Position{}
	for key, value := range rockPositions {
		mergedPositions[key] = value
	}
	for key, value := range sandPositions {
		mergedPositions[key] = value
	}

	return mergedPositions
}
