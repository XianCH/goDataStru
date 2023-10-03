package algorthmBasic

import (
    "testing"
)

func TestBinarySearch(t *testing.T) {
    arr := []int{1, 3, 5, 7, 9, 11, 13}
    tests := []struct {
        target   int
        expected int
    }{
        {3, 1},   // 3 is at index 1
        {13, 6},  // 13 is at index 6
        {6, -1},  // 6 is not in the array
        {0, -1},  // 0 is not in the array
    }

    for _, test := range tests {
        result := BinarySearch(arr, test.target)
        if result != test.expected {
            t.Errorf("BinarySearch(%v, %d) = %d; expected %d", arr, test.target, result, test.expected)
        }
    }
}
