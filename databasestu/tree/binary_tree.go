package main

import (
	"fmt"
)

type node struct {
	key        int
	LeftChild  *node
	RightChild *node
}

type binaryTree struct {
	root *node
}

// 插入原始
func (bt *binaryTree) insert(key int) {
	if bt.root == nil {
		bt.root = &node{key: key}
	}
	bt.insertRecursively(bt.root, key)

}

func (bt *binaryTree) insertRecursively(root *node, key int) {
	if key < root.key {
		if root.LeftChild == nil {
			root.LeftChild = &node{key: key}
		} else {
			bt.insertRecursively(root.LeftChild, key)
		}
	} else {
		if root.RightChild == nil {
			root.RightChild = &node{key: key}
		} else {
			bt.insertRecursively(root.RightChild, key)
		}
	}
}

// 搜索元素
func (bt *binaryTree) search(key int) bool {
	return bt.searchRecurisively(bt.root, key)

}

func (bt *binaryTree) searchRecurisively(root *node, key int) bool {
	if root == nil {
		fmt.Println("root is nil")
		return false
	}
	if root.key == key {
		return true
	}
	if root.key > key {
		return bt.searchRecurisively(root.LeftChild, key)
	} else {
		return bt.searchRecurisively(root.RightChild, key)
	}

}

// 前序遍历
func (bt *binaryTree) PerOrderTraversal(node *node) {
	if node != nil {
		fmt.Printf("%d", node.key)
		bt.PerOrderTraversal(node.LeftChild)
		bt.PerOrderTraversal(node.RightChild)
	}
}

// 中序遍历
func (bt *binaryTree) InOrderTraversal(root *node) {
	if root != nil {
		bt.InOrderTraversal(root.LeftChild)
		fmt.Printf("%d", root.key)
		bt.InOrderTraversal(root.RightChild)
	}
}

// 后序遍历
func (bt *binaryTree) PostOrderTraversal(root *node) {
	if root != nil {
		bt.PostOrderTraversal(root.LeftChild)
		bt.PostOrderTraversal(root.RightChild)
		fmt.Printf("%d", root.key)
	}
}

/*
func main() {
	arrs := []int{4, 2, 5, 3, 1, 6}
	bt := binaryTree{}
	for _, arr := range arrs {
		bt.insert(arr)
	}

	//前序遍历
	bt.PerOrderTraversal(bt.root)
	fmt.Println()
	//中序遍历
	bt.InOrderTraversal(bt.root)
	fmt.Println()
	//后序遍历
	bt.PostOrderTraversal(bt.root)
	fmt.Println()
}

*/
