package main

import (
	"sync"
)

type Queen struct {
	Size  int
	items []interface{}
	mutex sync.Mutex
}

func (queen *Queen) SetMaxSize(size int) {
	queen.Size = size
}

func (queen *Queen) IsFull() bool {
	return len(queen.items) == queen.Size
}

func (queen *Queen) EnQueen(item interface{}) {
	queen.mutex.Lock()
	defer queen.mutex.Unlock()
	if !queen.IsFull() {
		queen.items = append(queen.items, item)
	}
}

func (queen *Queen) Fornt() interface{} {
	return queen.items[0]
}

func (queen *Queen) IsEmpty() bool {
	return len(queen.items) == 0
}

func main() {

	queen := &Queen{
		Size: 20,
	}

	go func() {
		for i := 0; i < 20; i++ {

		}
	}()
}
