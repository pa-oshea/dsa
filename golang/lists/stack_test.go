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
	assert.Equal(t, ans, 9)
	assert.Equal(t, stack.length, 2)

	stack.push(11)
	ans, _ = stack.pop()
	assert.Equal(t, ans, 11)
	ans, _ = stack.pop()
	assert.Equal(t, ans, 7)
	assert.Equal(t, stack.peek(), 5)
	ans, _ = stack.pop()
	assert.Equal(t, ans, 5)
	_, err := stack.pop()
	assert.Equal(t, err, errors.New("no item found"))

	stack.push(69)
	assert.Equal(t, stack.peek(), 69)
	assert.Equal(t, stack.length, 1)
}
