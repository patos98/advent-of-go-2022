package monkeyinthemiddle

import (
	"sort"
)

func GetMonkeyBusiness() int {
	return getMonkeyBusiness(getMonkeysManually(), 20, func(i int) int { return i / 3 })
}

func GetMonkeyBusiness2() int {
	product := int(2 * 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23)
	return getMonkeyBusiness(getMonkeysManually(), 10000, func(i int) int {
		if i < product {
			return i
		}

		return i % product
	})
}

func getMonkeyBusiness(monkeys []*Monkey, rounds int, normalizer func(int) int) int {
	inspectCounts := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	for roundIndex := int(0); roundIndex < rounds; roundIndex++ {
		for monkeyIndex, monkey := range monkeys {
			for _, item := range monkey.StartingItems {
				inspectCounts[monkeyIndex]++
				newItem := normalizer(monkey.Operation(item))
				targetMonkeyIndex := monkey.Test(newItem)
				monkeys[targetMonkeyIndex].StartingItems = append(monkeys[targetMonkeyIndex].StartingItems, newItem)
			}
			monkey.StartingItems = []int{}
		}
	}

	sort.Ints(inspectCounts[:])

	return inspectCounts[int(6)] * inspectCounts[int(7)]
}
