package regolithreservoir

import (
	"aoc-2022-go/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const INPUT_PATH = "14_regolithreservoir/input.txt"
const TEST_INPUT_PATH = "14_regolithreservoir/input_test.txt"

type Position struct {
	x int
	y int
}

func (p Position) toString() string {
	return fmt.Sprintf("%03d%03d", p.x, p.y)
}

func parsePosition(s string) Position {
	coordinates := strings.Split(s, ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	return Position{
		x: x,
		y: y,
	}
}

type Rock struct {
	position Position
}

func GetNumberOfRestingSandUnits() int {
	rocks := []Rock{}
	rockPositions := map[string]struct{}{}

	utils.ProcessInputLines(TEST_INPUT_PATH, func(line string) {
		positions := strings.Split(line, " -> ")
		count := len(positions)
		for positionIndex := 0; positionIndex < count-1; positionIndex++ {
			position1 := parsePosition(positions[positionIndex])
			position2 := parsePosition(positions[positionIndex+1])

			if position1.x == position2.x {
				difference := position2.y - position1.y
				for i := 0; i <= difference; i++ {
					var y int
					if math.Signbit(float64(difference)) {
						y = position1.y + -1*i
					} else {
						y = position1.y + i
					}

					rock := Rock{
						position: Position{
							x: position1.x,
							y: y,
						},
					}

					rocks = append(rocks, rock)
					rockPositions[rock.position.toString()] = struct{}{}
				}
			} else {
				difference := position2.x - position1.x
				for i := 0; i <= difference; i++ {
					var x int
					if math.Signbit(float64(difference)) {
						x = position1.x + -1*i
					} else {
						x = position1.x + i
					}

					rock := Rock{
						position: Position{
							x: x,
							y: position1.y,
						},
					}

					rocks = append(rocks, rock)
					rockPositions[rock.position.toString()] = struct{}{}
				}
			}
		}
	})

	return 0
}
