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

func mergeTwoLists(list1, list2 *common.Node[int]) *common.Node[int] {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Data < list2.Data {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list1.Next)
	return list2
}

func deleteMiddle(head *common.Node[int]) *common.Node[int] {
	length := 0
	curr := head
	for curr != nil {
		curr = curr.Next
		length++
	}

	index := length / 2
	curr = head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	if curr.Next == nil {
		curr.Next = nil
	} else {
		curr.Next = curr.Next.Next
	}
	return head
}
