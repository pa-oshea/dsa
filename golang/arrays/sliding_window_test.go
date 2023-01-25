package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumSubArr(t *testing.T) {
	assert.Equal(t, 16, maxSumSubArr([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 3))
	assert.Equal(t, 10, maxSumSubArr([]int{5, 5, 2, 3, 4, 2, 1}, 2))
}

func TestSmallestSubArray(t *testing.T) {
	assert.Equal(t, 1, smallestSubArray([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 8))
	assert.Equal(t, 2, smallestSubArray([]int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}, 15))
}

func TestLongestSubStringNDistinct(t *testing.T) {
	assert.Equal(t, 5, longestSubStringNDistinct("aaahhibc", 2))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 3, lengthOfLongestSubstring("dvdf"))
}

func TestLengthOfLongestSubstringOptimised(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstringOptimised("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstringOptimised("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstringOptimised("pwwkew"))
	assert.Equal(t, 3, lengthOfLongestSubstringOptimised("dvdf"))
}

func TestHouseRobber(t *testing.T) {
	assert.Equal(t, 10, houseRobberRecursive([]int{3, 1, 2, 5, 4, 2}))
	assert.Equal(t, 10, houseRobber([]int{3, 1, 2, 5, 4, 2}))
	assert.Equal(t, 10, houseRobberAlt([]int{3, 1, 2, 5, 4, 2}))
}

func Test_MaxSlidingWindow(t *testing.T) {
	expected := []int{3, 3, 5, 5, 6, 7}
	input := []int{1, 3, -1, -3, 5, 3, 6, 7}
	assert.Equal(t, expected, maxSlidingWindow(input, 3))
	expected = []int{1,-1}
	input = []int{1,-1}
	assert.Equal(t, expected, maxSlidingWindow(input, 1))
}
