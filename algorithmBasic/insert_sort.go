package algorithmBasic

func InsertSort(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		for js := i; js > 0; js-- {
			if array[js] < array[js-1] {
				array[js], array[js-1] = array[js-1], array[js]
			}
		}
	}
}
