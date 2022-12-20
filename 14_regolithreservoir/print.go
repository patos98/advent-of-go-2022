package regolithreservoir

import (
	"fmt"
	"io/ioutil"

	tm "github.com/buger/goterm"
)

type CharPrinter func(p Position, s string)
type NewLinePrinter func()

func clearConsole() {
	tm.Clear()
	tm.Flush()
}

func printSandToConsole(sandPosition Position, lowestLevel int) {
	x, y := tm.GetXY(sandPosition.X, sandPosition.Y)
	tm.MoveCursor(x/2, y)
	tm.Flush()
	fmt.Print("O")
}

func printCharToConsole(p Position, character string) {
	x, y := tm.GetXY(p.X, p.Y)
	tm.MoveCursor(x/2, y)
	tm.Flush()
	fmt.Print(character)
}

func printMapToConsole(rockPositions map[string]Position, sandPositions map[string]Position, foreverFallingPositions map[string]Position, maxY int, hasFloor bool) {
	printMap(
		rockPositions,
		sandPositions,
		foreverFallingPositions,
		maxY,
		hasFloor,
		printCharToConsole,
		func() { fmt.Println() },
	)
}

func printMapToFile(rockPositions map[string]Position, sandPositions map[string]Position, foreverFallingPositions map[string]Position, maxY int, hasFloor bool, filePath string) {
	content := ""
	printMap(
		rockPositions,
		sandPositions,
		foreverFallingPositions,
		maxY,
		hasFloor,
		func(p Position, s string) {
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
			position := Position{
				X: x,
				Y: y,
			}
			charPrinter(position, " ")
			if _, contains := rockPositions[position.ToString()]; contains {
				charPrinter(position, "#")
			} else if _, contains := sandPositions[position.ToString()]; contains {
				charPrinter(position, "O")
			} else if _, contains := foreverFallingPositions[position.ToString()]; contains {
				charPrinter(position, "~")
			} else if hasFloor && y == maxY {
				charPrinter(position, "#")
			} else {
				charPrinter(position, " ")
			}
		}
		newLinePrinter()
	}

}
