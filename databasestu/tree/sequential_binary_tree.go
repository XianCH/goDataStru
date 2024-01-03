package tree

import "fmt"

type SequentialBinaryTree struct {
	Nodes []interface{}
}

func NewSequentialBinaryTree(size int) *SequentialBinaryTree {
	return &SequentialBinaryTree{
		Nodes: make([]interface{}, size),
	}
}

func (sbt *SequentialBinaryTree) Insert(data interface{}, index int) {
	if index >= len(sbt.Nodes) {
		fmt.Println("Full")
		return
	}
	sbt.Nodes[index] = data
}

func (sbt *SequentialBinaryTree) PreorderTraversal(index int) {
	if index >= len(sbt.Nodes) {
		return
	}
	if sbt.Nodes[index] != nil {
		fmt.Printf("%v ", sbt.Nodes[index])
		sbt.PreorderTraversal(2*index + 1) // Left child
		sbt.PreorderTraversal(2*index + 2) // Right child
	}
}

// func main() {
// 	tree := NewSequentialBinaryTree(10)
//
// 	// Insert nodes
// 	tree.Insert(10, 0)
// 	tree.Insert(2, 1)
// 	tree.Insert(3, 2)
// 	tree.Insert(4, 3)
// 	tree.Insert(5, 4)
// 	tree.Insert(6, 5)
//
// 	// Perform Preorder Traversal and print the tree
// 	fmt.Println("Preorder Traversal:")
// 	tree.PreorderTraversal(0)
// }
