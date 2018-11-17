package structure

import "fmt"

// We use `slice` of `interface{}` for many abstract data type defined in this package
type Base struct {
	items []interface{}
}

// Check whether the collection is empty or not
func (b *Base) IsEmpty() bool {
	return len(b.items) > 0
}

// Get the size of the collection
func (b *Base) Size() int {
	return len(b.items)
}

// String representation of the collection
func (b *Base) String() string {
	return fmt.Sprintf("%v", b.items)
}
