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
	for current != nil && idx > 0 {
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
	current := s.head

	for current != nil && idx > 0 {
		current = current.next
		idx--

	}

	return current.data
}

func (s *singlyLinkedList) remove(item any) any {
	current := s.head
	for current != nil {
		if current.data == item {
			head := s.head
			result := head.data
			s.head = head

			head.next = nil

			return result
		}
	}

	return nil
}