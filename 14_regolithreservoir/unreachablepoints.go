package regolithreservoir

/*

	Unreachable if these points also unreachable:
		- above left
		- above
		- above right

	#####
	#xxx.
	#xx..


	Pyramid area

	h - height
	A - area

	h = 1
	A = 1

	  O

	h = 2
	A = 4 (1 + 3)

	  O
	 OOO

	h = 3
	A = 9 (1 + 3 + 5)

      O
	 OOO
	OOOOO

*/

func UnreachablePositions() int {
	unreachablePositions := parseRockPositions(INPUT_PATH)
	lowestPoint, rightMostPoint := getLowestAndRightMostPoint(unreachablePositions)

	for y := 1; y < lowestPoint+2; y++ {
		for x := 0; x < rightMostPoint; x++ {
			aboveLeft := createPosition(x-1, y-1)
			above := createPosition(x, y-1)
			aboveRight := createPosition(x+1, y-1)
			if contains(unreachablePositions, aboveLeft.ToString()) &&
				contains(unreachablePositions, above.ToString()) &&
				contains(unreachablePositions, aboveRight.ToString()) {
				currentPosition := Position{
					X: x,
					Y: y,
				}
				unreachablePositions[currentPosition.ToString()] = currentPosition
			}
		}
	}

	pyramidArea := calculatePyramidArea(lowestPoint + 2)

	// ASSUMING THAT ALL UNREACHABLE POINTS FALL IN THE PYRAMID
	sandCount := pyramidArea - len(unreachablePositions)

	return sandCount
}

func calculatePyramidArea(height int) int {
	area := 0
	currentRow := 1
	for i := 0; i < height; i++ {
		area += currentRow
		currentRow += 2
	}

	return area
}

func contains(theMap map[string]Position, key string) bool {
	_, contains := theMap[key]
	return contains
}
