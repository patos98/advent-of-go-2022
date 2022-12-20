package beacon

import (
	"aoc-2022-go/utils"
	"math"
)

const INPUT_PATH = "15_beacon/input.txt"
const TEST_INPUT_PATH = "15_beacon/input_test.txt"

func GetNumberOfNonBeaconPositionInRow(inputPath string, y int) int {
	distanceCalculator := manhattanDistance
	sensors, beaconPositions, minX, maxX := getSensorsAndXLimitsFromFile(inputPath, distanceCalculator)

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

type Range struct {
	min int
	max int
}

func GetDistressBeaconFrequency(inputPath string, coordinatesMin int, coordinatesMax int) int {
	distanceCalculator := manhattanDistance
	sensors, _, _, _ := getSensorsAndXLimitsFromFile(inputPath, distanceCalculator)

	stopInit := utils.Timer("init")
	possibleBeaconX := map[int][]Range{}
	for i := coordinatesMin; i <= coordinatesMax; i++ {
		possibleBeaconX[i] = []Range{}
	}
	stopInit()

	stopSensors := utils.Timer("sensors")
	for _, sensor := range sensors {
		stopSingleSensor := utils.Timer("single sensor")
		distanceX := sensor.distance
		for i := -distanceX; i <= distanceX; i++ {
			x := sensor.self.x + i
			if x < coordinatesMin || x > coordinatesMax {
				continue
			}

			distanceY := distanceX - int(math.Abs(float64(i)))
			minY := sensor.self.y - distanceY
			maxY := sensor.self.y + distanceY

			possibleBeaconX[x] = utils.InsertIntoSortedSlice(possibleBeaconX[x], Range{
				min: minY,
				max: maxY,
			}, Range{}, func(itemInSlice Range, itemToInsert Range) bool {
				return itemInSlice.min > itemToInsert.min
			})
		}
		stopSingleSensor()
	}
	stopSensors()

	defer utils.Timer("finding distress beacon")()
	for x, yRanges := range possibleBeaconX {
		maxY := yRanges[0].max
		for _, yRange := range yRanges {
			if yRange.min > maxY+1 {
				return 4000000*x + maxY + 1
			}
			if yRange.max > maxY {
				maxY = yRange.max
			}
		}
	}

	return 0
}

func manhattanDistance(p1 Position, p2 Position) int {
	return int(math.Abs(float64(p1.x)-float64(p2.x))) +
		int(math.Abs(float64(p1.y)-float64(p2.y)))
}
