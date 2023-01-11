package common

type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}

type Binary_node struct {
	value       int
	left, right *Binary_node
}

type General_node struct {
	value    int
	children *[]General_node
}

/*
Remove element from slice keeping the original order
@param slice: slice to remove element from
@param s: position of element to remove
*/
func Remove_ordered(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
