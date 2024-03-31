package trees

import (
	"testing"

	"github.com/pa-oshea/dsa/common"
	"github.com/stretchr/testify/assert"
)

var tree = common.Binary_node[int]{
	Value: 20,
	Left: &common.Binary_node[int]{
		Value: 10,
		Left: &common.Binary_node[int]{
			Value: 5,
			Right: &common.Binary_node[int]{
				Value: 7,
			},
		},
		Right: &common.Binary_node[int]{Value: 15},
	},
	Right: &common.Binary_node[int]{
		Value: 50,
		Left: &common.Binary_node[int]{
			Value: 30,
			Left: &common.Binary_node[int]{
				Value: 29,
			},
			Right: &common.Binary_node[int]{
				Value: 45,
			},
		},
		Right: &common.Binary_node[int]{Value: 100},
	},
}

// Test in, pre, and post order traversal on binary tree
func Test_BinaryTreeDFT(t *testing.T) {
	binaryTree := binary_tree{head: &tree}

	preOrder := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	assert.Equal(t, preOrder, binaryTree.get_pre_order())

	inOrder := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	assert.Equal(t, inOrder, binaryTree.get_in_order())

	postOrder := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	assert.Equal(t, postOrder, binaryTree.get_post_order())
}

// Test Breadth first search in binary tree
func Test_BinaryTreeBFS(t *testing.T) {
	binaryTree := binary_tree{head: &tree}

	expectedOrder := []int{20, 10, 50, 5, 15, 30, 100, 7, 29, 45}
	assert.Equal(t, expectedOrder, binaryTree.breadth_first_traveral())

	assert.Equal(t, true, binaryTree.BFS(45))
	assert.Equal(t, true, binaryTree.BFS(7))
	assert.Equal(t, false, binaryTree.BFS(69))
}

func Test_BinaryTreeInvert(t *testing.T) {
	binaryTree := binary_tree{head: &tree}
	binaryTree.invert()

	expected := []int{20, 50, 10, 100, 30, 15, 5, 45, 29, 7}
	assert.Equal(t, expected, binaryTree.breadth_first_traveral())
}
