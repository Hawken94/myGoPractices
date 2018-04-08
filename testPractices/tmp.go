package main

import (
	"fmt"
	"sort"
	"testing"
)

var (
	intA     int
	float32A float32
	float64A float64
	boolA    bool
	stringA  string
	arrA     []int
	sliceA   []int
	mapA     map[string]int
	pointA   *int
)

func main() {
	fmt.Println(intA == 0)
	fmt.Println(float32A == 0)
	fmt.Println(float64A == 0)
	fmt.Println(boolA == false)
	fmt.Println(stringA == "")
	fmt.Println(arrA == nil)
	fmt.Println(sliceA == nil)
	fmt.Println(mapA == nil)
	fmt.Println(pointA == nil)

	// 对 map 的 key 进行排序
	// map 本身是无序的,需要将 key 取出来放到一个 slice 中,对 slice 进行排序
	temp_map := map[string]int{"e": 1, "f": 2, "c": 3, "g": 4, "w": 5}
	temp_slice := make([]string, 0)
	for k := range temp_map {
		temp_slice = append(temp_slice, k)
	}
	// 调用 sort 包的 string 排序方法
	sort.Strings(temp_slice)
	fmt.Println(temp_slice)
	for _, k := range temp_slice {
		fmt.Printf("\"%s\":%v", k, temp_map[k])
	}

	// 追加切片
	temp_a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	temp_b := [5]int{10, 11, 12, 13, 14}
	slice_c := make([]int, 0)
	for i := 0; i < len(temp_a); i++ {
		slice_c = append(slice_c, temp_a[i])
	}
	for i := 0; i < len(temp_b); i++ {
		slice_c = append(slice_c, temp_b[i])
	}
	fmt.Println(slice_c)

	// testing

	info := &Info{"北京"}
	info.changeInfo() //
	fmt.Println(info) //

	info2 := Info{"北京"}
	info2.changeInfo2() //
	fmt.Println(info2)  //

	fmt.Println()
	// 100以内素数之和
	sum := 0
	for i := 2; i < 100; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Printf("100以内素数之和是:%v \n", sum)
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

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
