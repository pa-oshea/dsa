package trees

import (
	"fmt"

	"github.com/pa-oshea/dsa/common"
	"github.com/pa-oshea/dsa/queue"
)

type binary_tree struct {
	height int
	head   *common.Binary_node[int]
}

func walk_pre_order(curr *common.Binary_node[int], path *[]int) {
	if curr == nil {
		return
	}

	*path = append(*path, curr.Value)
	walk_pre_order(curr.Left, path)
	walk_pre_order(curr.Right, path)
}

func (b *binary_tree) get_pre_order() []int {
	path := make([]int, 0)
	walk_pre_order(b.head, &path)
	return path
}

func walk_in_order(curr *common.Binary_node[int], path *[]int) {
	if curr == nil {
		return
	}

	walk_in_order(curr.Left, path)
	*path = append(*path, curr.Value)
	walk_in_order(curr.Right, path)
}

func (b *binary_tree) get_in_order() []int {
	path := []int{}
	walk_in_order(b.head, &path)
	return path
}

func walk_post_order(curr *common.Binary_node[int], path *[]int) {
	if curr == nil {
		return
	}

	walk_post_order(curr.Left, path)
	walk_post_order(curr.Right, path)
	*path = append(*path, curr.Value)
}

func (b *binary_tree) get_post_order() []int {
	path := make([]int, 0)
	walk_post_order(b.head, &path)
	fmt.Printf("path: %v\n", path)
	return path
}

// BFT
func (b *binary_tree) breadth_first_traveral() (r []int) {
	queue := queue.Queue[*common.Binary_node[int]]{}
	queue.Enqueue(b.head)

	if queue.Length == 0 {
		return
	}

	for queue.Length > 0 {
		curr, _ := queue.Dequeue()
		if curr == nil {
			continue
		}

		r = append(r, curr.Value)

		if curr.Left != nil {
			queue.Enqueue(curr.Left)
		}

		if curr.Right != nil {
			queue.Enqueue(curr.Right)
		}
	}

	return
}

// Breadth first search
func (b *binary_tree) BFS(item int) bool {
	queue := queue.Queue[*common.Binary_node[int]]{}
	queue.Enqueue(b.head)

	if queue.Length == 0 {
		return false
	}

	for queue.Length > 0 {
		curr, _ := queue.Dequeue()
		if curr == nil {
			continue
		}

		if curr.Value == item {
			return true
		}

		if curr.Left != nil {
			queue.Enqueue(curr.Left)
		}

		if curr.Right != nil {
			queue.Enqueue(curr.Right)
		}
	}

	return false
}

func invert_tree(curr *common.Binary_node[int]) {
	if curr == nil {
		return
	}

	curr.Right, curr.Left = curr.Left, curr.Right
	invert_tree(curr.Left)
	invert_tree(curr.Right)
}

func (b *binary_tree) invert() {
	invert_tree(b.head)
}

func longestPath(root *common.Binary_node[int], diameter *int) int {
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

func (b *binary_tree) findDiameter() int {
	diameter := 0
	longestPath(b.head, &diameter)
	return diameter
}
