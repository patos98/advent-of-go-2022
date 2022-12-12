package hillclimbingalgorithm

import (
	"aoc-2022-go/utils"
	"unicode/utf8"
)

func getHeightMap(inputPath string) ([][]int, Path, string) {
	heightMap := [][]int{}
	var startPosition Path
	var endPosition Path
	lineCount := 0
	utils.ProcessInputLines(inputPath, func(line string) {
		heights := []int{}
		count := utf8.RuneCountInString(line)
		for i := 0; i < count; i++ {
			height := line[i]
			if height == 'E' {
				height = 'z'
				endPosition = Path{
					steps:      "",
					currentRow: lineCount,
					currentCol: i,
				}
			} else if height == 'S' {
				height = 'a'
				startPosition = Path{
					steps:      "",
					currentRow: lineCount,
					currentCol: i,
				}
			}
			heights = append(heights, int(height))
		}
		heightMap = append(heightMap, heights)
		lineCount++
	})

	return heightMap, startPosition, endPosition.getPosition()
}
