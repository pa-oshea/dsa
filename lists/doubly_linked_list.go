package lists

import (
	"errors"

	"github.com/pa-oshea/dsa/common"
)

type doublyLinkedList[T any] struct {
	head   *common.Node[T]
	tail   *common.Node[T]
	length int
}

func (d *doublyLinkedList[T]) prepend(item T) {
	node := &common.Node[T]{Data: item}
	if d.length == 0 {
		d.tail = node
	} else {
		d.head.Prev = node
		node.Next = d.head
	}

	d.head = node
	d.length++
}

func (d *doublyLinkedList[T]) append(item T) {
	node := &common.Node[T]{Data: item}
	if d.length == 0 {
		d.head = node
	} else {
		d.tail.Next = node
		node.Prev = d.tail
	}

	d.tail = node
	d.length++
}

func (d *doublyLinkedList[T]) get(idx int) (T, error) {
	if idx >= d.length {
		// throw some error
		var result T
		return result, errors.New("index out of bounds")
	} else if idx == 0 {
		return d.head.Data, nil
	} else if idx == d.length-1 {
		return d.tail.Data, nil
	}

	curr := d.head
	for curr != nil && idx > 0 {
		curr = curr.Next
		idx--
	}

	return curr.Data, nil

}

func (d *doublyLinkedList[T]) insertAt(item T, idx int) {
	if idx >= d.length {
		// throw some error
		return
	} else if idx == 0 {
		d.prepend(item)
	} else if idx == d.length-1 {
		d.append(item)
	}

	curr := d.head
	for curr != nil && idx > 0 {
		curr = curr.Next
		idx--
	}

	node := &common.Node[T]{Data: item, Prev: curr.Prev, Next: curr}
	node.Prev.Next = node
	curr.Prev = node
	d.length++
}

func (d *doublyLinkedList[T]) remove(idx int) (T, error) {
	if idx >= d.length {
		// throw some error
		var result T
		return result, errors.New("index out of bounds")
	}

	d.length--
	if d.length == 0 {
		result := d.head.Data
		d.head = nil
		d.tail = nil
		return result, nil
	}

	if idx == 0 {
		result := d.head.Data
		d.head = d.head.Next
		return result, nil
	} else if idx == d.length-1 {
		result := d.tail.Data
		d.tail = d.tail.Prev
		return result, nil
	}

	curr := d.head

	for curr != nil && idx > 0 {
		curr = curr.Next
		idx--
	}

	result := curr.Data
	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev
	curr.Next = nil
	curr.Prev = nil

	return result, nil
}
