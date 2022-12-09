package treehouse

import (
	"aoc-2022-go/utils"
	"fmt"
	"sort"
	"strconv"
)

const INPUT_PATH = "08_treehouse/input.txt"

func GetVisibleTreesFromOutside() int {
	grid, height, width := getGrid()

	top := []Visibility{}
	bottom := []Visibility{}
	left := []Visibility{}
	right := []Visibility{}

	for col := 0; col < width; col++ {
		top = append(top, newVisibility())
		bottom = append(bottom, newVisibility())
	}

	visibleIndexes := map[string]struct{}{}
	for row := 0; row < height; row++ {
		left = append(left, newVisibility())
		right = append(right, newVisibility())
		for col := 0; col < width; col++ {
			topLeftNext := grid[row][col]
			bottomRightNext := grid[height-row-1][width-col-1]

			visibleFromTop := top[col].next(topLeftNext)
			visibleFromLeft := left[row].next(topLeftNext)
			if visibleFromTop || visibleFromLeft {
				index := fmt.Sprintf("%02d", row) + fmt.Sprintf("%02d", col)
				visibleIndexes[index] = struct{}{}
			}

			visibleFromBottom := bottom[col].next(bottomRightNext)
			visibleFromRight := right[row].next(bottomRightNext)
			if visibleFromBottom || visibleFromRight {
				index := fmt.Sprintf("%02d", height-row-1) + fmt.Sprintf("%02d", width-col-1)
				visibleIndexes[index] = struct{}{}
			}
		}
	}
	visibleIndexSlice := []string{}
	for key := range visibleIndexes {
		visibleIndexSlice = append(visibleIndexSlice, key)
	}
	sort.Strings(visibleIndexSlice)

	return len(visibleIndexes)
}

func GetTreeWithMaxScenicScore() int {
	grid, height, width := getGrid()
	maxScenicScore := 0
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			scenicScore := getScenicScore(grid, height, width, row, col)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	return maxScenicScore
}

func getGrid() ([][]int, int, int) {
	lines := utils.GetAllInputLines(INPUT_PATH)
	height := len(lines)
	width := len(lines[0])

	grid := [][]int{}
	for row := 0; row < height; row++ {
		gridRow := []int{}
		for col := 0; col < width; col++ {
			treeHeight, _ := strconv.Atoi(string(lines[row][col]))
			gridRow = append(gridRow, treeHeight)
		}
		grid = append(grid, gridRow)
	}

	return grid, height, width
}
