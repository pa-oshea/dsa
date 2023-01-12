package trees

import (
	"fmt"

	"github.com/pa-oshea/dsa/common"
	"github.com/pa-oshea/dsa/lists"
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

// Breadth first search
func (b *binary_tree) BFS(item int) bool {
	queue := lists.Queue[*common.Binary_node[int]]{}
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
