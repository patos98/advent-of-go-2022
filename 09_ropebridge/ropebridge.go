package ropebridge

import (
	"aoc-2022-go/utils"
)

const INPUT_PATH = "09_ropebridge/input.txt"

type PositionUpdateRule func(Position, int) Position
type Rules map[string]PositionUpdateRule

func RopeBridge() {
	var rules = Rules{
		"U": func(p Position, amount int) Position { return Position{x: p.x, y: p.y - amount} },
		"D": func(p Position, amount int) Position { return Position{x: p.x, y: p.y + amount} },
		"L": func(p Position, amount int) Position { return Position{x: p.x - amount, y: p.y} },
		"R": func(p Position, amount int) Position { return Position{x: p.x + amount, y: p.y} },
	}

	rope := newRope()
	utils.ProcessInputLines(INPUT_PATH, func(line string) {
		executeStep(rules, rope, parseStep(line))
	})
}

func executeStep(rules Rules, rope *Rope, step Step) {
	for i := 0; i < step.amount; i++ {

	}
}
