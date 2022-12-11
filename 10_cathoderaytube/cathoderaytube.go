package cathoderaytube

import (
	"aoc-2022-go/utils"
	"fmt"
	"strconv"
	"strings"
)

const INPUT_PATH = "10_cathoderaytube/input.txt"

type Command struct {
	timeToLive int
	execute    func()
}

type CommandFactory func([]string) Command

func GetSignalStrengths() int {
	sum := 0
	cyclesToInspect := map[int]struct{}{
		20:  {},
		60:  {},
		100: {},
		140: {},
		180: {},
		220: {},
	}
	inspectCycles(func(registerX int, cycle int) {
		if _, contains := cyclesToInspect[cycle]; contains {
			signalStrength := registerX * cycle
			sum += signalStrength
		}
	})

	return sum
}

func DrawSprite() {
	spriteRadius := 1
	screenWidth := 40
	screen := []string{}
	inspectCycles(func(registerX int, cycle int) {
		rowIndex := (cycle - 1) / screenWidth
		if len(screen) == rowIndex {
			screen = append(screen, "")
		}

		cyclePosition := (cycle - 1) % screenWidth
		var currentChar string
		if registerX >= cyclePosition-spriteRadius && registerX <= cyclePosition+spriteRadius {
			currentChar = "#"
		} else {
			currentChar = " "
		}

		screen[rowIndex] += currentChar
	})

	for _, row := range screen {
		fmt.Println(row)
	}
}

func inspectCycles(cycleInspector func(int, int)) {
	registerX := 1
	cycle := 1

	commands := map[string]CommandFactory{
		"noop": func(split []string) Command { return Command{timeToLive: 1, execute: func() {}} },
		"addx": func(split []string) Command {
			return Command{
				timeToLive: 2,
				execute: func() {
					value, _ := strconv.Atoi(split[1])
					registerX += value
				},
			}
		},
	}

	utils.ProcessInputLines(INPUT_PATH, func(line string) {
		split := strings.Split(line, " ")
		command := commands[split[0]](split)
		for {
			if command.timeToLive == 0 {
				break
			}

			cycleInspector(registerX, cycle)

			command.timeToLive--
			if command.timeToLive == 0 {
				command.execute()
			}
			cycle++
		}
	})
}
