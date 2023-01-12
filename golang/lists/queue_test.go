package lists

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
	assert.Equal(t, ans, 5)
	assert.Equal(t, queue.Length, 2)

	queue.Enqueue(11)
	ans, _ = queue.Dequeue()
	assert.Equal(t, ans, 7)
	ans, _ = queue.Dequeue()
	assert.Equal(t, ans, 9)
	assert.Equal(t, queue.Peek(), 11)
	ans, _ = queue.Dequeue()
	assert.Equal(t, ans, 11)
	ans, err := queue.Dequeue()
	assert.Equal(t, err, errors.New("first node is nil"))
	assert.Equal(t, ans, 0)
	assert.Equal(t, queue.Length, 0)

	queue.Enqueue(69)
	assert.Equal(t, queue.Peek(), 69)
	assert.Equal(t, queue.Length, 1)

}
