package queue

type Queue struct {
	size     int
	top      int
	queueArr []int
}

func NewQueue(size int) *Queue {
	return &Queue{
		size: size,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.top == 0
}

func (q *Queue) IsFull() bool {
	return q.top >= q.size
}

func (q *Queue) Add(i int) string {
	if q.IsFull() {
		return "cannot add item because queue is full"
	} else {
		q.queueArr = append(q.queueArr, i)
		q.top++
		return "success add item in queue"
	}
}

func (q *Queue) Show() []int {
	if q.IsEmpty() {
		return nil
	} else {
		return q.queueArr
	}
}

func (q *Queue) Pop() string {
	if q.IsEmpty() {
		return "cannot remove item because queue is empty"
	} else {
		q.queueArr = q.queueArr[1:]
		return "success remove item in queue"
	}
}
