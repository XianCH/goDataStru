package sort

import "fmt"

func main() {
	arrs := []int{9, 3, 5, 7, 2, 1, 6, 4, 5}
	bubbleSort(arrs)
	printBubble(arrs)
}

func bubbleSort(arrs []int) {
	len := len(arrs)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arrs[j] > arrs[j+1] {
				arrs[j], arrs[j+1] = arrs[j+1], arrs[j]
			}
		}
	}
}

func printBubble(arrs []int) {
	for _, arr := range arrs {
		fmt.Println(arr)
	}
}
