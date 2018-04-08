package main

import "fmt"

func main() {
	num := 0
	a := 0
	b := 0
	c := 0
	for {
		n, _ := fmt.Scan(&num)
		if n == 0 {
			break
		} else {
			fmt.Printf("请输入%d个数字\n", num)
			n, _ := fmt.Scan(&a, &b, &c)
			if n == 0 {
				break
			} else {
				fmt.Println(a, b, c)
			}
		}
	}
}
