package proboscideavolcanium

import "aoc-2022-go/utils"

type DistanceTable map[string]map[string]int

func getDistanceTable(valves map[string]Valve) DistanceTable {
	distanceTable := DistanceTable{}
	for valveName, valve := range valves {
		distanceTable[valveName] = getDistancesFromValve(valve, valves)
	}
	return distanceTable
}

func getDistancesFromValve(valve Valve, valves map[string]Valve) map[string]int {
	currentDistance := 0
	distances := map[string]int{valve.GetName(): currentDistance}

	currentValves := []Valve{valve}
	for len(currentValves) > 0 {
		currentDistance++
		nextValves := []Valve{}

		for _, currentValve := range currentValves {
			for _, nextValve := range currentValve.GetReachableValves() {
				if !utils.MapContains(distances, nextValve) {
					distances[nextValve] = currentDistance
					nextValves = append(nextValves, valves[nextValve])
				}
			}
		}

		currentValves = nextValves
	}

	return distances
}
