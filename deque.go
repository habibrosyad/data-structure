package structure

type Deque struct {
	*Base
}

func NewDeque() *Deque {
	return &Deque{&Base{[]interface{}{}}}
}

func (d *Deque) AddRear(item interface{}) {
	d.Base.items = append(d.Base.items, item)
}

func (d *Deque) RemoveRear() interface{} {
	if len(d.Base.items) > 0 {
		v := d.Base.items[len(d.Base.items)-1]  // Data to be popped
		d.Base.items[len(d.Base.items)-1] = nil // Erase data
		d.Base.items = d.Base.items[:len(d.Base.items)-1]
		return v
	}
	return nil
}

func (d *Deque) AddFront(item interface{}) {
	d.Base.items = append([]interface{}{item}, d.Base.items...)
}

func (d *Deque) RemoveFront() interface{} {
	if len(d.Base.items) > 0 {
		v := d.Base.items[0]
		d.Base.items = d.Base.items[1:]
		return v
	}
	return nil
}
