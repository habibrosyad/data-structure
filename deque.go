package structure

// Deque abstract data type
type Deque struct {
	*Base
}

// Create a new Deque
func NewDeque() *Deque {
	return &Deque{&Base{[]interface{}{}}}
}

// Add a new item at the rear of the queue
func (d *Deque) AddRear(item interface{}) {
	d.Base.items = append(d.Base.items, item)
}

// Retrieve and return an item from the rear of the queue
func (d *Deque) RemoveRear() interface{} {
	if len(d.Base.items) > 0 {
		v := d.Base.items[len(d.Base.items)-1]  // Data to be popped
		d.Base.items[len(d.Base.items)-1] = nil // Erase data
		d.Base.items = d.Base.items[:len(d.Base.items)-1]
		return v
	}
	return nil
}

// Add a new item at the front of the queue
func (d *Deque) AddFront(item interface{}) {
	d.Base.items = append([]interface{}{item}, d.Base.items...)
}

// Retrieve and return an item from the front of the queue
func (d *Deque) RemoveFront() interface{} {
	if len(d.Base.items) > 0 {
		v := d.Base.items[0]
		d.Base.items = d.Base.items[1:]
		return v
	}
	return nil
}
