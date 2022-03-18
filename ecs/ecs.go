package ecs

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
