package queue

type Queue struct {
	elements []interface{}
}

func (q *Queue) Push(elem interface{}) {
	q.elements = append(q.elements, elem)
}

func (q *Queue) Pop() {
	if q.Empty() {
		panic("pop empty queue")
	}
	q.elements = q.elements[1:]
}

func (q *Queue) Front() interface{} {
	if q.Empty() {
		return nil
	}
	return q.elements[0]
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}
