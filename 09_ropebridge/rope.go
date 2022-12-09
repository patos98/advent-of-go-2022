package ropebridge

import "math"

type Rope interface {
	moveHead(Rules, Step)
	moveTail()
	getTail() string
}

type SimpleRope struct {
	head Position
	tail Position
}

func newSimpleRope() *SimpleRope {
	return &SimpleRope{
		head: startingPosition(),
		tail: startingPosition(),
	}
}

func (r *SimpleRope) moveHead(rules Rules, step Step) {
	r.head = rules[step.direction](r.head, 1)
}

func (r *SimpleRope) moveTail() {
	r.tail = getNextTailPosition(r.head, r.tail)
}

func getNextTailPosition(head Position, tail Position) Position {
	diff := head.getDifference(tail)
	isAdjacent := math.Abs(float64(diff.x)) <= 1 && math.Abs(float64(diff.y)) <= 1
	if isAdjacent {
		return tail
	}

	maxTailStep := 1
	movement := Position{
		x: trimCoordinate(diff.x, maxTailStep),
		y: trimCoordinate(diff.y, maxTailStep),
	}

	return Position{
		x: tail.x + movement.x,
		y: tail.y + movement.y,
	}
}

func (r *SimpleRope) getTail() string {
	return r.tail.toString()
}

type SuperRope struct {
	head  Position
	tails []Position
}

func newSuperRope(tailCount int) *SuperRope {
	tails := []Position{}
	for i := 0; i < tailCount; i++ {
		tails = append(tails, startingPosition())
	}
	return &SuperRope{
		head:  startingPosition(),
		tails: tails,
	}
}

func (r *SuperRope) moveHead(rules Rules, step Step) {
	r.head = rules[step.direction](r.head, 1)
}

func (r *SuperRope) moveTail() {
	tailCount := len(r.tails)
	head := r.head
	for i := 0; i < tailCount; i++ {
		r.tails[i] = getNextTailPosition(head, r.tails[i])
		head = r.tails[i]
	}
}

func (r *SuperRope) getTail() string {
	return r.tails[len(r.tails)-1].toString()
}
