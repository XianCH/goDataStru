package tree

func heapSort(arr []int) {
	n := len(arr)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 依次取出堆顶元素，交换位置，重新堆化
	for i := n - 1; i > 0; i-- {
		// 交换堆顶和当前堆的最后一个元素
		arr[i], arr[0] = arr[0], arr[i]

		// 对交换后的堆顶执行堆化
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// 找出左子节点和右子节点中的最大值
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是父节点，交换父节点和最大值，然后递归堆化
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
