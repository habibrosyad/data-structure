package structure

type List struct {
	head *Node
	size int
}

type Node struct {
	data interface{}
	next *Node
}

func NewList() *List {
	return &List{nil, 0}
}

func (l *List) IsEmpty() bool {
	return l.size == 0
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Add(item interface{}) {
	node := &Node{item, l.head}
	l.head = node
	l.size++
}

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

func (l *List) PopLast() interface{} {
	return l.Pop(l.size - 1)
}
