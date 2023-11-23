package main

import (
	"fmt"
)

type JNode struct {
	data int
	next *JNode
}

type JList struct {
	FirstNode *JNode
	k         int
}

func (jl *JList) addJNode(newNode *JNode) {
	if jl.FirstNode == nil {
		jl.FirstNode = newNode
		jl.FirstNode.next = jl.FirstNode
		return
	}

	currentNode := jl.FirstNode
	for currentNode.next != jl.FirstNode {
		currentNode = currentNode.next
	}

	currentNode.next = newNode
	newNode.next = jl.FirstNode
}

func (jl *JList) Display() {
	if jl.FirstNode == nil {
		println("链表为空")
		return
	}

	currentNode := jl.FirstNode

	for {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next

		if currentNode == jl.FirstNode {
			break // 循环到达链表末尾，退出循环
		}
	}
}

func (jl *JList) josephu(m int) []int {
	if jl.FirstNode == nil {
		fmt.Println("链表为空")
		return nil
	}
	var result []int
	currentNode := jl.FirstNode

	for jl.FirstNode.next != jl.FirstNode {
		// 移动到第 m-1 个节点
		for i := 0; i < m-1; i++ {
			currentNode = currentNode.next
		}

		// 将第 m 个节点的数据添加到结果数组中
		result = append(result, currentNode.next.data)

		// 删除第 m 个节点
		currentNode.next = currentNode.next.next
	}

	// 最后一个节点
	fmt.Println("最后一个节点", jl.FirstNode.data)
	result = append(result, jl.FirstNode.data)

	return result
}

func (jl *JList) countJosephuSize() int {
	if jl.FirstNode == nil {
		return 0
	}
	var size int
	current := jl.FirstNode
	for current != nil {
		size++
		current = current.next
		if current == jl.FirstNode {
			break
		}
	}
	return size
}

func main() {

	JList := &JList{
		k: 5,
	}

	arrayDatas := []int{1, 2, 3, 4, 5, 6, 6}
	for _, data := range arrayDatas {
		newNode := &JNode{data: data}
		JList.addJNode(newNode)
	}

	JList.Display()

	JList.josephu(5)

	// for _, i := range is {
	// 	fmt.Println(i)
	// }

}
