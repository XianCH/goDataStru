package algorithmBasic

import (
	"reflect"
	"testing"
)

func TestSelectSort(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{5, 1, 4, 2, 8}, []int{1, 2, 4, 5, 8}},
		{[]int{7, 3, 9, 2}, []int{2, 3, 7, 9}},
	}

	for _, tc := range testCases {
		arr := make([]int, len(tc.input))
		copy(arr, tc.input)

		selectSort(arr)

		if !reflect.DeepEqual(arr, tc.expected) {
			t.Errorf("SelectSort(%v) = %v, want %v", tc.input, arr, tc.expected)
		}
	}
}
