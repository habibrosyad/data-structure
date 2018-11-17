package structure

// Stack abstract data type, using slice as the base
type Stack struct {
	*Base
}

// Create a new Stack
func NewStack() *Stack {
	return &Stack{&Base{[]interface{}{}}}
}

// Insert a new item into the Stack
func (s *Stack) Push(item interface{}) {
	s.Base.items = append(s.Base.items, item)
}

// Remove an item from the Stack and return it in LIFO order
func (s *Stack) Pop() interface{} {
	v := s.Base.items[len(s.Base.items)-1]  // Data to be popped
	s.Base.items[len(s.Base.items)-1] = nil // Erase data
	s.Base.items = s.Base.items[:len(s.Base.items)-1]
	return v
}

// Return the last inserted item from the Stack without removing it
func (s *Stack) Peek() interface{} {
	return s.Base.items[len(s.Base.items)-1]
}
