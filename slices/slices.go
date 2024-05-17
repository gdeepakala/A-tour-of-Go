package main

import (
	"fmt"
	"strings"
)

func printBoard(b [][]string, c int) {
	fmt.Println("Board after move:", c)
	for i := 0; i < len(b); i++ {
		//fmt.Println(b[i])
		fmt.Printf("%s\n", strings.Join(b[i], " "))
	}

}

func printUsingRange(b [][]string) {
	fmt.Println("Printing board using range:")
	for _, row := range b {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
func ticTacToe() {
	var count int
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	board[0][0] = "O"
	count++
	printBoard(board, count)
	board[1][0] = "X"
	count++
	printBoard(board, count)
	board[2][2] = "O"
	count++
	printBoard(board, count)
	board[1][1] = "X"
	count++
	printBoard(board, count)
	board[1][2] = "O"
	count++
	printBoard(board, count)
	board[0][2] = "X"
	count++
	printBoard(board, count)
	board[2][0] = "O"
	count++
	printBoard(board, count)
	board[2][1] = "X"
	count++
	printBoard(board, count)
	board[0][1] = "O"
	count++
	printBoard(board, count)
	printUsingRange(board)

}
func main() {
	mySlice := []int{1, 2, 3, 4, 5}

	mySlice = append(mySlice, 6)

	anotherSlice := mySlice[1:4]

	fmt.Println("mySlice : ", mySlice, " \t Len (mySlice):",
		len(mySlice), "\t Cap(mySlice)", cap(mySlice))
	fmt.Println("anotherSlice : ", anotherSlice, " \t Len (anotherSlice):",
		len(anotherSlice), "\t cap(anotherSlice)", cap(anotherSlice))

	s := []int{1, 2, 3, 4, 5, 6}
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	ticTacToe()
}
