package stack_queue

type Queue struct {
	elements []int
}

// Enqueue 入队
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// Dequeue 出队
func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		return -1
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

func (q *Queue) Peek() int {
	if len(q.elements) == 0 {
		return -1
	}
	return q.elements[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}
