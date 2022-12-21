package monkey

import (
	"strconv"
	"strings"
)

func Parse(s string) Monkey {
	split := strings.Split(s, ": ")
	name := split[0]

	if strings.Contains(split[1], " ") {
		return parseMonkeyWithOperation(name, split[1])
	} else {
		return parseMonkeyWithNumber(name, split[1])
	}
}

var operators = map[string]Operator{
	"+": func(i1, i2 int) int { return i1 + i2 },
	"-": func(i1, i2 int) int { return i1 - i2 },
	"*": func(i1, i2 int) int { return i1 * i2 },
	"/": func(i1, i2 int) int { return i1 / i2 },
}

func parseMonkeyWithOperation(name string, s string) Monkey {
	split := strings.Split(s, " ")
	return MonkeyWithOperation{
		name:        name,
		monkeyLeft:  split[0],
		monkeyRight: split[2],
		operator:    operators[split[1]],
	}
}

func parseMonkeyWithNumber(name string, s string) Monkey {
	number, _ := strconv.Atoi(s)
	return MonkeyWithNumber{
		name:   name,
		number: number,
	}
}
