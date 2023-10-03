package sort

func selectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i] // 交换元素
		}
	}
}

// func main() {
// 	arrs := []int{3, 5, 7, 2, 3, 8, 1}
// 	fmt.Println(arrs)
// 	selectSort(arrs)
// 	fmt.Println(arrs)
// }
