package tree

type Tree interface {
	insert() any
	delete() any
	isEmpty() bool
	Deep() int
	getValue(any) any
}

type BaseNode struct {
	value any
	left  *BaseNode
	right *BaseNode
}
