package tree

import (
	"fmt"
)

// 定义B+树节点的结构
type Node struct {
	leaf     bool
	keys     []int
	children []*Node
}

// 初始化一个新的节点
func NewNode(leaf bool) *Node {
	return &Node{
		leaf:     leaf,
		keys:     []int{},
		children: []*Node{},
	}
}

// 在节点中插入键
func (n *Node) insert(key int) {
	// 找到要插入的位置
	i := len(n.keys) - 1
	for i >= 0 && n.keys[i] > key {
		i--
	}
	i++

	// 插入键
	newKeys := make([]int, len(n.keys)+1)
	copy(newKeys, n.keys[:i])
	newKeys[i] = key
	copy(newKeys[i+1:], n.keys[i:])
	n.keys = newKeys
}

// 定义B+树结构
type BPlusTree struct {
	root *Node
}

// 初始化一个新的B+树
func NewBPlusTree() *BPlusTree {
	return &BPlusTree{
		root: NewNode(true),
	}
}

// 在B+树中插入键
func (b *BPlusTree) Insert(key int) {
	// 如果根节点已满，则拆分根节点
	if len(b.root.keys) >= 3 {
		newRoot := NewNode(false)
		newRoot.children = append(newRoot.children, b.root)
		b.splitChild(newRoot, 0)
		b.root = newRoot
	}
	b.insertNonFull(b.root, key)
}

// 在非满节点中插入键
func (b *BPlusTree) insertNonFull(node *Node, key int) {
	// 如果是叶子节点，则直接插入键
	if node.leaf {
		node.insert(key)
		return
	}

	// 否则，找到合适的子节点递归插入
	i := len(node.keys) - 1
	for i >= 0 && node.keys[i] > key {
		i--
	}
	i++

	// 检查子节点是否已满
	if len(node.children[i].keys) >= 3 {
		b.splitChild(node, i)
		if node.keys[i] < key {
			i++
		}
	}
	b.insertNonFull(node.children[i], key)
}

// 拆分节点
func (b *BPlusTree) splitChild(parent *Node, index int) {
	// 获取要拆分的节点和其子节点
	nodeToSplit := parent.children[index]
	newChild := NewNode(nodeToSplit.leaf)

	// 将后半部分键和子节点移至新节点
	newChild.keys = append(newChild.keys, nodeToSplit.keys[2:]...)
	nodeToSplit.keys = nodeToSplit.keys[:2]
	if !nodeToSplit.leaf {
		newChild.children = append(newChild.children, nodeToSplit.children[2:]...)
		nodeToSplit.children = nodeToSplit.children[:2]
	}

	// 插入新节点到父节点
	parent.insert(nodeToSplit.keys[1])
	parent.children = append(parent.children, nil)
	copy(parent.children[index+2:], parent.children[index+1:])
	parent.children[index+1] = newChild
}

// 打印B+树
func (b *BPlusTree) Print() {
	printNode(b.root, 0)
}

// 打印节点
func printNode(node *Node, level int) {
	if node == nil {
		return
	}
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "\t"
	}
	fmt.Printf("%s%v\n", prefix, node.keys)
	for _, child := range node.children {
		printNode(child, level+1)
	}
}

func main() {
	bpt := NewBPlusTree()
	keys := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	for _, key := range keys {
		bpt.Insert(key)
	}
	bpt.Print()
}
