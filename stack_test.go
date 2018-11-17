package structure

import "testing"

func TestStackPush(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	if s.Size() != 2 {
		t.Errorf("Expect %v got %v", 2, s.Size())
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	for _, expect := range []interface{}{2, 1} {
		if item := s.Pop(); item != expect {
			t.Errorf("Expect %v got %v", expect, item)
		}
	}

	if s.Size() != 0 {
		t.Errorf("Expect 0 got %v", s.Size())
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)

	if v := s.Peek(); v != 2 {
		t.Errorf("Expect 2 got %v", v)
	}

	if s.Size() != 2 {
		t.Errorf("Expect 2 got %v", s.Size())
	}
}
