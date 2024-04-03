package trees

import (
	"github.com/pa-oshea/dsa/common"
	"github.com/pa-oshea/dsa/queue"
)

type BinaryTree[T comparable] struct {
	height int
	head   *common.Node[T]
}

func (b *BinaryTree[T]) GetInOrder() []T {
	var path []T
	walkInOrder(b.head, &path)
	return path
}

func (b *BinaryTree[T]) GetPreOrder() []T {
	var path []T
	walkPreOrder(b.head, &path)
	return path
}

func (b *BinaryTree[T]) GetPostOrder() []T {
	var path []T

	walkPostOrder(b.head, &path)
	return path
}

// Breadth First Search
// FIXME: this will panic if the struct is not comparable. Need a compare function parameter
func (b *BinaryTree[T]) BFS(item T) (T, bool) {
	var result T
	queue := queue.Queue[*common.Node[T]]{}
	queue.Enqueue(b.head)

	if queue.Length == 0 {
		return result, false
	}

	for queue.Length > 0 {
		curr, _ := queue.Dequeue()
		if curr == nil {
			continue
		}

		if curr.Value == item {
			result = curr.Value
			return result, true
		}

		if curr.Left != nil {
			queue.Enqueue(curr.Left)
		}

		if curr.Right != nil {
			queue.Enqueue(curr.Right)
		}
	}

	return result, false
}

// Breadth first traveral
func (b *BinaryTree[T]) BFT() []T {
	var result []T
	queue := queue.Queue[*common.Node[T]]{}
	queue.Enqueue(b.head)

	if queue.Length == 0 {
		return result
	}

	for queue.Length > 0 {
		curr, _ := queue.Dequeue()
		if curr == nil {
			continue
		}

		result = append(result, curr.Value)

		if curr.Left != nil {
			queue.Enqueue(curr.Left)
		}

		if curr.Right != nil {
			queue.Enqueue(curr.Right)
		}
	}

	return result
}

// TODO: make this it's own function ?
func (b *BinaryTree[T]) InvertTree() {
	invertTree(b.head)
}

func (b *BinaryTree[T]) Diameter() int {
	diameter := 0
	longestPath(b.head, &diameter)
	return diameter
}

func walkInOrder[T any](curr *common.Node[T], path *[]T) {
	if curr == nil {
		return
	}

	walkInOrder(curr.Left, path)
	*path = append(*path, curr.Value)
	walkInOrder(curr.Right, path)
}

func walkPreOrder[T any](curr *common.Node[T], path *[]T) {
	if curr == nil {
		return
	}

	*path = append(*path, curr.Value)
	walkPreOrder(curr.Left, path)
	walkPreOrder(curr.Right, path)
}

func walkPostOrder[T any](curr *common.Node[T], path *[]T) {
	if curr == nil {
		return
	}

	walkPostOrder(curr.Left, path)
	walkPostOrder(curr.Right, path)
	*path = append(*path, curr.Value)
}

func invertTree[T any](curr *common.Node[T]) {
	if curr == nil {
		return
	}

	curr.Right, curr.Left = curr.Left, curr.Right
	invertTree(curr.Left)
	invertTree(curr.Right)
}

func longestPath[T any](root *common.Node[T], diameter *int) int {
	if root == nil {
		return 0
	}

	leftPath := longestPath(root.Left, diameter)
	rightPath := longestPath(root.Right, diameter)

	if *diameter < leftPath+rightPath {
		*diameter = leftPath + rightPath
	}

	if leftPath > rightPath {
		return leftPath + 1
	}
	return rightPath + 1
}
