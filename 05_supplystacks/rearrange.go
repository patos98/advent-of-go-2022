package supplystacks

import (
	"strconv"
	"strings"
)

type Rearrangement struct {
	amount int
	from   int
	to     int
}

func executeRearrangementOneByOne(stacks []CrateStack, rearrangementString string) {
	rearrangement := parseRearrangement(rearrangementString)
	for i := 0; i < rearrangement.amount; i++ {
		stacks[rearrangement.to].push(stacks[rearrangement.from].pop())
	}
}

func executeRearrangementPreservingOrder(stacks []CrateStack, rearrangementString string) {
	rearrangement := parseRearrangement(rearrangementString)
	temp := newCrateStack()
	for i := 0; i < rearrangement.amount; i++ {
		temp.push(stacks[rearrangement.from].pop())
	}
	for i := 0; i < rearrangement.amount; i++ {
		stacks[rearrangement.to].push(temp.pop())
	}
}

func parseRearrangement(rearrangementString string) Rearrangement {
	splitted := strings.Split(strings.Split(rearrangementString, "move ")[1], " from ")
	amount, _ := strconv.Atoi(splitted[0])
	splitted = strings.Split(splitted[1], " to ")
	from, _ := strconv.Atoi(splitted[0])
	to, _ := strconv.Atoi(splitted[1])
	return Rearrangement{
		amount: amount,
		from:   from - 1,
		to:     to - 1,
	}
}
