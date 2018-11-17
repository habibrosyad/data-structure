package structure

import "testing"

func TestDequeAddFront(t *testing.T) {
	d := NewDeque()
	d.AddFront(1)
	d.AddFront(2)

	if d.Size() != 2 {
		t.Errorf("Expect 2 got %v", d.Size())
	}
}

func TestDequeRemoveFront(t *testing.T) {
	d := NewDeque()
	d.AddFront(1)
	d.AddFront(2)

	for _, expect := range []interface{}{2, 1} {
		if item := d.RemoveFront(); item != expect {
			t.Errorf("Expect %v got %v", expect, item)
		}
	}
}

func TestDequeAddRear(t *testing.T) {
	d := NewDeque()
	d.AddRear(1)
	d.AddRear(2)

	if d.Size() != 2 {
		t.Errorf("Expect 2 got %v", d.Size())
	}
}

func TestDequeRemoveRear(t *testing.T) {
	d := NewDeque()
	d.AddRear(1)
	d.AddRear(2)

	for _, expect := range []interface{}{2, 1} {
		if item := d.RemoveRear(); item != expect {
			t.Errorf("Expect %v got %v", expect, item)
		}
	}
}
