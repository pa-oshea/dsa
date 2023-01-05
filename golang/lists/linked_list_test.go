package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {
	list := singlyLinkedList{}
	list.prepend(5)

	assert.Equal(t, 5, list.head.data)
	assert.Equal(t, 5, list.tail.data)
	assert.Equal(t, 1, list.length)

	list.removeAt(0)
	assert.Equal(t, 0, list.length)
	list.append(5)
	list.append(6)
	list.append(7)

	assert.Equal(t, 6, list.get(1))
	list.insertAt(9, 1)
	assert.Equal(t, 9, list.get(1))
	assert.Equal(t, 4, list.length)

	list.prepend(2)
	assert.Equal(t, 2, list.get(0))
	assert.Equal(t, 2, list.remove(2))
	assert.Equal(t, 9, list.remove(9))
	assert.Equal(t, 7, list.remove(7))
	assert.Equal(t, 2, list.length)
}
