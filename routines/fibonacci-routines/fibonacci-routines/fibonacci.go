package main

import "fmt"

func fibonacci(c chan int, n int) {
	n1 := 0
	n2 := 1
	for i := 0; i < n; i++ {
		c <- n1
		n1, n2 = n2, n1+n2
	}
	close(c)
}
func main() {
	c := make(chan int, 10)
	go fibonacci(c, cap(c))
	for n := range c {
		fmt.Println(n)
	}
}
