package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityManager(t *testing.T) {
	em := NewEntityManager()

	a := em.CreateEntity()
	b := em.CreateEntity()
	c := em.CreateEntity()

	em.SetSignature(a, 0b01)
	em.SetSignature(b, 0b10)

	assert.Equal(t, Signature(0b01), em.GetSignature(a))
	assert.Equal(t, Signature(0b10), em.GetSignature(b))
	assert.Equal(t, Signature(0b00), em.GetSignature(c))

	em.DestroyEntity(b)
	assert.Equal(t, Signature(0b00), em.GetSignature(b))
	assert.Equal(t, uint(2), em.count)

	d := em.CreateEntity()
	assert.Greater(t, d, c)
	assert.Equal(t, uint(3), em.count)
}
