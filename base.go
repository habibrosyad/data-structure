package structure

import "fmt"

type Base struct {
	items []interface{}
}

func (b *Base) IsEmpty() bool {
	return len(b.items) > 0
}

func (b *Base) Size() int {
	return len(b.items)
}

func (b *Base) String() string {
	return fmt.Sprintf("%v", b.items)
}
