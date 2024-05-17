package main

import "fmt"

func main() {
	i, j := 1, 4.5

	p1 := &i
	fmt.Println("i=", i, "*p1=", *p1)
	*p1++
	fmt.Println("=", i, "*p1+1=", *p1)

	p2 := &j
	fmt.Println("j=", j, "*p2=", *p2)
	*p2++
	fmt.Println("j=", j, "*p2+1=", *p2)

}
