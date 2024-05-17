package main

import (
	"errors"
	"fmt"
)

type linkNode[T any] struct {
	val  T
	next *linkNode[T]
}

type singlyLinkedList[T any] struct {
	head *linkNode[T]
	len  int
}

func makeNewNode[T any](value T) *linkNode[T] {
	return &linkNode[T]{
		val:  value,
		next: nil,
	}

}

func newSinglyLinkedList[T any]() *singlyLinkedList[T] {

	return &singlyLinkedList[T]{
		head: nil,
		len:  0,
	}
}

func (myList *singlyLinkedList[T]) insertAtHead(value T) {
	newNode := makeNewNode(value)
	newNode.next = myList.head
	myList.head = newNode
	myList.len++
}

func (myList *singlyLinkedList[T]) insertAtTail(value T) {
	newNode := makeNewNode(value)
	if myList.isEmpty() {
		fmt.Println("EMpty Linked List. Calling insertAtHead() ")
		myList.insertAtHead(value)
	} else {
		temp := myList.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
}

func (myList *singlyLinkedList[T]) deleteAtHead() {
	if myList.isEmpty() {
		fmt.Println("Empty List! Can not delete")
		return
	} else {
		myList.head = myList.head.next
	}
}
func (myList *singlyLinkedList[T]) isEmpty() bool {
	if myList.head == nil {
		return true
	} else {
		return false
	}
}
func (myList *singlyLinkedList[T]) deleteAtIndex(index int) {
	if myList.isEmpty() {
		fmt.Println("Empty List! Can not delete")
		return
	} else if index == 0 {
		myList.deleteAtHead()
	} else if index > myList.getLength() {
		fmt.Println(index, " > List length (", myList.getLength(), "). Can not delete")
	} else {
		temp := myList.head
		var prev *linkNode[T] = temp

		for i := 0; i < index && temp.next != nil; i++ {
			prev = temp
			temp = temp.next
		}
		prev.next = temp.next

	}

}

func (myList *singlyLinkedList[T]) searchAtIndex(index int) (T, error) {
	var ret T
	if myList.isEmpty() {
		fmt.Println("Empty List! Nothing to Search")
		return ret, errors.New("nothing to searchs")
	} else if index > myList.getLength() {
		fmt.Println(index, " > List length (", myList.getLength(), "). Can not search")
		return ret, errors.New("can not search")
	} else {
		temp := myList.head
		for i := 0; i < index && temp.next != nil; i++ {
			temp = temp.next
		}
		ret = temp.val
	}
	return ret, nil
}

func (myList *singlyLinkedList[T]) getLength() int {
	return myList.len
}
func (myList *singlyLinkedList[T]) deleteAtTail() {
	if myList.isEmpty() {
		fmt.Println("Empty List! Can not delete")
		return
	} else {
		temp := myList.head
		var prev *linkNode[T] = temp
		for temp.next != nil {
			prev = temp
			temp = temp.next
		}
		prev.next = nil
	}

}

func (myList *singlyLinkedList[T]) getValues() []T {
	temp := myList.head
	var values []T
	for temp != nil {
		values = append(values, temp.val)
		temp = temp.next
	}
	return values
}

func main() {

	fmt.Println("Linked List Implementation:")
	var head singlyLinkedList[int]
	head.insertAtHead(3)
	head.insertAtHead(1)

	head.insertAtTail(5)
	head.insertAtTail(7)
	head.insertAtTail(9)
	head.insertAtTail(11)
	fmt.Print("Value at index 0=")
	result, err := head.searchAtIndex(0)
	if err == nil {
		fmt.Println(result)
	}
	fmt.Print("Value at index 1=")
	result, err = head.searchAtIndex(1)
	if err == nil {
		fmt.Println(result)
	}
	fmt.Print("Value at index 2=")
	result, err = head.searchAtIndex(2)
	if err == nil {
		fmt.Println(result)
	}
	fmt.Print("Value at index 7=")
	result, err = head.searchAtIndex(7)
	if err == nil {
		fmt.Println(result)
	}
	fmt.Println("Complete Linked list:")
	fmt.Println(head.getValues())
	head.deleteAtHead()
	fmt.Println(head.getValues())
	head.deleteAtTail()
	fmt.Println(head.getValues())
	head.deleteAtIndex(0)
	fmt.Println(head.getValues())
	head.deleteAtIndex(1)
	fmt.Println(head.getValues())

}
