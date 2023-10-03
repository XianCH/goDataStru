package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) Append(data interface{}) {
	newNode := &Node{data: data, prev: nil, next: nil}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

func (dll *DoublyLinkedList) Display() {
	current := dll.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}


func main() {
	dll := DoublyLinkedList{}

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	fmt.Println("Doubly Linked List:")
	dll.Display()
}
