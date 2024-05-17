package main

import "fmt"

func fibonacci() func() int {
	f := 0
	next := 1
	return func() int {
		res := f + next
		f = next
		next = res
		return res
	}
}

func main() {
	fn := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}
