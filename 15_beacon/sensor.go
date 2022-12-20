package beacon

import "fmt"

type Position struct {
	x int
	y int
}

func (p Position) ToString() string {
	return fmt.Sprintf("%d;%d", p.x, p.y)
}

type Sensor struct {
	self          Position
	closestBeacon Position
	distance      int
}
