package rockpaperscissors

import (
	"aoc-2022-go/utils"
)

const INPUT_PATH = "02_rock_paper_scissors/input.txt"

type Asset string

const (
	ROCK     = "X"
	PAPER    = "Y"
	SCISSORS = "Z"
)

var signs1 = map[string]Asset{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var signs2 = map[string]Asset{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

var assetPoints = map[Asset]int{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,
}

const (
	WIN  int = 6
	DRAW     = 3
	LOSS     = 0
)

const (
	ROCK_ROCK         = "A X"
	ROCK_PAPER        = "A Y"
	ROCK_SCISSORS     = "A Z"
	PAPER_ROCK        = "B X"
	PAPER_PAPER       = "B Y"
	PAPER_SCISSORS    = "B Z"
	SCISSORS_ROCK     = "C X"
	SCISSORS_PAPER    = "C Y"
	SCISSORS_SCISSORS = "C Z"
)

var pointsByOutComes = map[string]int{
	ROCK_ROCK:         assetPoints[signs2["X"]] + DRAW,
	ROCK_PAPER:        assetPoints[signs2["Y"]] + WIN,
	ROCK_SCISSORS:     assetPoints[signs2["Z"]] + LOSS,
	PAPER_ROCK:        assetPoints[signs2["X"]] + LOSS,
	PAPER_PAPER:       assetPoints[signs2["Y"]] + DRAW,
	PAPER_SCISSORS:    assetPoints[signs2["Z"]] + WIN,
	SCISSORS_ROCK:     assetPoints[signs2["X"]] + WIN,
	SCISSORS_PAPER:    assetPoints[signs2["Y"]] + LOSS,
	SCISSORS_SCISSORS: assetPoints[signs2["Z"]] + DRAW,
}

func CalculateScore() int {
	totalPoints := 0

	for line := range utils.GetInputLines(INPUT_PATH) {
		totalPoints += pointsByOutComes[line]
	}

	return totalPoints
}

const (
	ROCK_LOSS     = "A X"
	ROCK_DRAW     = "A Y"
	ROCK_WIN      = "A Z"
	PAPER_LOSS    = "B X"
	PAPER_DRAW    = "B Y"
	PAPER_WIN     = "B Z"
	SCISSORS_LOSS = "C X"
	SCISSORS_DRAW = "C Y"
	SCISSORS_WIN  = "C Z"
)

var expectedAsset = map[string]string{
	ROCK_LOSS:     SCISSORS,
	ROCK_DRAW:     ROCK,
	ROCK_WIN:      PAPER,
	PAPER_LOSS:    ROCK,
	PAPER_DRAW:    PAPER,
	PAPER_WIN:     SCISSORS,
	SCISSORS_LOSS: PAPER,
	SCISSORS_DRAW: SCISSORS,
	SCISSORS_WIN:  ROCK,
}

func CalculateScore2() int {
	totalPoints := 0

	for line := range utils.GetInputLines(INPUT_PATH) {
		line = line[0:2] + expectedAsset[line]
		totalPoints += pointsByOutComes[line]
	}

	return totalPoints
}
