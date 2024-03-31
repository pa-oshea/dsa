package arrays

import (
	"github.com/pa-oshea/dsa/common"
)

// Given an array of integers rotate the array to left by P positions

// Method 1: Using a temp array
func rotate_with_temp(arr []int, p int) []int {
	// going to do this imperitively
	// Create a temperary array / slice
	var temp []int

	// loop through the array and when the index [arr[index]] >= p [Position]
	// add the arr[index] to the temp array
	for index, element := range arr {
		if index >= p {
			temp = append(temp, element)
		}
	}

	// Add the start of the arr to the temp arr up to the position provided
	for index, element := range arr {
		if index < p {
			temp = append(temp, element)
		}
	}
	return temp
}

// Method 2: Using append
func rotate_with_append(arr []int, p int) []int {
	return append(arr[p:], arr[:p]...)
}

// Method 3: Shifting one by one
func rotate_one_by_one_loop(arr []int, p int) []int {
	i := 0
	for i < p {
		tmp := arr[0]
		arr = common.Remove_ordered(arr, 0)
		arr = append(arr, tmp)
		i++
	}
	return arr
}

func get_gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return get_gcd(b, a%b)
	}
}

// Method 4: Some weird way off geeksforgeeks.org
func rotate_with_gcd(arr []int, d int, n int) []int{
	d = d % n
	var j, k, temp int
	gcd := get_gcd(d, n)

	for i := 0; i < gcd; i++ {
		temp = arr[i]
		j = i
		for {
			k = j + d
			if k >= n {
				k = k - n
			}
			if k == i {
				break
			}
			arr[j] = arr[k]
			j = k
		}
		arr[j] = temp
	}
    return arr
}
