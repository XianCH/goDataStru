package main

import (
	"errors"
	"fmt"
)

type StackArray struct {
	item []interface{}
	top  int16
}

func (sa *StackArray) push(data interface{}) {
	sa.item = append(sa.item, data)
	sa.top++
}

func (sa *StackArray) Pop() (interface{}, error) {
	if sa.top < 0 {
		return nil, errors.New("stack is empty")
	}

	item := sa.item[sa.top]
	sa.item = sa.item[:sa.top]
	sa.top--

	return item, nil
}

func (sa *StackArray) Peek() (interface{}, error) {
	if sa.top < 0 {
		return nil, errors.New("stack is empty")
	}

	return sa.item[sa.top], nil
}

func main() {
	stack := StackArray{}

	stack.push(1)
	stack.push(2)
	stack.push(3)

	if top, err := stack.Peek(); err == nil {
		fmt.Println("Top element:", top)
	} else {
		fmt.Println("Error:", err)
	}

	if popped, err := stack.Pop(); err == nil {
		fmt.Println("Popped element:", popped)
	} else {
		fmt.Println("Error:", err)
	}

	if top, err := stack.Peek(); err == nil {
		fmt.Println("Top element after Pop:", top)
	} else {
		fmt.Println("Error:", err)
	}
}
