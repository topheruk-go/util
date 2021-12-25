package main

func char(s string, c chan string) {
	defer close(c)
	for i := 0; i < len(s); i++ {

		c <- string(s[i])
	}
}
