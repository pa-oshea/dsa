package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expected     = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	unsorted_arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
)

func Test_SelectionSort(t *testing.T) {
	arr := make([]int, len(unsorted_arr))
	copy(arr, unsorted_arr)
	assert.Equal(t, expected, selection_sort(arr))
}

func Test_InsertionSort(t *testing.T) {
	arr := make([]int, len(unsorted_arr))
	copy(arr, unsorted_arr)
	assert.Equal(t, expected, insertion_sort(arr))
}

func Test_BubbleSort(t *testing.T) {
	arr := make([]int, len(unsorted_arr))
	copy(arr, unsorted_arr)
	assert.Equal(t, expected, bubble_sort(arr))
}

func Test_QuickSort(t *testing.T) {
	arr := make([]int, len(unsorted_arr))
	copy(arr, unsorted_arr)
	assert.Equal(t, expected, quick_sort(arr))
}
