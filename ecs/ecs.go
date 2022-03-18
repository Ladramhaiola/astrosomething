package ecs

import bitmap "asteroids/bits"

// global engine
var engine = NewCoordinator()

type (
	Entity        uint32
	ComponentType uint32
	Signature     uint32
)

var (
	MaxEntities   = 500
	MaxComponents = 32
)

func SignatureFromComponentTypes(types ...ComponentType) Signature {
	var signature Signature
	for _, t := range types {
		signature = bitmap.SetBit(signature, Signature(t))
	}
	return signature
}
