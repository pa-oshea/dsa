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
	assert.Equal(t, 7, ans)

	// add 5 , 7 back
	list.append(5)
	list.append(7)
	list.append(6)

	ans, _ = list.get(0)
	assert.Equal(t, 5, ans)
	ans, _ = list.get(1)
	assert.Equal(t, 7, ans)
	ans, _ = list.get(2)
	assert.Equal(t, 6, ans)

	list.insertAt(9, 1)

	ans, _ = list.get(1)
	assert.Equal(t, 9, ans)

	ans, _ = list.remove(1)
	assert.Equal(t, 9, ans)

	ans, _ = list.get(0)
	assert.Equal(t, 5, ans)
	ans, _ = list.get(1)
	assert.Equal(t, 7, ans)
	ans, _ = list.get(2)
	assert.Equal(t, 6, ans)
}

func Test_DoublyLinkedList(t *testing.T) {
	list := doublyLinkedList[int]{}
	list.prepend(5)
	assert.Equal(t, 5, list.head.Data)
	assert.Equal(t, 5, list.tail.Data)
	assert.Equal(t, 1, list.length)

	// remove 5
	ans, _ := list.remove(0)
	assert.Equal(t, 5, ans)

	list.append(7)
	assert.Equal(t, 7, list.head.Data)
	assert.Equal(t, 7, list.tail.Data)
	assert.Equal(t, 1, list.length)

	ans, _ = list.remove(0)
	assert.Equal(t, 7, ans)

	// add 5 , 7 back
	list.append(5)
	list.append(7)
	list.append(6)

	ans, _ = list.get(0)
	assert.Equal(t, 5, ans)
	ans, _ = list.get(1)
	assert.Equal(t, 7, ans)
	ans, _ = list.get(2)
	assert.Equal(t, 6, ans)

	list.insertAt(9, 1)

	ans, _ = list.get(1)
	assert.Equal(t, 9, ans)

	ans, _ = list.remove(1)
	assert.Equal(t, 9, ans)

	ans, _ = list.get(0)
	assert.Equal(t, 5, ans)
	ans, _ = list.get(1)
	assert.Equal(t, 7, ans)
	ans, _ = list.get(2)
	assert.Equal(t, 6, ans)
}

func Test_Merge(t *testing.T) {
	list1 := singlyLinkedList[int]{}
	list2 := singlyLinkedList[int]{}

	list1.append(1)
	list1.append(2)
	list1.append(4)
	list2.append(1)
	list2.append(3)
	list2.append(4)
	list2.append(7)
	list2.append(1)
	list2.append(2)
	list2.append(6)

	deleteMiddle(list2.head)
	// mergeTwoLists(list1.head, list2.head)
}
