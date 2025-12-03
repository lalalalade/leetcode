package codetop3

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

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	this.fillOut()
	return this.stackOut.Pop()
}

func (this *MyQueue) Peek() int {
	this.fillOut()
	return this.stackIn.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.stackOut.Empty() && this.stackIn.Empty()
}

func (this *MyQueue) fillOut() {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			this.stackOut.Push(this.stackIn.Pop())
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
