package tree

import (
	"container/heap"
	"fmt"
)

type HuffmanNode struct {
	char  rune
	freq  int
	left  *HuffmanNode
	right *HuffmanNode
}

type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*HuffmanNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func buildHuffmanTree(freqMap map[rune]int) *HuffmanNode {
	pq := make(PriorityQueue, 0)
	for char, freq := range freqMap {
		pq = append(pq, &HuffmanNode{char: char, freq: freq})
	}
	heap.Init(&pq)

	for len(pq) > 1 {
		left := heap.Pop(&pq).(*HuffmanNode)
		right := heap.Pop(&pq).(*HuffmanNode)
		internalNode := &HuffmanNode{char: 0, freq: left.freq + right.freq, left: left, right: right}
		heap.Push(&pq, internalNode)
	}

	return pq[0]
}

func printHuffmanCodes(node *HuffmanNode, code string) {
	if node.left == nil && node.right == nil {
		fmt.Printf("Character: %c, Frequency: %d, Code: %s\n", node.char, node.freq, code)
		return
	}
	if node.left != nil {
		printHuffmanCodes(node.left, code+"0")
	}
	if node.right != nil {
		printHuffmanCodes(node.right, code+"1")
	}
}

// func main() {
// 	text := "hello huffman"
// 	freqMap := make(map[rune]int)
// 	for _, char := range text {
// 		freqMap[char]++
// 	}

// 	root := buildHuffmanTree(freqMap)
// 	printHuffmanCodes(root, "")

// }
