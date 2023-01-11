package lists

import (
	"errors"

	"github.com/pa-oshea/dsa/common"
)

type Queue[T any] struct {
	Length      int
	first, last *common.Node[T]
}

func (q *Queue[T]) Enqueue(item T) {
	node := &common.Node[T]{Data: item}

	q.Length++
	if q.last == nil {
		q.last = node
		q.last.Next = node
		q.first = q.last
		return
	}

	q.last.Next = node
	q.last = node

}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.first == nil {
		var result T
		return result, errors.New("first node is nil")
	}

	result := q.first.Data
	q.first = q.first.Next
	q.Length--
	if q.Length == 0 {
		q.last = nil
	}

	return result, nil
}

func (q *Queue[T]) Peek() T {
	return q.first.Data
}
