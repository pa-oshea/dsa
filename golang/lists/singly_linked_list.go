package lists

type singlyLinkedList struct {
	length     int
	head, tail *node
}

func (s *singlyLinkedList) prepend(item any) {
	node := &node{data: item}
	if s.length == 0 {
		s.tail = node
	} else {
		node.next = s.head
	}
	s.head = node
	s.length++
}

func (s *singlyLinkedList) append(item any) {
	node := &node{data: item}
	if s.length == 0 {
		s.head = node
	} else {
		s.tail.next = node
	}

	s.tail = node
	s.length++

}

func (s *singlyLinkedList) insertAt(item any, idx int) {
	if s.length <= idx {
		s.append(item)
		return
	}

	if idx == 0 {
		s.prepend(item)
		return
	}

	current := s.head
	for current != nil && idx > 1 {
		current = current.next
		idx--
	}

	if current != nil {
		node := &node{data: item, next: current.next}
		current.next = node
		s.length++
	}
}

func (s *singlyLinkedList) removeAt(index int) any {
	// check if the index is out of bounds
	if index >= s.length {
		print("index out of bounds")
		return nil
	}

	// If there is only 1 node in the list, set head and tail to nil
	if s.length == 1 {
		result := s.head.data
		s.head = nil
		s.tail = nil
		s.length = 0
		return result
	}

	// If the index is 0, then set the head to the next node
	if index == 0 {
		head := s.head
		s.head = head.next
		s.length--
		return head.data
	}

	current := s.head
	for current != nil && index > 0 {
		current = current.next
		index--
	}

	if current != nil {
		nextNode := current.next
		result := nextNode.data
		current.next = nextNode.next
		nextNode.next = nil

		if index == s.length-1 {
			s.tail = current
		}
		return result
	}
	return nil
}

func (s *singlyLinkedList) get(idx int) any {
	if idx > s.length {
		return nil
	} else if idx == 0 {
		return s.head.data
	} else if idx == s.length-1 {
		return s.tail.data
	}

	current := s.head

	for current != nil && idx > 0 {
		current = current.next
		idx--

	}

	if current == nil {
		return nil
	}

	return current.data
}

func (s *singlyLinkedList) remove(item any) any {
	if s.head.data == item {
		result := s.head.data
		s.head = s.head.next
		s.length--
		return result
	}

	current := s.head
	for current.next != nil {
		if current.next.data == item {
			if current.next == s.tail {
				result := s.tail.data
				s.tail = current
				s.length--
				return result
			}

			result := current.next.data
			current.next = current.next.next
			s.length--
			return result
		}
		current = current.next
	}

	return nil
}
