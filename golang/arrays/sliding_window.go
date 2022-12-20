package arrays

import (
	"math"
)

/*
Fixed size sliding window.

Find the maximum sum of a contiguous subarray of size 3.

In sub arrays of 3 (size), find the maximun sum of that size.
*/
func maxSumSubArr(nums []int, size int) int {
	currentSum := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		if i >= size-1 {
			max = int(math.Max(float64(max), float64(currentSum)))
			currentSum -= nums[i-(size-1)]
		}
	}
	return max
}

/*
Dynamic size sliding window

Find the size of the smallest subarray >= to sum of N
*/
func smallestSubArray(nums []int, n int) int {
	minWindowSize := math.MaxInt
	currentWindowSum := 0
	windowStart := 0
	for windowEnd := 0; windowEnd < len(nums); windowEnd++ {
		currentWindowSum += nums[windowEnd]

		for currentWindowSum >= n {
			currentWindowSum -= nums[windowStart]
			minWindowSize = int(math.Min(float64(minWindowSize), float64(windowEnd)-float64(windowStart)+1))
			windowStart++
		}
	}
	return minWindowSize
}

/*
Dynamic sliding window with ADT (Abstract data type)

Find the longest substring with N Distinct characters
*/
func longestSubStringNDistinct(str string, n int) int {
	charFequencyMap := map[string]int{}
	windowStart, maxLength := 0, 0
	for windowEnd, char := range str {
		charFequencyMap[string(char)]++

		if len(charFequencyMap) > n {
			temp := charFequencyMap[string(str[windowStart])]
			delete(charFequencyMap, string(str[windowStart]))
			windowStart += temp
		}
		maxLength = int(math.Max(float64(maxLength), float64(windowEnd)-float64(windowStart)+1))
	}
	return maxLength
}

/*
Dynamic sliding window with ADT

Find the longest sub string with non repeating characters
leetcode 3
*/
func lengthOfLongestSubstring(s string) (r int) {
	charMap := map[byte]int{}
	windowStart := 0
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++

		for charMap[s[i]] > 1 {
			charMap[s[windowStart]] -= 1
			windowStart++
		}
		r = int(math.Max(float64(r), float64(i)-float64(windowStart)+1))
	}
	return
}

/*
Same as above just optimised
TODO: understand this better
*/
func lengthOfLongestSubstringOptimised(s string) (r int) {
	n := len(s)
	charMap := map[string]int{}
	for i, j := 0, 0; j < n; j++ {
		if _, ok := charMap[string(s[j])]; ok {
			i = int(math.Max(float64(charMap[string(s[j])]), float64(i)))
		}
		r = int(math.Max(float64(r), float64(j-i+1)))
		charMap[string(s[j])] = j + 1
	}
	return
}

/*
House robber problem
Rob the most gold from the houses, cannot rob houses next to each other
Meaning there must be at least one house between any two that you rob
*/
func stealTheGold(index int, arr []int) int {
	if index >= len(arr) {
		return 0
	}

	return int(math.Max(float64(arr[index])+float64(stealTheGold(index+2, arr)), float64(stealTheGold(index+1, arr))))
}

func houseRobberRecursive(arr []int) (r int) {
	return stealTheGold(0, arr)
}

/*
House robber problem with window sliding (kinda)
*/
func houseRobber(arr []int) int {
	maxGold := []int{}
	for index, val := range arr {
		prevMax := 0
		prevMax2Back := 0
		if index-1 >= 0 {
			prevMax = maxGold[index-1]
		}
		if index-2 >= 0 {
			prevMax2Back = maxGold[index-2]
		}

		maxGold = append(maxGold, int(math.Max(float64(val+prevMax2Back), float64(prevMax))))
	}
	return maxGold[len(maxGold)-1]
}

func houseRobberAlt(arr []int) int {
	currentMax, prevMax := 0, 0
	for i := 0; i < len(arr); i++ {
		currentHouse := arr[i]
		newMax := math.Max(float64(currentHouse)+float64(prevMax), float64(currentMax))
		prevMax = currentMax
		currentMax = int(newMax)
	}
	return currentMax
}
