package beacon

import (
	"aoc-2022-go/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const INPUT_PATH = "15_beacon/input.txt"
const TEST_INPUT_PATH = "15_beacon/input_test.txt"

type Position struct {
	x int
	y int
}

func (p Position) ToString() string {
	return fmt.Sprintf("%d;%d", p.x, p.y)
}

type Sensor struct {
	self          Position
	closestBeacon Position
	distance      int
}

func GetNumberOfNonBeaconPositionInRow(inputPath string, y int) int {
	distanceCalculator := manhattanDistance
	sensors := []Sensor{}
	beaconPositions := map[string]struct{}{}

	minX := 0
	maxX := 0
	utils.ProcessInputLines(inputPath, func(line string) {
		sensor := parseSensor(line, distanceCalculator)
		sensors = append(sensors, sensor)
		beaconPositions[sensor.closestBeacon.ToString()] = struct{}{}

		sensorMinX := sensor.self.x - sensor.distance
		if sensorMinX < minX {
			minX = sensorMinX
		}

		sensorMaxX := sensor.self.x + sensor.distance
		if sensorMaxX > maxX {
			maxX = sensorMaxX
		}
	})

	nonBeaconPositions := 0
	for x := minX; x <= maxX; x++ {
		position := Position{x: x, y: y}
		for _, sensor := range sensors {
			if !utils.MapContains(beaconPositions, position.ToString()) &&
				distanceCalculator(position, sensor.self) <= sensor.distance {
				nonBeaconPositions++
				break
			}
		}
	}

	return nonBeaconPositions
}

type DistanceCalculator func(Position, Position) int

func parseSensor(line string, dc DistanceCalculator) Sensor {
	split := strings.Split(strings.Split(line, "Sensor at x=")[1], ": closest beacon is at x=")
	self := parsePosition(split[0])
	closestBeacon := parsePosition(split[1])
	return Sensor{
		self:          self,
		closestBeacon: closestBeacon,
		distance:      dc(self, closestBeacon),
	}
}

func parsePosition(positionString string) Position {
	positionStrings := strings.Split(positionString, ", y=")
	x, _ := strconv.Atoi(positionStrings[0])
	y, _ := strconv.Atoi(positionStrings[1])
	return Position{
		x: x,
		y: y,
	}
}

func manhattanDistance(p1 Position, p2 Position) int {
	return int(math.Abs(float64(p1.x)-float64(p2.x))) +
		int(math.Abs(float64(p1.y)-float64(p2.y)))
}
