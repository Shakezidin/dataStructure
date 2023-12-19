package hello

import (
	"math"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func FindHeigt(root *Node) int {
	if root == nil {
		return -1
	}

	left := FindHeigt(root.Left)
	Right := FindHeigt(root.Right)

	return int(math.Max(float64(left), float64(Right)) + 1)
}

var flag = 1

func ISbalance(root *Node) bool {
	isBalancehelper(root)

	if flag != 1 {
		return false
	}
	return true
}

func isBalancehelper(root *Node) int {
	if root == nil {
		return -1
	}

	left := isBalancehelper(root.Left)
	right := isBalancehelper(root.Right)
	if left-right > 1 || right-left > 1 {
		flag = 0

	}

	return int(math.Max(float64(left), float64(right)) + 1)
}

func (t *Tree) Insert(value int) {
	newNode := Node{Data: value}

	if t.Root == nil {
		t.Root = &newNode
		return
	}

	current := t.Root
	for {
		if current.Data > value {
			if current.Left == nil {
				current.Left = &newNode
				return
			}
			current = current.Left
		} else if current.Data < value {
			if current.Right == nil {
				current.Right = &newNode
				return
			}
			current = current.Right
		}
	}
}

func (t *Tree) Contains(data int) bool {
	if t.Root == nil {
		return false
	}

	current := t.Root

	for current != nil {
		if data > current.Data {
			current = current.Right
		} else if data < current.Data {
			current = current.Left
		} else {
			return true
		}
	}
	return false
}

func checkIsBST(root *Node) bool {
	var arr []int

	Inorder(root, &arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func Inorder(root *Node, arr *[]int) {
	if root == nil {
		return
	}
	Inorder(root.Left, arr)
	*arr = append(*arr, root.Data)
	Inorder(root.Right, arr)
}
