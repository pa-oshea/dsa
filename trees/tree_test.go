package trees

import (
	"testing"

	"github.com/pa-oshea/dsa/common"
	"github.com/stretchr/testify/assert"
)

var tree = common.Node[int]{
	Value: 20,
	Left: &common.Node[int]{
		Value: 10,
		Left: &common.Node[int]{
			Value: 5,
			Right: &common.Node[int]{
				Value: 7,
			},
		},
		Right: &common.Node[int]{Value: 15},
	},
	Right: &common.Node[int]{
		Value: 50,
		Left: &common.Node[int]{
			Value: 30,
			Left: &common.Node[int]{
				Value: 29,
			},
			Right: &common.Node[int]{
				Value: 45,
			},
		},
		Right: &common.Node[int]{Value: 100},
	},
}

// Test in, pre, and post order traversal on binary tree
func TestBinaryTreeDFT(t *testing.T) {
	binaryTree := BinaryTree[int]{head: &tree}

	preOrder := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	assert.Equal(t, preOrder, binaryTree.GetPreOrder())

	inOrder := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	assert.Equal(t, inOrder, binaryTree.GetInOrder())

	postOrder := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	assert.Equal(t, postOrder, binaryTree.GetPostOrder())
}

func TestBreathFirstTraversal(t *testing.T) {
	binaryTree := BinaryTree[int]{head: &tree}
	expectedOrder := []int{20, 10, 50, 5, 15, 30, 100, 7, 29, 45}
	assert.Equal(t, expectedOrder, binaryTree.BFT())
}

// Test Breadth first search in binary tree
func TestBinaryTreeBFS(t *testing.T) {
	tests := []struct {
		value int
		found bool
	}{
		{value: 45, found: true},
		{value: 7, found: true},
		{value: 69, found: false},
	}
	binaryTree := BinaryTree[int]{head: &tree}

	for _, test := range tests {
		_, found := binaryTree.BFS(test.value)
		assert.Equal(t, test.found, found)
	}
}

func TestBinaryTreeInvert(t *testing.T) {
	binaryTree := BinaryTree[int]{head: &tree}
	binaryTree.InvertTree()

	expected := []int{20, 50, 10, 100, 30, 15, 5, 45, 29, 7}
	assert.Equal(t, expected, binaryTree.BFT())
}
