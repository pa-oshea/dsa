package main

type stack struct {
	length int
	top   *node
}

func (s *stack) push(item int) {
	s.length++
	node := &node{data: item}
	if s.top == nil {
		s.top = node
		return
	}
	node.next = s.top
	s.top = node
}

func (s *stack) pop() int {
	if s.length != 0 {

		s.length--
		result := s.top.data
		s.top = s.top.next
		return result
	}
	// should be err maybe
	return -1
}

func (s *stack) peek() int {
	return s.top.data
}
