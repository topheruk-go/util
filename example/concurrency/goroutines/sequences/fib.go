package main

func fib(n int, c chan int) {
	defer close(c)

	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
}
