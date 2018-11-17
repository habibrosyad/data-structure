package structure

type Stack struct {
	*Base
}

func NewStack() *Stack {
	return &Stack{&Base{[]interface{}{}}}
}

func (s *Stack) Push(item interface{}) {
	s.Base.items = append(s.Base.items, item)
}

func (s *Stack) Pop() interface{} {
	v := s.Base.items[len(s.Base.items)-1]  // Data to be popped
	s.Base.items[len(s.Base.items)-1] = nil // Erase data
	s.Base.items = s.Base.items[:len(s.Base.items)-1]
	return v
}

func (s *Stack) Peek() interface{} {
	return s.Base.items[len(s.Base.items)-1]
}
