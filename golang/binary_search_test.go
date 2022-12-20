package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	assert.Equal(t, true, binary_search(foo, 69))
	assert.Equal(t, false, binary_search(foo, 1336))
	assert.Equal(t, true, binary_search(foo, 69420))
	assert.Equal(t, false, binary_search(foo, 69421))
	assert.Equal(t, true, binary_search(foo, 1))
	assert.Equal(t, false, binary_search(foo, 0))
}

func TestBinarySearchRecursive(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	assert.Equal(t, true, binary_search_recursive(foo, 69))
	assert.Equal(t, false, binary_search_recursive(foo, 1336))
	assert.Equal(t, true, binary_search_recursive(foo, 69420))
	assert.Equal(t, false, binary_search_recursive(foo, 69421))
	assert.Equal(t, true, binary_search_recursive(foo, 1))
	assert.Equal(t, false, binary_search_recursive(foo, 0))
}

func TestBinarySearchCorrect(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	assert.Equal(t, true, binary_search_correct(foo, 69))
	assert.Equal(t, false, binary_search_correct(foo, 1336))
	assert.Equal(t, true, binary_search_correct(foo, 69420))
	assert.Equal(t, false, binary_search_correct(foo, 69421))
	assert.Equal(t, true, binary_search_correct(foo, 1))
	assert.Equal(t, false, binary_search_correct(foo, 0))
}
