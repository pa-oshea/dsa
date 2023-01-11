package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var expected = []int{3, 4, 7, 9, 42, 69, 420}
var unsorted_arr = []int{9, 3, 7, 4, 69, 420, 42}

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
