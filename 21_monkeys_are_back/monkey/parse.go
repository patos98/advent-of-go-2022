package monkey

import (
	"aoc-2022-go/utils"
	"strconv"
	"strings"
)

func ParseMonkeys(inputPath string) map[string]Monkey {
	monkeys := map[string]Monkey{}
	utils.ProcessInputLines(inputPath, func(s string) {
		m := parseSingleMonkey(s)
		monkeys[m.GetName()] = m
	})
	return monkeys
}

func parseSingleMonkey(s string) Monkey {
	split := strings.Split(s, ": ")
	name := split[0]

	if strings.Contains(split[1], " ") {
		return parseMonkeyWithOperation(name, split[1])
	} else {
		return parseMonkeyWithNumber(name, split[1])
	}
}

func parseMonkeyWithOperation(name string, s string) Monkey {
	split := strings.Split(s, " ")
	operatorString := split[1]
	return MonkeyWithOperation{
		name:           name,
		monkeyLeft:     split[0],
		monkeyRight:    split[2],
		operatorString: operatorString,
		operator:       operators[operatorString],
		antiOperation:  antiOperations[operatorString],
	}
}

func parseMonkeyWithNumber(name string, s string) Monkey {
	number, _ := strconv.Atoi(s)
	return MonkeyWithNumber{
		name:   name,
		number: int(number),
	}
}
