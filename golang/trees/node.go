package trees

type binary_node struct {
	value int
	left, right *binary_node
}

type general_node struct {
	value int
	children *[]general_node
}
