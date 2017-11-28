package algo

import (
	"fmt"
	"testing"
)

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

func TestDefer(t *testing.T) {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
	t.Error("xhk")
}
