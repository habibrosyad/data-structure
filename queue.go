package structure

// Queue abstract data type, using slice as the base
type Queue struct {
	*Base
}

// Create a new Queue
func NewQueue() *Queue {
	return &Queue{&Base{[]interface{}{}}}
}

// Insert a new item to the Queue
func (q *Queue) Enqueue(item interface{}) {
	q.Base.items = append(q.Base.items, item)
}

// Remove an item from the Queue and return it in FIFO order
func (q *Queue) Dequeue() interface{} {
	v := q.Base.items[0]
	q.Base.items = q.Base.items[1:]
	return v
}
