package proboscideavolcanium

import (
	"aoc-2022-go/16_proboscidea_volcanium/valve"
	"aoc-2022-go/utils"
)

const INPUT_PATH = "16_proboscidea_volcanium/input.txt"
const TEST_INPUT_PATH = "16_proboscidea_volcanium/input_test.txt"

const MINUTES_TO_OPEN_A_VALVE = 1

type Valve interface {
	GetName() string
	GetFlowRate() int
	GetReachableValves() []string
}

type Path struct {
	currentValve         string
	visitedValves        map[string]struct{}
	visitedValvesInOrder []string
	releasedPressure     int
	remainingMinutes     int
}

func GetMostReleasablePressure(inputPath string) int {
	valves := parseValves(inputPath)
	distanceTable := getDistanceTable(valves)
	paths := map[string]Path{
		"AA": {
			currentValve:         "AA",
			visitedValves:        map[string]struct{}{"AA": {}},
			visitedValvesInOrder: []string{},
			releasedPressure:     0,
			remainingMinutes:     30,
		},
	}

	for {
		newPaths := map[string]Path{}
		finalPaths := map[string]Path{}
		for previousValves, path := range paths {
			pathContinues := false
			for targetValve, distance := range distanceTable[path.currentValve] {
				if utils.MapContains(path.visitedValves, targetValve) {
					continue
				}

				remainingMinutes := path.remainingMinutes - distance - MINUTES_TO_OPEN_A_VALVE
				if remainingMinutes < 0 {
					continue
				}

				additionalReleasedPressure := valves[targetValve].GetFlowRate() * remainingMinutes
				if additionalReleasedPressure == 0 {
					continue
				}

				pathContinues = true

				visitedValves := utils.CopyMap(path.visitedValves, map[string]struct{}{})
				visitedValves[targetValve] = struct{}{}
				releasedPressure := path.releasedPressure + additionalReleasedPressure

				newPath := Path{
					currentValve:         targetValve,
					visitedValves:        visitedValves,
					visitedValvesInOrder: append(path.visitedValvesInOrder, targetValve),
					releasedPressure:     releasedPressure,
					remainingMinutes:     remainingMinutes,
				}

				currentValves := previousValves + "," + targetValve
				if !utils.MapContains(newPaths, currentValves) ||
					newPath.releasedPressure > newPaths[currentValves].releasedPressure {
					newPaths[currentValves] = newPath
				}
			}

			if !pathContinues {
				finalPaths[previousValves] = path
			}
		}

		if len(newPaths) == 0 {
			break
		}

		paths = newPaths
		for k, v := range finalPaths {
			paths[k] = v
		}
	}

	mostReleasablePressurePath := paths["TM"]
	for _, path := range paths {
		if path.releasedPressure > mostReleasablePressurePath.releasedPressure {
			mostReleasablePressurePath = path
		}
	}

	return mostReleasablePressurePath.releasedPressure
}

func parseValves(inputPath string) map[string]Valve {
	valves := map[string]Valve{}
	utils.ProcessInputLines(inputPath, func(line string) {
		valve := valve.Parse(line)
		valves[valve.GetName()] = valve
	})
	return valves
}
