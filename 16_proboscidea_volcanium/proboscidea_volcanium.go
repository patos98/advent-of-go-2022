package proboscideavolcanium

import (
	"aoc-2022-go/16_proboscidea_volcanium/valve"
	"aoc-2022-go/utils"
)

type Valve interface {
	GetName() string
	GetFlowRate() int
}

func getMostReleasablePressure(inputPath string) int {
	valves := parseValves(inputPath)
	valveCount := len(valves)

	valveMap := map[string]map[string]int{}
	for _, targetValve := range valves {
		for _, sourceValve := range valves {
			// do what?
		}
	}

	return 0
}

func parseValves(inputPath string) map[string]Valve {
	valves := map[string]Valve{}
	utils.ProcessInputLines(inputPath, func(s string) {
		valve := valve.Parse(inputPath)
		valves[valve.GetName()] = valve
	})
	return valves
}
