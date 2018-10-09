package tests

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPointer(t *testing.T) {

	x := 100
	var p = &x
	var e interface{} = &x

	type eface struct {
		itable uintptr
		data   *int
	}

	ep := (*eface)(unsafe.Pointer(&e))
	t.Errorf("%v \n", e)
	t.Errorf("%v \n", ep)
	t.Errorf("%p %v \n", p, *ep.data)
}

func TestPointerArr(t *testing.T) {
	var data [24]byte
	p := (*[24 / unsafe.Sizeof(0)]int)(unsafe.Pointer(&data))

	fmt.Printf("%p \n", p)
	fmt.Printf("%p \n", p)

}
