package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumInRange(t *testing.T) {
	assert.Equal(t, 5050, sum_in_range(1, 100))
}
