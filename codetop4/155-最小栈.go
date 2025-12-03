package codetop4

import "math"

type Pair struct {
	val, preMin int
}

type MinStack []Pair

func Constructor2() MinStack {
	return MinStack{{0, math.MaxInt}}
}

func (s *MinStack) Push(val int) {
	*s = append(*s, Pair{val, min(val, s.GetMin())})
}

func (s *MinStack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *MinStack) Top() int {
	return (*s)[len(*s)-1].val
}

func (s *MinStack) GetMin() int {
	return (*s)[len(*s)-1].preMin
}
