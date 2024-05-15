package algorithmBasic

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
