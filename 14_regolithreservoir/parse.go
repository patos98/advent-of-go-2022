package regolithreservoir

import (
	"aoc-2022-go/utils"
	"math"
	"strings"
)

func parseRockPositions(inputPath string) map[string]Position {
	rockPositions := map[string]Position{}

	utils.ProcessInputLines(inputPath, func(line string) {
		positions := strings.Split(line, " -> ")
		count := len(positions)
		for positionIndex := 0; positionIndex < count-1; positionIndex++ {
			position1 := parsePosition(positions[positionIndex])
			position2 := parsePosition(positions[positionIndex+1])

			if position1.X == position2.X {
				difference := position2.Y - position1.Y
				absoluteDiff := math.Abs(float64(difference))
				for i := 0; i <= int(absoluteDiff); i++ {
					var y int
					if math.Signbit(float64(difference)) {
						y = position1.Y + -1*i
					} else {
						y = position1.Y + i
					}
					position := Position{
						X: position1.X,
						Y: y,
					}
					rockPositions[position.ToString()] = position
				}
			} else {
				difference := position2.X - position1.X
				absoluteDiff := math.Abs(float64(difference))
				for i := 0; i <= int(absoluteDiff); i++ {
					var x int
					if math.Signbit(float64(difference)) {
						x = position1.X + -1*i
					} else {
						x = position1.X + i
					}

					position := Position{
						X: x,
						Y: position1.Y,
					}

					rockPositions[position.ToString()] = position
				}
			}
		}
	})

	return rockPositions
}
