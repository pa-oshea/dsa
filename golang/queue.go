package main

type queue struct {
	length      int
	first, last *node
}

func (q *queue) enqueue(item int) {
	node := &node{data: item}
	q.length++
	if q.last == nil {
		q.last = node
		q.last.next = node
		q.first = q.last
		return
	}

	q.last.next = node
	q.last = node

}

func (q *queue) dequeue() int {
	if q.first == nil {
		return -1
	}

	result := q.first.data
	q.first = q.first.next
	q.length--
	if q.length == 0 {
		q.last = nil
	}

	return result
}

func (q *queue) peek() int {
	return q.first.data
}
