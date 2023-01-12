package lists

import (
	"errors"

	"github.com/pa-oshea/dsa/common"
)

type singlyLinkedList[T any] struct {
	length     int
	head, tail *common.Node[T]
}

func (s *singlyLinkedList[T]) prepend(item T) {
	node := &common.Node[T]{Data: item}
	if s.length == 0 {
		s.tail = node
	} else {
		node.Next = s.head
	}
	s.head = node
	s.length++
}

func (s *singlyLinkedList[T]) append(item T) {
	node := &common.Node[T]{Data: item}
	if s.length == 0 {
		s.head = node
	} else {
		s.tail.Next = node
	}

	s.tail = node
	s.length++

}

func (s *singlyLinkedList[T]) insertAt(item T, idx int) {
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
		current = current.Next
		idx--
	}

	if current != nil {
		node := &common.Node[T]{Data: item, Next: current.Next}
		current.Next = node
		s.length++
	}
}

func (s *singlyLinkedList[T]) remove(idx int) (T, error) {
	
	if idx >= s.length {
		// throw some error
		var result T
		return result, errors.New("index out of bounds")
	}

	s.length--
	if s.length == 0 {
		result := s.head.Data
		s.head = nil
		s.tail = nil
		return result, nil
	}

	if idx == 0 {
		result := s.head.Data
		s.head = s.head.Next
		return result, nil
	} 
	curr := s.head

	for curr != nil && idx > 1 {
		curr = curr.Next
		idx--
	}

	result := curr.Next.Data
	curr.Next = curr.Next.Next

	return result, nil

}

func (s *singlyLinkedList[T]) get(idx int) (T, error) {
	if idx > s.length {
		var result T
		return result, errors.New("index out of bounds")
	} else if idx == 0 {
		return s.head.Data, nil
	} else if idx == s.length-1 {
		return s.tail.Data, nil
	}

	current := s.head

	for current != nil && idx > 0 {
		current = current.Next
		idx--

	}

	if current == nil {
		var result T
		return result, errors.New("item not found")
	}

	return current.Data, nil
}

// This is stupid
// func (s *singlyLinkedList) remove(item any) any {
// 	if s.head.data == item {
// 		result := s.head.data
// 		s.head = s.head.next
// 		s.length--
// 		return result
// 	}
//
// 	current := s.head
// 	for current.next != nil {
// 		if current.next.data == item {
// 			if current.next == s.tail {
// 				result := s.tail.data
// 				s.tail = current
// 				s.length--
// 				return result
// 			}
//
// 			result := current.next.data
// 			current.next = current.next.next
// 			s.length--
// 			return result
// 		}
// 		current = current.next
// 	}
//
// 	return nil
// }
