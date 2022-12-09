package ropebridge

type Position struct {
	x int
	y int
}

type Rope struct {
	head Position
	tail Position
}

func startingPosition() Position {
	return Position{
		x: 0,
		y: 0,
	}
}

func newRope() *Rope {
	return &Rope{
		head: startingPosition(),
		tail: startingPosition(),
	}
}
