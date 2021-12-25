package main

func sum(n int, c chan int) {
	defer close(c)

	y := 0
	for i := 0; i < n; i++ {
		y += i
		c <- y
	}
}
