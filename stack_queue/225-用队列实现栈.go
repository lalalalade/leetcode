package stack_queue

type MyStack1 struct {
	q []int
}

func Constructor1() MyStack1 {
	return MyStack1{
		q: []int{},
	}
}

func (s *MyStack1) Push(x int) {
	n := len(s.q)
	// x 加入
	s.q = append(s.q, x)
	// 依次弹出加入到后面
	for i := 0; i < n; i++ {
		val := s.q[0]
		s.q = s.q[1:]
		s.q = append(s.q, val)
	}
}

func (s *MyStack1) Pop() int {
	val := s.q[0]
	s.q = s.q[1:]
	return val
}

func (s *MyStack1) Top() int {
	return s.q[0]
}

func (s *MyStack1) Empty() bool {
	return len(s.q) == 0
}
