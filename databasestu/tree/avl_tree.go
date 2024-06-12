package tree

type AVLNode struct {
	key    int
	height int
	left   *AVLNode
	right  *AVLNode
}

func height(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func getBalance(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}

func rightRotate(y *AVLNode) *AVLNode {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x
}

func leftRotate(x *AVLNode) *AVLNode {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y
}

func insert(node *AVLNode, key int) *AVLNode {
	if node == nil {
		return &AVLNode{key: key, height: 1}
	}

	if key < node.key {
		node.left = insert(node.left, key)
	} else if key > node.key {
		node.right = insert(node.right, key)
	} else {
		return node // 处理重复键
	}

	node.height = 1 + max(height(node.left), height(node.right))

	balance := getBalance(node)

	// 左左情况
	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	// 右右情况
	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	// 左右情况
	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// 右左情况
	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func deleteNode(root *AVLNode, key int) *AVLNode {
	if root == nil {
		return root
	}

	if key < root.key {
		root.left = deleteNode(root.left, key)
	} else if key > root.key {
		root.right = deleteNode(root.right, key)
	} else {
		if root.left == nil || root.right == nil {
			var temp *AVLNode
			if root.left != nil {
				temp = root.left
			} else {
				temp = root.right
			}

			if temp == nil {
				temp = root
				root = nil
			} else {
				*root = *temp
			}
		} else {
			temp := minValueNode(root.right)
			root.key = temp.key
			root.right = deleteNode(root.right, temp.key)
		}
	}

	if root == nil {
		return root
	}

	root.height = 1 + max(height(root.left), height(root.right))

	balance := getBalance(root)

	// 左左情况
	if balance > 1 && getBalance(root.left) >= 0 {
		return rightRotate(root)
	}

	// 左右情况
	if balance > 1 && getBalance(root.left) < 0 {
		root.left = leftRotate(root.left)
		return rightRotate(root)
	}

	// 右右情况
	if balance < -1 && getBalance(root.right) <= 0 {
		return leftRotate(root)
	}

	// 右左情况
	if balance < -1 && getBalance(root.right) > 0 {
		root.right = rightRotate(root.right)
		return leftRotate(root)
	}

	return root
}

func minValueNode(node *AVLNode) *AVLNode {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
