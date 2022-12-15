package regolithreservoir

import (
	"fmt"
	"io/ioutil"
)

type CharPrinter func(s string)
type NewLinePrinter func()

func printMapToConsole(rockPositions map[string]Position, sandPositions map[string]Position, foreverFallingPositions map[string]Position, maxY int, hasFloor bool) {
	printMap(rockPositions, sandPositions, foreverFallingPositions, maxY, hasFloor, func(s string) { fmt.Print(s) }, func() { fmt.Println() })
}

func printMapToFile(rockPositions map[string]Position, sandPositions map[string]Position, foreverFallingPositions map[string]Position, maxY int, hasFloor bool, filePath string) {
	content := ""
	printMap(
		rockPositions,
		sandPositions,
		foreverFallingPositions,
		maxY,
		hasFloor,
		func(s string) {
			content += s
		},
		func() {
			content += "\n"
		},
	)
	ioutil.WriteFile(filePath, []byte(content), 0644)
}

func printMap(rockPositions map[string]Position, sandPositions map[string]Position, foreverFallingPositions map[string]Position, maxY int, hasFloor bool, charPrinter CharPrinter, newLinePrinter NewLinePrinter) {
	minX := 999
	minY := 999
	maxX := 0
	for _, position := range mergePosititonMaps(rockPositions, sandPositions) {
		if position.X < minX {
			minX = position.X
		}
		if position.Y < minY {
			minY = position.Y
		}
		if position.X > maxX {
			maxX = position.X
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			charPrinter(" ")
			position := Position{
				X: x,
				Y: y,
			}
			if _, contains := rockPositions[position.ToString()]; contains {
				charPrinter("#")
			} else if _, contains := sandPositions[position.ToString()]; contains {
				charPrinter("O")
			} else if _, contains := foreverFallingPositions[position.ToString()]; contains {
				charPrinter("~")
			} else if hasFloor && y == maxY {
				charPrinter("#")
			} else {
				charPrinter(".")
			}
		}
		newLinePrinter()
	}

}
