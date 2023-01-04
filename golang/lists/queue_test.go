package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {

	queue := queue{}
	queue.enqueue(5)
	queue.enqueue(7)
	queue.enqueue(9)

	assert.Equal(t, queue.dequeue(), 5)
	assert.Equal(t, queue.length, 2)

	queue.enqueue(11)
	assert.Equal(t, queue.dequeue(), 7)
	assert.Equal(t, queue.dequeue(), 9)
	assert.Equal(t, queue.peek(), 11)
	assert.Equal(t, queue.dequeue(), 11)
	assert.Equal(t, queue.dequeue(), nil)
	assert.Equal(t, queue.length, 0)

	queue.enqueue(69)
	assert.Equal(t, queue.peek(), 69)
	assert.Equal(t, queue.length, 1)

}
