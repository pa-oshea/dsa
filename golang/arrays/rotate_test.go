package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateTemp(t *testing.T) {
	ans := []int{3, 4, 5, 6, 7, 1, 2}
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	result := rotate_with_temp(arr, 2)

	assert.Equal(t, ans, result)
}

func TestRotateTempStart(t *testing.T) {
	ans := []int{2, 3, 4, 5, 6, 7, 1}
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	result := rotate_with_temp(arr, 1)

	assert.Equal(t, ans, result)
}

func TestRotateTempEnd(t *testing.T) {
	ans := []int{1, 2, 3, 4, 5, 6, 7}
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	result := rotate_with_temp(arr, len(arr))

	assert.Equal(t, ans, result)
}

func TestRotateAppend(t *testing.T) {
	ans := []int{3, 4, 5, 6, 7, 1, 2}
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	result := rotate_with_append(arr, 2)

	assert.Equal(t, ans, result)
}

func TestRotateAppendStart(t *testing.T) {
	ans := []int{2, 3, 4, 5, 6, 7, 1}
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	result := rotate_with_append(arr, 1)

	assert.Equal(t, ans, result)
}

func TestRotateOneByOneLoop(t *testing.T) {
	ans := []int{3, 4, 5, 6, 7, 1, 2}
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	result := rotate_one_by_one_loop(arr, 2)
	assert.Equal(t, ans, result)
}

// func TestRotateDivide(t *testing.T) {
// 	ans := []int{3, 4, 5, 6, 7, 1, 2}
// 	arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	result := rotate_with_gcd(arr, len(arr), 2)
// 	assert.Equal(t, ans, result)
// }
