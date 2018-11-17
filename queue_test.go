package structure

import "testing"

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)

	if q.Size() != 2 {
		t.Errorf("Expect 2 got %v", q.Size())
	}
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)

	for _, expect := range []interface{}{1, 2} {
		if item := q.Dequeue(); item != expect {
			t.Errorf("Expect %v got %v", expect, item)
		}
	}

	if q.Size() != 0 {
		t.Errorf("Expect 0 got %v", q.Size())
	}
}
