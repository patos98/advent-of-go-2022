package proboscideavolcanium

import (
	"aoc-2022-go/16_proboscidea_volcanium/valve"
	"aoc-2022-go/utils"
)

const INPUT_PATH = "16_proboscidea_volcanium/input.txt"
const TEST_INPUT_PATH = "16_proboscidea_volcanium/input_test.txt"

type Valve interface {
	GetName() string
	GetFlowRate() int
	GetReachableValves() []string
}

type Path struct {
	currentValve     string
	visitedValves    map[string]struct{}
	openValves       map[string]struct{}
	releasedPressure int
}

func GetMostReleasablePressure(inputPath string) int {
	valves := parseValves(inputPath)
	paths := []Path{{
		currentValve:     "AA",
		visitedValves:    map[string]struct{}{},
		openValves:       map[string]struct{}{},
		releasedPressure: 0,
	}}

	remainingMinutes := 30
	maxReleasedPressureAtValves := map[string]int{}

	for remainingMinutes > 0 {
		remainingMinutes--
		newPaths := []Path{}

		for _, path := range paths {
			currentValve := valves[path.currentValve]

			// add next possible valves to path
			for _, nextValve := range currentValve.GetReachableValves() {
				visitedValves := utils.CopyMap(path.visitedValves, map[string]struct{}{})

				if !utils.MapContains(visitedValves, nextValve) {
					visitedValves[nextValve] = struct{}{}

					if !utils.MapContains(maxReleasedPressureAtValves, nextValve) ||
						path.releasedPressure > maxReleasedPressureAtValves[nextValve] {
						maxReleasedPressureAtValves[nextValve] = path.releasedPressure

						newPaths = append(newPaths, Path{
							currentValve:     nextValve,
							visitedValves:    visitedValves,
							openValves:       utils.CopyMap(path.openValves, map[string]struct{}{}),
							releasedPressure: path.releasedPressure,
						})
					}
				} else {
					newPaths = append(newPaths, Path{
						currentValve:     nextValve,
						visitedValves:    visitedValves,
						openValves:       utils.CopyMap(path.openValves, map[string]struct{}{}),
						releasedPressure: path.releasedPressure,
					})
				}
			}

			// add opening current valve to path
			if !utils.MapContains(path.openValves, path.currentValve) &&
				valves[path.currentValve].GetFlowRate() > 0 {
				openValves := utils.CopyMap(path.openValves, map[string]struct{}{})
				openValves[path.currentValve] = struct{}{}
				releasedPressure := remainingMinutes * valves[path.currentValve].GetFlowRate()

				newPaths = append(newPaths, Path{
					currentValve:     path.currentValve,
					visitedValves:    path.visitedValves,
					openValves:       path.openValves,
					releasedPressure: path.releasedPressure + releasedPressure,
				})
			}

		}

		paths = newPaths
	}

	maxReleasablePressure := 0
	for _, path := range paths {
		if path.releasedPressure > maxReleasablePressure {
			maxReleasablePressure = path.releasedPressure
		}
	}

	return maxReleasablePressure
}

func parseValves(inputPath string) map[string]Valve {
	valves := map[string]Valve{}
	utils.ProcessInputLines(inputPath, func(line string) {
		valve := valve.Parse(line)
		valves[valve.GetName()] = valve
	})
	return valves
}
