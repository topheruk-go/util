package main

func char(s string, c chan byte) {
	for i := 0; i < len(s); i++ {
		c <- s[i]
	}
}
