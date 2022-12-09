package treehouse

import (
	"aoc-2022-go/utils"
	"fmt"
	"sort"
	"strconv"
)

const INPUT_PATH = "08_treehouse/input.txt"

type Visibility struct {
	count     int
	maxHeight int
}

func (v *Visibility) next(nextHeight int) bool {
	if nextHeight > v.maxHeight {
		v.maxHeight = nextHeight
		return true
	}
	return false
}

func newVisibility() Visibility {
	return Visibility{
		count:     0,
		maxHeight: -1,
	}
}

func GetVisibleTreesFromOutside() int {
	grid := utils.GetAllInputLines(INPUT_PATH)
	height := len(grid)
	width := len(grid[0])

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
			topLeftNext, _ := strconv.Atoi(string(grid[row][col]))
			bottomRightNext, _ := strconv.Atoi(string(grid[height-row-1][width-col-1]))

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
