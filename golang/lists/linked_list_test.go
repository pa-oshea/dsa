package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {
	list := singlyLinkedList[int]{}
	list.prepend(5)
	assert.Equal(t, 5, list.head.Data)
	assert.Equal(t, 5, list.tail.Data)
	assert.Equal(t, 1, list.length)

	// remove 5
	ans, _ := list.remove(0)
	assert.Equal(t, ans, 5)

	list.append(7)
	assert.Equal(t, 7, list.head.Data)
	assert.Equal(t, 7, list.tail.Data)
	assert.Equal(t, 1, list.length)

	ans, _ = list.remove(0)
	assert.Equal(t, ans, 7)

	// add 5 , 7 back
	list.append(5)
	list.append(7)
	list.append(6)

	ans, _ = list.get(0)
	assert.Equal(t, ans, 5)
	ans, _ = list.get(1)
	assert.Equal(t, ans, 7)
	ans, _ = list.get(2)
	assert.Equal(t, ans, 6)

	list.insertAt(9, 1)

	ans, _ = list.get(1)
	assert.Equal(t, ans, 9)

	ans, _ = list.remove(1)
	assert.Equal(t, ans, 9)

	ans, _ = list.get(0)
	assert.Equal(t, ans, 5)
	ans, _ = list.get(1)
	assert.Equal(t, ans, 7)
	ans, _ = list.get(2)
	assert.Equal(t, ans, 6)
}

func Test_DoublyLinkedList(t *testing.T) {
	list := doublyLinkedList[int]{}
	list.prepend(5)
	assert.Equal(t, 5, list.head.Data)
	assert.Equal(t, 5, list.tail.Data)
	assert.Equal(t, 1, list.length)

	// remove 5
	ans, _ := list.remove(0)
	assert.Equal(t, ans, 5)

	list.append(7)
	assert.Equal(t, 7, list.head.Data)
	assert.Equal(t, 7, list.tail.Data)
	assert.Equal(t, 1, list.length)

	ans, _ = list.remove(0)
	assert.Equal(t, ans, 7)

	// add 5 , 7 back
	list.append(5)
	list.append(7)
	list.append(6)

	ans, _ = list.get(0)
	assert.Equal(t, ans, 5)
	ans, _ = list.get(1)
	assert.Equal(t, ans, 7)
	ans, _ = list.get(2)
	assert.Equal(t, ans, 6)

	list.insertAt(9, 1)

	ans, _ = list.get(1)
	assert.Equal(t, ans, 9)

	ans, _ = list.remove(1)
	assert.Equal(t, ans, 9)

	ans, _ = list.get(0)
	assert.Equal(t, ans, 5)
	ans, _ = list.get(1)
	assert.Equal(t, ans, 7)
	ans, _ = list.get(2)
	assert.Equal(t, ans, 6)

}
