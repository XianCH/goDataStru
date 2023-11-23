package main

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Append(data interface{}) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// func main() {
// 	ll := LinkedList{}

// 	ll.Append(1)
// 	ll.Append(2)
// 	ll.Append(3)

// 	fmt.Println("Single Linked List:")
// 	ll.Display()
// }
