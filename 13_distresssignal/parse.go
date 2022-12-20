package distresssignal

import (
	"aoc-2022-go/13_distresssignal/types/intorlist"
	"strconv"
	"strings"
)

func parse(s string) *intorlist.IntOrList {
	s = s[1 : len(s)-1] // remove [] from outside

	result := intorlist.NewList(nil)
	currentNode := result
	for {

		if len(s) == 0 {
			break
		}

		nextChar := s[0]
		if nextChar == '[' {
			newNode := intorlist.NewList(currentNode)
			currentNode.ListValue = append(currentNode.ListValue, newNode)
			currentNode = newNode
			s = s[1:]
		} else if nextChar == ']' {
			currentNode = currentNode.Parent
			s = s[1:]
		} else if nextChar != ',' {
			// 1,2,3,[4,5,6]]...
			split := strings.SplitN(s, ",", 2)
			numberString := strings.Split(split[0], "]")[0]
			value, err := strconv.Atoi(numberString)
			if err != nil {
				panic(err)
			}
			currentNode.ListValue = append(currentNode.ListValue, intorlist.NewInt(value, currentNode))
			if len(split) > 1 {
				s = split[1]
				if split[0][len(split[0])-1] == ']' {
					s = "]" + s
				}
			} else {
				s = ""
			}
		} else {
			s = s[1:]
		}

	}

	return result
}
