package beacon

import (
	"aoc-2022-go/utils"
	"strconv"
	"strings"
)

func getSensorsAndXLimitsFromFile(inputPath string, distanceCalculator DistanceCalculator) ([]Sensor, map[string]struct{}, int, int) {
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

	return sensors, beaconPositions, minX, maxX
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
