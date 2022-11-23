package common

/*
Remove element from slice keeping the original order
@param slice: slice to remove element from
@param s: position of element to remove
*/
func Remove_ordered(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
