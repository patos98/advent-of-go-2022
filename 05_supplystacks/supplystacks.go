package supplystacks

import (
	"aoc-2022-go/utils"
	"strings"
)

const INPUT_PATH = "05_supplystacks/input.txt"

func Rearrange() {
	stacksDrawingSection := true
	stacksDrawing := []string{}
	stacks := []CrateStack{}
	for line := range utils.GetInputLines(INPUT_PATH) {
		if line == "" {
			stacks = createStacksFromDrawing(stacksDrawing)
			stacksDrawingSection = false
			continue
		}

		if !stacksDrawingSection {
			stacksDrawing = append(stacksDrawing, line)
		} else {
			executeRearrangement(stacks, line)
		}
	}

}

func createStacksFromDrawing(drawing []string) []CrateStack {
	stacks := [9]CrateStack{
		newCrateStack(),
		newCrateStack(),
		newCrateStack(),
		newCrateStack(),
		newCrateStack(),
		newCrateStack(),
		newCrateStack(),
		newCrateStack(),
		newCrateStack(),
	}
	count := len(drawing)
	for i := count - 2; i >= 0; i-- {
		line := drawing[i]
		for j := 0; j < 9; j++ {
			crateLetterIndex := 1 + j*4
			crateLetter := strings.Split(line, "")[crateLetterIndex]
			if crateLetter != " " {
				stacks[j].push(crateLetter)
			}
		}
	}
	return stacks[:]
}

func executeRearrangement(stacks []CrateStack, rearrangementString string) {
	splitted := strings.Split(strings.Split(rearrangementString, "move ")[1], " from ")
}
