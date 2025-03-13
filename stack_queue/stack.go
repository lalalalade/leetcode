package stack_queue

type Stack struct {
	elements []int
}

func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1
	}
	index := len(s.elements) - 1
	value := s.elements[index]
	s.elements = s.elements[:len(s.elements)-1]
	return value
}

func (s *Stack) Peek() int {
	if len(s.elements) == 0 {
		return -1
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
