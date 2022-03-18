package ecs

import (
	"golang.org/x/exp/constraints"
)

type Bits constraints.Unsigned

func SetBit[T Bits](b, flag T) T {
	return b | flag
}

func ClearBit[T Bits](b, flag T) T {
	return b &^ flag
}

func ToggleBit[T Bits](b, flag T) T {
	return b ^ flag
}

func HasBit[T Bits](b, flag T) bool {
	return b&flag != 0
}
