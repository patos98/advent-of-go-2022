package ropebridge

import (
	"strconv"
	"strings"
)

type Step struct {
	direction string
	amount    int
}

func parseStep(s string) Step {
	split := strings.Split(s, " ")
	amount, _ := strconv.Atoi(split[1])
	return Step{
		direction: split[0],
		amount:    amount,
	}
}
