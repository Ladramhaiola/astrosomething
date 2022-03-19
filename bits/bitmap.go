package bitmap

import (
	"golang.org/x/exp/constraints"
)

type Bits constraints.Unsigned

func SetBit[T Bits](b, pos T) T {
	return b | (1 << pos)
}

func ClearBit[T Bits](b, pos T) T {
	return b &^ (1 << pos)
}

func ToggleBit[T Bits](b, pos T) T {
	return b ^ (1 << pos)
}

func HasBit[T Bits](b, pos T) bool {
	return b&(1<<pos) != 0
}
