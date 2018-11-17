package structure

import "testing"

func TestListAdd(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)

	if l.Size() != 2 {
		t.Errorf("Expect 2 got %v", l.Size())
	}
}

func TestListRemove(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)

	expects := []struct {
		item interface{}
		size int
	}{
		{3, 2},
		{2, 1},
		{4, 1},
		{1, 0},
		{5, 0},
	}

	for _, expect := range expects {
		l.Remove(expect.item)
		if l.Size() != expect.size {
			t.Errorf("Expect %v got %v", expect.size, l.Size())
		}
	}
}

func TestListSearch(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)

	expects := []struct {
		item  interface{}
		found bool
	}{
		{3, false},
		{2, true},
		{4, false},
		{1, true},
		{5, false},
	}

	for _, expect := range expects {
		if found := l.Search(expect.item); found != expect.found {
			t.Errorf("Expect %v got %v", expect.found, found)
		}
	}
}

func TestListAppend(t *testing.T) {
	l := NewList()
	l.Append(1)
	l.Append(2)

	if l.Size() != 2 {
		t.Errorf("Expect 2 got %v", l.Size())
	}
}

func TestListIndex(t *testing.T) {
	l := NewList()
	l.Add(3)
	l.Add(4)
	l.Append(1)
	l.Append(2)

	expects := []struct {
		item interface{}
		pos  int
	}{
		{3, 1},
		{2, 3},
		{4, 0},
		{1, 2},
		{5, -1},
	}

	for _, expect := range expects {
		if pos := l.Index(expect.item); pos != expect.pos {
			t.Errorf("Expect %v got %v", expect.pos, pos)
		}
	}
}

func TestListInsert(t *testing.T) {
	l := NewList()

	expects := []struct {
		pos  int
		item interface{}
		size int
	}{
		{3, 1, 0},
		{0, 1, 1},
		{1, 2, 2},
		{0, 3, 3},
		{1, 4, 4},
	}

	for _, expect := range expects {
		l.Insert(expect.pos, expect.item)

		// Check size after insertion
		if l.Size() != expect.size {
			t.Errorf("Expect %v got %v", expect.size, l.Size())
		}

		// Check index after insertion
		if pos := l.Index(expect.item); expect.size > 0 && pos != expect.pos {
			t.Errorf("Expect %v got %v", expect.pos, pos)
		}
	}
}

func TestListPop(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)
	l.Append(3)
	l.Append(4)

	expects := []struct {
		pos  int
		item interface{}
	}{
		{4, nil},
		{0, 2},
		{1, 3},
		{0, 1},
		{0, 4},
	}

	for _, expect := range expects {
		if item := l.Pop(expect.pos); item != expect.item {
			t.Errorf("Expect %v got %v", expect.item, item)
		}
	}

	if l.Size() != 0 {
		t.Errorf("Expect 0 got %v", l.Size())
	}
}

func TestListPopLast(t *testing.T) {
	l := NewList()
	l.Add(1)
	l.Add(2)
	l.Append(3)
	l.Append(4)

	for _, expect := range []interface{}{4, 3, 1, 2} {
		if item := l.PopLast(); item != expect {
			t.Errorf("Expect %v got %v", expect, item)
		}
	}

	if l.Size() != 0 {
		t.Errorf("Expect 0 got %v", l.Size())
	}
}
