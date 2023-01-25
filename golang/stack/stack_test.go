package lists

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	stack := stack[int]{}
	stack.push(5)
	stack.push(7)
	stack.push(9)

	ans, _ := stack.pop()
	assert.Equal(t, 9, ans)
	assert.Equal(t, 2, stack.length)

	stack.push(11)
	ans, _ = stack.pop()
	assert.Equal(t, 11, ans)
	ans, _ = stack.pop()
	assert.Equal(t, 7, ans)
	assert.Equal(t, 5, stack.peek())
	ans, _ = stack.pop()
	assert.Equal(t, 5, ans)
	_, err := stack.pop()
	assert.Equal(t, errors.New("no item found"), err)

	stack.push(69)
	assert.Equal(t, 69, stack.peek())
	assert.Equal(t, 1, stack.length)
}
