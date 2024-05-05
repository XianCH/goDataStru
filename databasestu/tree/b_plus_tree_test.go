package tree

import (
	"testing"
)

func TestBPlusTree_Insert(t *testing.T) {
	tree := NewBPlusTree()

	// Insert keys into the tree
	keys := []int{5, 3, 8, 1, 4, 6, 9}
	for _, key := range keys {
		tree.Insert(key)
	}

	// Define expected tree structure after insertion
	expectedKeys := [][]int{
		{4},
		{1, 3, 5},
		{6, 8},
		{9},
	}
	expectedLeaf := []bool{true, true, true, true}

	// Validate the tree structure
	validateTree(t, tree.root, expectedKeys, expectedLeaf)
}

func validateTree(t *testing.T, node *Node, expectedKeys [][]int, expectedLeaf []bool) {
	if node == nil {
		return
	}

	// Validate keys in the node
	expected := expectedKeys[0]
	expectedKeys = expectedKeys[1:]
	if !equal(node.keys, expected) {
		t.Errorf("Expected keys %v, got %v", expected, node.keys)
	}

	// Validate leaf status
	expectedLeafStatus := expectedLeaf[0]
	expectedLeaf = expectedLeaf[1:]
	if node.leaf != expectedLeafStatus {
		t.Errorf("Expected leaf status %v, got %v", expectedLeafStatus, node.leaf)
	}

	// Recursively validate children
	for _, child := range node.children {
		validateTree(t, child, expectedKeys, expectedLeaf)
	}
}

func equal(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}
