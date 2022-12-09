package campcleanup

import (
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func createRangeFromString(s string) Range {
	boundaries := strings.Split(s, "-")
	start, _ := strconv.Atoi(boundaries[0])
	end, _ := strconv.Atoi(boundaries[1])
	return Range{
		start: start,
		end:   end,
	}
}

func (r Range) fullyContains(other Range) bool {
	return r.start <= other.start && r.end >= other.end
}

func rangesFullyContain(r1 Range, r2 Range) bool {
	return r1.fullyContains(r2) || r2.fullyContains(r1)
}

func rangesOverlap(r1 Range, r2 Range) bool {
	return r1.start <= r2.end && r1.end >= r2.start
}
