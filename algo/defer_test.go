package algo

import (
	"errors"
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

func TestDefer2(t *testing.T) {
	err := doStuff()
	fmt.Println("123123", err)
	t.Error(err)
}

func pressButton() {
	fmt.Println("I'm Mr. Meeseeks, look at me!!")
	// other stuff then happens, but if Jerry asks to
	// remove 2 strokes from his golf game...
	panic("It's gettin' weird!")
}

func doStuff() (err error) {
	// var err error
	// If there is a panic we need to recover in a deferred func
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("the meeseeks went crazy!")
		}
	}()

	pressButton()
	fmt.Println("we got here!")
	return err
}
