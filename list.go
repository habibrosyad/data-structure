package structure

import "fmt"

// List abstract data type using linked list
type List struct {
	head *Node
	size int
}

// Linked list node
type Node struct {
	data interface{}
	next *Node
}

// Create a new List
func NewList() *List {
	return &List{nil, 0}
}

// Check whether the List is empty or not
func (l *List) IsEmpty() bool {
	return l.size == 0
}

// Get the size of the List
func (l *List) Size() int {
	return l.size
}

// Add a new item at the beginning of the List
func (l *List) Add(item interface{}) {
	node := &Node{item, l.head}
	l.head = node
	l.size++
}

// Remove an item from the List
func (l *List) Remove(item interface{}) {
	var previous, current *Node = nil, l.head

	for current != nil {
		if current.data == item {
			if previous == nil {
				l.head = current.next
			} else {
				previous.next = current.next
			}
			current.next = nil
			l.size--
			break
		}
		previous = current
		current = current.next
	}
}

// Find whether an item exist in the List
func (l *List) Search(item interface{}) bool {
	current := l.head

	for current != nil {
		if current.data == item {
			return true
		}
		current = current.next
	}

	return false
}

// Insert a new item at the end of the List
func (l *List) Append(item interface{}) {
	var previous, current *Node = nil, l.head

	for current != nil {
		previous = current
		current = current.next
	}

	if previous == nil {
		l.Add(item)
	} else {
		previous.next = &Node{item, nil}
		l.size++
	}
}

// Get the index of an item in the List
func (l *List) Index(item interface{}) int {
	current := l.head

	for i := 0; i < l.size; i++ {
		if current.data == item {
			return i
		}
		current = current.next
	}

	return -1
}

// Insert a new item into the List at any position
func (l *List) Insert(pos int, item interface{}) {
	if pos == l.size {
		l.Append(item)
	} else if pos == 0 {
		l.Add(item)
	} else if pos < l.size {
		var previous, current *Node = nil, l.head

		for i := 0; i < pos; i++ {
			previous = current
			current = current.next
		}

		if previous != nil {
			node := &Node{item, nil}
			node.next = current
			previous.next = node
			l.size++
		}
	}
}

// Retrieve an item from the List at any position, removing it from the List
func (l *List) Pop(pos int) interface{} {
	if pos < l.size {
		var previous, current *Node = nil, l.head

		for i := 0; i < pos; i++ {
			previous = current
			current = current.next
		}

		if current != nil {
			v := current.data

			if previous == nil {
				l.head = current.next
			} else {
				previous.next = current.next
			}
			current.next = nil

			l.size--

			return v
		}
	}

	return nil
}

// Retrieve and remove the list item of the List
func (l *List) PopLast() interface{} {
	return l.Pop(l.size - 1)
}

// String representation of the list
func (l *List) String() string {
	s := []interface{}{}
	current := l.head

	for current != nil {
		s = append(s, current.data)
		current = current.next
	}

	return fmt.Sprintf("%v", s)
}
