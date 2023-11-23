package main

import "fmt"

type DoubleNode struct {
	Pre  *DoubleNode
	Next *DoubleNode
	Data int
}

type DoubleLinkedList struct {
	head *DoubleNode
	tail *DoubleNode
}

// 在末尾添加
func (l *DoubleLinkedList) AppendToTail(data int) {
	newNode := &DoubleNode{Data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.Pre = l.tail
	l.tail.Next = newNode
	l.tail = newNode
}

func (l *DoubleLinkedList) AppendToHeader(data int) {
	newNode := &DoubleNode{Data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.Pre = l.head
	l.head.Next.Pre = newNode
	newNode.Next = l.head.Next
	l.head.Next = newNode
}

func (l *DoubleLinkedList) DeleteToTail(data any) {
	if l.head == nil {
		return
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	}
	l.tail = l.tail.Pre
	l.tail.Next.Pre = nil
	l.tail.Next = nil
}

func (l *DoubleLinkedList) DisplayList() {
	if l.head == nil {
		return
	}

	tmp := l.head
	for tmp != nil {
		fmt.Println(tmp.Data)
		tmp = tmp.Next
	}
}

// func main() {
// 	dl := &DoubleLinkedList{}

// 	datas := []int{1, 2, 3, 4, 5, 6, 7, 8}

// 	for _, data := range datas {
// 		dl.AppendToTail(data)
// 	}

// 	dl.DisplayList()
// }
