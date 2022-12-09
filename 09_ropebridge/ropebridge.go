package ropebridge

import (
	"aoc-2022-go/utils"
)

const INPUT_PATH = "09_ropebridge/input.txt"
const TEST_INPUT_PATH = "09_ropebridge/input_test.txt"

type PositionUpdateRule func(Position, int) Position
type Rules map[string]PositionUpdateRule

func SimulateSimpleRope() int {
	return simulateRope(newSimpleRope())
}

func SimulateSuperRope() int {
	return simulateRope(newSuperRope(9))
}

func simulateRope(rope Rope) int {
	var rules = Rules{
		"U": func(p Position, amount int) Position { return Position{x: p.x, y: p.y + amount} },
		"D": func(p Position, amount int) Position { return Position{x: p.x, y: p.y - amount} },
		"L": func(p Position, amount int) Position { return Position{x: p.x - amount, y: p.y} },
		"R": func(p Position, amount int) Position { return Position{x: p.x + amount, y: p.y} },
	}

	visitedPositions := map[string]struct{}{}
	utils.ProcessInputLines(INPUT_PATH, func(line string) {
		step := parseStep(line)
		for i := 0; i < step.amount; i++ {
			rope.moveHead(rules, step)
			rope.moveTail()
			visitedPositions[rope.getTail()] = struct{}{}
		}
	})

	return len(visitedPositions)
}
