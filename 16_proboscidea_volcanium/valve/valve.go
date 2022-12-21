package valve

import (
	"strconv"
	"strings"
)

type Valve struct {
	name            string
	flowRate        int
	reachableValves []string
}

func (v Valve) GetName() string {
	return v.name
}

func (v Valve) GetFlowRate() int {
	return v.flowRate
}

func (v Valve) GetReachableValves() []string {
	return v.reachableValves
}

func Parse(s string) Valve {
	split := strings.Split(strings.SplitN(s, "Valve ", 2)[1], " has flow rate=")
	name := split[0]
	if strings.Contains(split[1], "; tunnels lead to valves ") {
		split = strings.Split(split[1], "; tunnels lead to valves ")
	} else {
		split = strings.Split(split[1], "; tunnel leads to valve ")
	}
	flowRate, _ := strconv.Atoi(split[0])
	valves := strings.Split(split[1], ", ")

	return Valve{
		name:            name,
		flowRate:        flowRate,
		reachableValves: valves,
	}
}
