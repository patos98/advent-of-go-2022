package monkeyinthemiddle

type Monkey struct {
	StartingItems []int
	Operation     func(int) int
	Test          func(int) int // returns the index of the target monkey
}

func getMonkeysManually() []*Monkey {
	return []*Monkey{
		{
			StartingItems: []int{72, 97},
			Operation:     func(i int) int { return i * 13 },
			Test: func(i int) int {
				if i%19 == 0 {
					return 5
				} else {
					return 6
				}
			},
		},
		{
			StartingItems: []int{55, 70, 90, 74, 95},
			Operation:     func(i int) int { return i * i },
			Test: func(i int) int {
				if i%7 == 0 {
					return 5
				} else {
					return 0
				}
			},
		},
		{
			StartingItems: []int{74, 97, 66, 57},
			Operation:     func(i int) int { return i + 6 },
			Test: func(i int) int {
				if i%17 == 0 {
					return 1
				} else {
					return 0
				}
			},
		},
		{
			StartingItems: []int{86, 54, 53},
			Operation:     func(i int) int { return i + 2 },
			Test: func(i int) int {
				if i%13 == 0 {
					return 1
				} else {
					return 2
				}
			},
		},
		{
			StartingItems: []int{50, 65, 78, 50, 62, 99},
			Operation:     func(i int) int { return i + 3 },
			Test: func(i int) int {
				if i%11 == 0 {
					return 3
				} else {
					return 7
				}
			},
		},
		{
			StartingItems: []int{90},
			Operation:     func(i int) int { return i + 4 },
			Test: func(i int) int {
				if i%2 == 0 {
					return 4
				} else {
					return 6
				}
			},
		},
		{
			StartingItems: []int{88, 92, 63, 94, 96, 82, 53, 53},
			Operation:     func(i int) int { return i + 8 },
			Test: func(i int) int {
				if i%5 == 0 {
					return 4
				} else {
					return 7
				}
			},
		},
		{
			StartingItems: []int{70, 60, 71, 69, 77, 70, 98},
			Operation:     func(i int) int { return i * 7 },
			Test: func(i int) int {
				if i%3 == 0 {
					return 2
				} else {
					return 3
				}
			},
		},
	}
}
