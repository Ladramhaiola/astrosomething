package bitmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBit(t *testing.T) {
	assert.Equal(t, uint(5), SetBit[uint](4, 0))
	assert.Equal(t, uint(16), SetBit[uint](0, 4))

	assert.Equal(t, uint(4), ClearBit[uint](5, 0))
	assert.Equal(t, uint(1), ClearBit[uint](3, 1))

	signature := SetBit[uint](1, 3)
	assert.Equal(t, uint(9), signature)
	assert.True(t, HasBit(signature, 3))
	signature = ToggleBit(signature, 3)
	assert.Equal(t, uint(1), signature)
	assert.False(t, HasBit(signature, 3))
}
