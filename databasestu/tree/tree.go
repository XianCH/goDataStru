package tree

type BaseNode struct {
	value     any
	LeftNode  *BaseTree
	RightNode *BaseTree
}

type BaseTree interface {
	insert() error
	findValue() (any, error)
	delete() (any, error)
	display()
	deep() int
	CountNode() int
	isEmpty() bool
}
