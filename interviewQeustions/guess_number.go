package main

import ("fmt")

func main() {
	for i := 1; ; i++ {
		if i%11 == 0 && i%9 == 8 && i%7 == 6 && i%5 == 4 && i%3 == 2 {
			fmt.Println(i)
			if i > 100000 {
				break
			}

		}
	}
}
