package queue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := Queue[int]{}
	queue.Enqueue(5)
	queue.Enqueue(7)
	queue.Enqueue(9)

	ans, _ := queue.Dequeue()
	assert.Equal(t, 5, ans)
	assert.Equal(t, 2, queue.Length)

	queue.Enqueue(11)
	ans, _ = queue.Dequeue()
	assert.Equal(t, 7, ans)
	ans, _ = queue.Dequeue()
	assert.Equal(t, 9, ans)
	assert.Equal(t, 11, queue.Peek())
	ans, _ = queue.Dequeue()
	assert.Equal(t, 11, ans)
	ans, err := queue.Dequeue()
	assert.Equal(t, errors.New("first node is nil"), err)
	assert.Equal(t, 0, ans)
	assert.Equal(t, 0, queue.Length)

	queue.Enqueue(69)
	assert.Equal(t, 69, queue.Peek())
	assert.Equal(t, 1, queue.Length)
}
