package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	assert.Equal(t, []int{3, 4, 7, 9, 42, 69, 420}, insertion_sort([]int{9, 3, 7, 4, 69, 420, 42}))
}
