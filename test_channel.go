package main

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			println("quit")
			return
		}
	}
}
func main1() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func main() {
	var a string
	var c = make(chan int)

	go func() {
		a = "hello world"
		c <- 11
		// <-c
	}()
	a = "xhk"
	<-c
	// c <- 11
	println(a)

	// close(c)
	// c <- 1
	// cc := <-c
	// println(<-c)

}
