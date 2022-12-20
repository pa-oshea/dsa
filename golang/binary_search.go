package main

import (
	"fmt"
	"math"
)

func binary_search(arr []int, val int) bool {
	high := len(arr)
	low := 0

	for low < high {
		mid := int(math.Floor(float64(low + (high-low)/2)))
		if arr[mid] == val {
			return true
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false

}

func helper(arr []int, low, high, val int) bool {
	if low >= high {
		return false
	}
	mid := low + (high-low)/2
	if val == arr[mid] {
		return true
	} else if val > arr[mid] {
		return helper(arr, mid+1, high, val)
	}
	return helper(arr, low, mid-1, val)
}

func binary_search_recursive(arr []int, val int) bool {
	high := len(arr)
	fmt.Println(val)
	return helper(arr, 0, high, val)
}

func binary_search_correct(arr []int, val int) bool {
	l := -1
	r := len(arr)
	for l+1 != r {
		m := (l + r) / 2
		if val == arr[m] {
			return true
		} else if val > arr[m] {
			l = m
		} else {
			r = m
		}
	}
	return false
}
