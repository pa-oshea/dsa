package queue

// Circular queue / ring buffer
// front and rear should be set to -1
type cqueue struct {
	size, front, rear int
	items             [5]int
}

func (cq *cqueue) isFull() bool {
	return cq.front == 0 && cq.rear == cq.size-1 || cq.front == cq.rear+1
}

func (cq *cqueue) isEmpty() bool {
	return cq.front == -1
}

func (cq *cqueue) enqueue(element int) {
	if cq.isFull() {
		return
	}
	if cq.front == -1 {
		cq.front = 0
	}

	cq.rear = (cq.rear + 1) % cq.size
	cq.items[cq.rear] = element
}
