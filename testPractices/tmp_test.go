package testPractices_test

import (
	"fmt"
	"testing"
)

type Info struct {
	address string
}

func (this Info) changeInfo() {
	this.address = "上海"
	fmt.Println(this.address)
}

func (this *Info) changeInfo2() {
	this.address = "上海"
	fmt.Println(this.address)
}

func Test_demo5(t *testing.T) {
	info := &Info{"北京"}
	info.changeInfo() //
	fmt.Println(info) //

	info2 := Info{"北京"}
	info2.changeInfo2() //
	fmt.Println(info2)  //
}
