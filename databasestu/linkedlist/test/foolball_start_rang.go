package main

import (
	"fmt"
)

type Star struct {
	Name     string
	No       int
	Position string
}

type Node struct {
	Data Star
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) AppendNode(node Node) {
	if ll.Head == nil {
		ll.Head = &node
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &node
}

func (ll *LinkedList) InsertNode(node Node, position int) {
	if position <= 0 {
		node.Next = ll.Head
		ll.Head = &node
		return
	}

	current := ll.Head
	for i := 1; i < position; i++ {
		if current == nil {
			fmt.Println("Position not found.")
			return
		}
		current = current.Next
	}

	node.Next = current.Next
	current.Next = &node
}

func (ll *LinkedList) DeleteNode(name string) {
	if ll.Head == nil {
		fmt.Println("List is empty.")
		return
	}

	if ll.Head.Data.Name == name {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for current.Next != nil && current.Next.Data.Name != name {
		current = current.Next
	}

	if current.Next != nil {
		current.Next = current.Next.Next
	} else {
		fmt.Println("Player not found.")
	}
}

func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("Name: %s, No: %d, Position: %s\n", current.Data.Name, current.Data.No, current.Data.Position)
		current = current.Next
	}
}

func (ll *LinkedList) Sort() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	// 使用插入排序来对链表进行排序
	sorted := &Node{Data: Star{}}
	current := ll.Head
	for current != nil {
		next := current.Next
		insert(sorted, current)
		current = next
	}

	ll.Head = sorted.Next
}

func insert(sorted, node *Node) {
	current := sorted
	for current.Next != nil && current.Next.Data.No < node.Data.No {
		current = current.Next
	}
	node.Next = current.Next
	current.Next = node
}

func main() {
	ll := LinkedList{}

	player1 := Node{Data: Star{Name: "Player1", No: 1, Position: "Forward"}}
	player2 := Node{Data: Star{Name: "Player2", No: 2, Position: "Midfielder"}}
	player3 := Node{Data: Star{Name: "Player3", No: 3, Position: "Goalkeeper"}}

	ll.AppendNode(player3)
	ll.AppendNode(player1)
	ll.AppendNode(player2)

	fmt.Println("Initial List:")
	ll.Display()

	player4 := Node{Data: Star{Name: "Player4", No: 4, Position: "Defender"}}
	ll.InsertNode(player4, 2)

	fmt.Println("\nList after inserting Player4 at position 2:")
	ll.Display()

	ll.DeleteNode("Player2")

	fmt.Println("\nList after deleting Player2:")
	ll.Display()

	fmt.Println("\nList after sorting by Player No:")
	ll.Sort()
	ll.Display()
}
