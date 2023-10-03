package main

import "fmt"

type SparseElement struct {
	row, col, value int
}

type SparseArray struct {
	row, col int
	data     []SparseElement
}

func NewSparseArray(row, col int) *SparseArray {
	return &SparseArray{
		row:  row,
		col:  col,
		data: []SparseElement{},
	}
}

func (sa *SparseArray) Set(row, col, value int) {
	element := &SparseElement{
		row:   row,
		col:   col,
		value: value,
	}
	sa.data = append(sa.data, *element)
}

func (sa *SparseArray) Get(row, col int) int {
	for _, data := range sa.data {
		if row == data.row && col == data.col {
			return data.value
		}
	}
	return 0
}

func (sa *SparseArray) printSparse() {
	fmt.Println("num row col value")
	i := 0
	for _, elemt := range sa.data {
		fmt.Printf("%d  %d  %d  %d", i, elemt.row, elemt.col, elemt.value)
		fmt.Println()
		i++
	}
}

func main() {
	sparseArray := NewSparseArray(10, 10)
	sparseArray.Set(1, 3, 1)
	sparseArray.Set(4, 2, 2)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d  ", sparseArray.Get(i, j))
		}
		fmt.Println()
	}

	sparseArray.printSparse()
}
