package hillclimbingalgorithm

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Path struct {
	steps      string
	currentRow int
	currentCol int
}

func (p Path) getPosition() string {
	return fmt.Sprintf("%03d%03d", p.currentRow, p.currentCol)
}

func (p Path) drawStepsToFile(filePath string) {
	row, col, minrow, mincol, maxrow, maxcol := 0, 0, 0, 0, 0, 0
	stepMap := map[int]map[int]string{}
	for _, step := range strings.Split(p.steps, "") {
		if minrow > row {
			minrow = row
		}
		if mincol > col {
			mincol = col
		}
		if maxrow < row {
			maxrow = row
		}
		if maxcol < col {
			maxcol = col
		}

		if _, isPresent := stepMap[row]; !isPresent {
			stepMap[row] = map[int]string{}
		}

		stepMap[row][col] = step
		if step == "<" {
			col--
		} else if step == ">" {
			col++
		} else if step == "^" {
			row--
		} else if step == "v" {
			row++
		}
	}

	steps := ""
	for i := minrow; i <= maxrow; i++ {
		for j := mincol; j <= maxcol; j++ {
			step, isPresent := stepMap[i][j]
			if !isPresent {
				step = " "
			}
			steps += step
		}
		steps += "\n"
	}

	ioutil.WriteFile(filePath, []byte(steps), 0644)

}
