package ropebridge

import (
	"fmt"
	"math"
)

type Position struct {
	x int
	y int
}

func startingPosition() Position {
	return Position{
		x: 0,
		y: 0,
	}
}

func (p Position) getDifference(o Position) Position {
	return Position{
		x: p.x - o.x,
		y: p.y - o.y,
	}
}

func (p Position) trim(limit int) Position {
	return Position{
		x: trimCoordinate(p.x, limit),
		y: trimCoordinate(p.y, limit),
	}
}

func trimCoordinate(coordinate int, limit int) int {
	if coordinate > 0 {
		return int(math.Min(float64(coordinate), float64(limit)))
	} else {
		return int(math.Max(float64(coordinate), float64(-limit)))
	}
}

func (p Position) toString() string {
	return fmt.Sprintf("%30d%30d", p.x, p.y)
}
