package regolithreservoir

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

func (p Position) ToString() string {
	return fmt.Sprintf("%d;%d", p.X, p.Y)
}

func createPosition(x int, y int) Position {
	return Position{
		X: x,
		Y: y,
	}
}

func parsePosition(s string) Position {
	coordinates := strings.Split(s, ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	return Position{
		X: x,
		Y: y,
	}
}
