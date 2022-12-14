package regolithreservoir

type Sand struct {
	position Position
}

func (s *Sand) fall(rockPositions map[string]struct{}) {
	nextPosition := Position{
		x: s.position.x,
		y: s.position.y + 1,
	}

	if _, contains := rockPositions[nextPosition.toString()]; !contains {
		s.position = nextPosition
		return
	}

	nextPosition = Position{
		x: nextPosition.x - 1,
		y: nextPosition.y,
	}

	if _, contains := rockPositions[nextPosition.toString()]; !contains {
		s.position = nextPosition
		return
	}

	nextPosition = Position{
		x: nextPosition.x + 2,
		y: nextPosition.y,
	}

	if _, contains := rockPositions[nextPosition.toString()]; !contains {
		s.position = nextPosition
		return
	}
}
