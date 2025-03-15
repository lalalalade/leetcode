package stack_queue

type MyQueue struct {
	stackIn  *MyStack
	stackOut *MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  &MyStack{},
		stackOut: &MyStack{},
	}
}

func (q *MyQueue) Push(x int) {
	q.stackIn.Push(x)
}

func (q *MyQueue) Pop() int {
	q.fillOut()
	return q.stackOut.Pop()
}

func (q *MyQueue) Peek() int {
	q.fillOut()
	return q.stackOut.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.stackIn.Empty() && q.stackOut.Empty()
}

func (q *MyQueue) fillOut() {
	if q.stackOut.Empty() {
		for !q.stackIn.Empty() {
			q.stackOut.Push(q.stackIn.Pop())
		}
	}
}

type MyStack []int

func (s *MyStack) Push(x int) {
	*s = append(*s, x)
}

func (s *MyStack) Pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func (s *MyStack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *MyStack) Empty() bool {
	return len(*s) == 0
}
