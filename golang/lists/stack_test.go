package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	stack := stack{}
	stack.push(5)
	stack.push(7)
	stack.push(9)

	assert.Equal(t, stack.pop(), 9)
	assert.Equal(t, stack.length, 2)

	stack.push(11)
	assert.Equal(t, stack.pop(), 11)
	assert.Equal(t, stack.pop(), 7)
	assert.Equal(t, stack.peek(), 5)
	assert.Equal(t, stack.pop(), 5)
	assert.Equal(t, stack.pop(), nil)

	stack.push(69)
	assert.Equal(t, stack.peek(), 69)
	assert.Equal(t, stack.length, 1)
}

