package comparator

import "aoc-2022-go/13_distresssignal/types/intorlist"

func IsPairInCorrectOrder(left *intorlist.IntOrList, right *intorlist.IntOrList) bool {
	return comparePairs(left, right) == VALID
}

type Validity int

const (
	VALID Validity = iota
	INVALID
	UNDEFINED
)

func comparePairs(left *intorlist.IntOrList, right *intorlist.IntOrList) Validity {
	if left.ValueType == intorlist.INT && right.ValueType == intorlist.INT {
		return compareInts(left.IntValue, right.IntValue)
	} else {
		return compareLists(left.GetListValue(), right.GetListValue())
	}
}

func compareInts(left int, right int) Validity {
	if left < right {
		return VALID
	} else if left > right {
		return INVALID
	}

	return UNDEFINED
}

func compareLists(left []*intorlist.IntOrList, right []*intorlist.IntOrList) Validity {
	i := 0
	leftCount := len(left)
	rightCount := len(right)
	for {
		if leftCount == i && rightCount == i {
			return UNDEFINED
		}
		if leftCount == i {
			return VALID
		}
		if rightCount == i {
			return INVALID
		}

		validity := comparePairs(left[i], right[i])
		if validity != UNDEFINED {
			return validity
		}

		i++
	}
}
