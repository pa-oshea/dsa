package lists

import (
	"errors"

	"github.com/pa-oshea/dsa/common"
)

type stack[T any] struct {
	length int
	top   *common.Node[T]
}

func (s *stack[T]) push(item T) {
	s.length++
	node := &common.Node[T]{Data: item}
	if s.top == nil {
		s.top = node
		return
	}
	node.Next = s.top
	s.top = node
}

func (s *stack[T]) pop() ( T, error ) {
	if s.length != 0 {

		s.length--
		result := s.top.Data
		s.top = s.top.Next
		return result, nil
	}
	var result T
	return result, errors.New("no item found")
}

func (s *stack[T]) peek() T {
	return s.top.Data
}
