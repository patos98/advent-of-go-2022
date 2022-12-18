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
	rockPositions := parseRockPositions(INPUT_PATH)
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

	clearConsole()
	printMapToConsole(rockPositions, sandPositions, foreverFallingPositions, lowestLevel, true)

	sandCount := 0
	sands := []Sand{}
	sandPaths := [][]Position{}
	for !exit {

		sands = append(sands, sandFactory())
		sandPaths = append(sandPaths, []Position{})
		solidPositions := mergePosititonMaps(rockPositions, sandPositions)
		for i, sand := range sands {
			previousPosition := sand.GetPosition()
			sandPaths[i] = append(sandPaths[i], previousPosition)
			sand.Fall(solidPositions)
			solidPositions[sand.GetPosition().ToString()] = sand.GetPosition()
			if sand.GetPosition().ToString() == previousPosition.ToString() && !exitCondition(sand) {
				sandPositions[sand.GetPosition().ToString()] = sand.GetPosition()
				sands = sands[1:]
				sandCount++
				printSandToConsole(sand.GetPosition(), lowestLevel)
			}
		}
		if exitCondition(sands[0]) {
			exit = true
			if len(sandPaths[0]) > 2 {
				for _, position := range sandPaths[0] {
					foreverFallingPositions[position.ToString()] = position
				}
			} else {
				sandPositions[sands[0].GetPosition().ToString()] = sands[0].GetPosition()
			}
			break
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
