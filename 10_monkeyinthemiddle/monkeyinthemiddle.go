package monkeyinthemiddle

import "sort"

func GetMonkeyBusiness() int {
	monkeys := getMonkeysManually()
	inspectCounts := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	rounds := 20
	for roundIndex := 0; roundIndex < rounds; roundIndex++ {
		for monkeyIndex, monkey := range monkeys {
			for _, item := range monkey.StartingItems {
				inspectCounts[monkeyIndex]++
				newItem := monkey.Operation(item) / 3
				targetMonkeyIndex := monkey.Test(newItem)
				monkeys[targetMonkeyIndex].StartingItems = append(monkeys[targetMonkeyIndex].StartingItems, newItem)
			}
			monkey.StartingItems = []int{}
		}
	}

	sort.Ints(inspectCounts[:])

	return inspectCounts[6] * inspectCounts[7]
}
