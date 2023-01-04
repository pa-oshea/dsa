package lists

type stack struct {
	length int
	top   *node
}

func (s *stack) push(item any) {
	s.length++
	node := &node{data: item}
	if s.top == nil {
		s.top = node
		return
	}
	node.next = s.top
	s.top = node
}

func (s *stack) pop() any {
	if s.length != 0 {

		s.length--
		result := s.top.data
		s.top = s.top.next
		return result
	}
	return nil
}

func (s *stack) peek() any {
	return s.top.data
}
