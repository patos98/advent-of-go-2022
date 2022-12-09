package treehouse

type ScenicScore struct {
	score int
	final bool
}

func newScenicScore() *ScenicScore {
	return &ScenicScore{
		score: 0,
		final: false,
	}
}

func getScenicScore(grid [][]int, height int, width int, row int, col int) int {
	currentTree := grid[row][col]
	up := newScenicScore()
	down := newScenicScore()
	left := newScenicScore()
	right := newScenicScore()

	distance := 1
	for {
		updateScenicScore(up, currentTree, grid, row-distance, col)
		updateScenicScore(down, currentTree, grid, row+distance, col)
		updateScenicScore(left, currentTree, grid, row, col-distance)
		updateScenicScore(right, currentTree, grid, row, col+distance)

		if up.final && down.final && left.final && right.final {
			break
		}

		distance++
	}

	return up.score * down.score * left.score * right.score
}

func updateScenicScore(scenicScore *ScenicScore, currentTree int, grid [][]int, nextRow int, nextCol int) {
	if nextRow < 0 || nextRow > 98 || nextCol < 0 || nextCol > 98 {
		scenicScore.final = true
		return
	}

	if scenicScore.final {
		return
	}

	scenicScore.score++
	nextTree := grid[nextRow][nextCol]
	if currentTree <= nextTree {
		scenicScore.final = true
	}
}
