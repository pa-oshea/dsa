package common

type Node[T any] struct {
	Data T
	// Data
	Value T
	// Prev
	Left *Node[T]
	// Next
	Right      *Node[T]
	Next, Prev *Node[T]
}

/*
Remove element from slice keeping the original order
@param slice: slice to remove element from
@param s: position of element to remove
*/
func Remove_ordered(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
