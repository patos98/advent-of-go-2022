package hillclimbingalgorithm

import "fmt"

const INPUT_PATH = "12_hillclimbingalgorithm/input.txt"
const TEST_INPUT_PATH = "12_hillclimbingalgorithm/input_test.txt"
const RESULT_DIRECTORY = "12_hillclimbingalgorithm/results"

func GetShortestPathLengthFromS() int {
	heightMap, startPosition, endPosition := getHeightMap(INPUT_PATH)
	return hillClimbingAlgorithm(heightMap, startPosition, endPosition, RESULT_DIRECTORY+"/result_part1.txt")
}

func GetShortestPathFromAnyA() int {
	heightMap, _, endPosition := getHeightMap(INPUT_PATH)
	startPositions := []Path{}
	rowCount := len(heightMap)
	colCount := len(heightMap[0])
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if heightMap[i][j] == int('a') {
				startPositions = append(startPositions, Path{
					steps:      "",
					currentRow: i,
					currentCol: j,
				})
			}
		}
	}

	paths := [][]Path{}
	visitedPositions := []map[string]struct{}{}
	for _, startPosition := range startPositions {
		paths = append(paths, []Path{startPosition})
		visitedPositions = append(visitedPositions, map[string]struct{}{startPosition.getPosition(): {}})
	}

	startPositionCount := len(startPositions)
	for {
		for i := 0; i < startPositionCount; i++ {
			newPaths, stepCount := getNewPaths(paths[i], heightMap, rowCount, colCount, visitedPositions[i], endPosition, RESULT_DIRECTORY+fmt.Sprintf("/result_part2_%d.txt", i))
			if stepCount > 0 {
				return stepCount
			}
			paths[i] = newPaths
		}
	}
}

func hillClimbingAlgorithm(heightMap [][]int, startPosition Path, endPosition string, resultPath string) int {
	paths := []Path{startPosition}
	visitedPositions := map[string]struct{}{
		startPosition.getPosition(): {},
	}

	rowCount := len(heightMap)
	colCount := len(heightMap[0])
	for {
		newPaths, stepCount := getNewPaths(paths, heightMap, rowCount, colCount, visitedPositions, endPosition, resultPath)
		if stepCount > 0 {
			return stepCount
		}
		paths = newPaths
	}
}

func getNewPaths(paths []Path, heightMap [][]int, rowCount int, colCount int, visitedPositions map[string]struct{}, endPosition string, resultPath string) ([]Path, int) {
	newPaths := []Path{}
	for _, path := range paths {
		currentHeight := heightMap[path.currentRow][path.currentCol]
		newPaths = nextStep(path.currentRow, path.currentCol-1, rowCount, colCount, visitedPositions, path.steps+"<", currentHeight, heightMap, newPaths)
		if len(newPaths) > 0 && newPaths[len(newPaths)-1].getPosition() == endPosition {
			newPaths[len(newPaths)-1].drawStepsToFile(resultPath)
			return paths, len(newPaths[len(newPaths)-1].steps)
		}
		newPaths = nextStep(path.currentRow, path.currentCol+1, rowCount, colCount, visitedPositions, path.steps+">", currentHeight, heightMap, newPaths)
		if len(newPaths) > 0 && newPaths[len(newPaths)-1].getPosition() == endPosition {
			newPaths[len(newPaths)-1].drawStepsToFile(resultPath)
			return paths, len(newPaths[len(newPaths)-1].steps)
		}
		newPaths = nextStep(path.currentRow-1, path.currentCol, rowCount, colCount, visitedPositions, path.steps+"^", currentHeight, heightMap, newPaths)
		if len(newPaths) > 0 && newPaths[len(newPaths)-1].getPosition() == endPosition {
			newPaths[len(newPaths)-1].drawStepsToFile(resultPath)
			return paths, len(newPaths[len(newPaths)-1].steps)
		}
		newPaths = nextStep(path.currentRow+1, path.currentCol, rowCount, colCount, visitedPositions, path.steps+"v", currentHeight, heightMap, newPaths)
		if len(newPaths) > 0 && newPaths[len(newPaths)-1].getPosition() == endPosition {
			newPaths[len(newPaths)-1].drawStepsToFile(resultPath)
			return paths, len(newPaths[len(newPaths)-1].steps)
		}
	}

	return newPaths, 0
}

func nextStep(nextRow int, nextCol int, rowCount int, colCount int, visitedPositions map[string]struct{}, steps string, currentHeight int, heightMap [][]int, newPaths []Path) []Path {
	if nextRow < 0 || nextRow >= rowCount || nextCol < 0 || nextCol >= colCount {
		return newPaths
	}

	newPath := Path{
		steps:      steps,
		currentRow: nextRow,
		currentCol: nextCol,
	}
	newPathPosition := newPath.getPosition()

	if _, contains := visitedPositions[newPathPosition]; contains {
		return newPaths
	}

	newHeight := heightMap[nextRow][nextCol]
	if newHeight-currentHeight > 1 {
		return newPaths
	}

	visitedPositions[newPathPosition] = struct{}{}
	return append(newPaths, newPath)
}
