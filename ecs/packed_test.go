package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackedArray(t *testing.T) {
	ca := NewComponentArray[int]()

	// entities
	var a, b, c Entity = 1, 3, 7

	ca.Insert(a, 1)
	ca.Insert(b, 2)
	ca.Insert(c, 3)

	assert.Equal(t, 1, ca.GetData(a))
	assert.Equal(t, 3, ca.GetData(c))

	ca.Remove(b)

	ca.Insert(b, 10)
	assert.Equal(t, 10, ca.GetData(b))
}
