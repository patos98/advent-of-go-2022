package regolithreservoir

type Sand interface {
	GetPosition() Position
	Fall(solidPositions map[string]Position)
}

type SimpleSand struct {
	position Position
}

func createSimpleSand() Sand {
	return &SimpleSand{
		position: Position{
			X: 500,
			Y: 0,
		},
	}
}

func (s *SimpleSand) GetPosition() Position {
	return s.position
}

func (s *SimpleSand) Fall(solidPositions map[string]Position) {
	nextPosition := Position{
		X: s.position.X,
		Y: s.position.Y + 1,
	}

	if _, contains := solidPositions[nextPosition.ToString()]; !contains {
		s.position = nextPosition
		return
	}

	nextPosition = Position{
		X: nextPosition.X - 1,
		Y: nextPosition.Y,
	}

	if _, contains := solidPositions[nextPosition.ToString()]; !contains {
		s.position = nextPosition
		return
	}

	nextPosition = Position{
		X: nextPosition.X + 2,
		Y: nextPosition.Y,
	}

	if _, contains := solidPositions[nextPosition.ToString()]; !contains {
		s.position = nextPosition
		return
	}
}

type FloorAwareSand struct {
	sand       Sand
	floorLevel int
}

func createFloorAwareSand(floorLevel int) Sand {
	return &FloorAwareSand{
		sand:       createSimpleSand(),
		floorLevel: floorLevel,
	}
}

func (s *FloorAwareSand) GetPosition() Position {
	return s.sand.GetPosition()
}

func (s *FloorAwareSand) Fall(solidPositions map[string]Position) {
	currentPosition := s.sand.GetPosition()
	for x := currentPosition.X - 1; x <= currentPosition.X+1; x++ {
		floorPosition := Position{
			X: x,
			Y: s.floorLevel,
		}
		solidPositions[floorPosition.ToString()] = floorPosition
	}

	s.sand.Fall(solidPositions)
}
