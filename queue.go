package structure

type Queue struct {
	*Base
}

func NewQueue() *Queue {
	return &Queue{&Base{[]interface{}{}}}
}

func (q *Queue) Enqueue(item interface{}) {
	q.Base.items = append(q.Base.items, item)
}

func (q *Queue) Dequeue() interface{} {
	v := q.Base.items[0]
	q.Base.items = q.Base.items[1:]
	return v
}
