package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func Walk(t *tree.Tree, ch chan int) {
	defer close(ch) // <- closes the channel when this function returns
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for k := range ch1 {
		select {
		case g := <-ch2:
			if k != g {
				return false
			}
		default:
			break
		}
	}
	return true
}
func main() {
	fmt.Println("Equivalent Binary trees?")
	t1 := tree.New(3)
	t1.Left = tree.New(1)
	t1.Right = tree.New(8)
	t1.Left.Left = tree.New(1)
	t1.Left.Right = tree.New(2)
	t1.Right.Left = tree.New(5)
	t1.Right.Right = tree.New(13)

	t2 := tree.New(8)
	t2.Left = tree.New(3)
	t2.Right = tree.New(13)
	t2.Left.Left = tree.New(1)
	t2.Left.Right = tree.New(5)
	t2.Left.Left.Left = tree.New(1)
	t2.Left.Left.Right = tree.New(2)

	if Same(t1, t2) {
		fmt.Println("Trees are Equivalent")
	} else {
		fmt.Println("Not equivalent")
	}

	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Equal")
	}
	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("Not Equal")
	}

}
