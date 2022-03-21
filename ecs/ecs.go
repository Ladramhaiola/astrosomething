package ecs

import (
	bitmap "asteroids/bits"

	"github.com/hajimehoshi/ebiten/v2"
)

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

func Update() {
	// fmt.Printf("[DEBUG] current entities count: %d (signatures: %d)\n",
	// 	engine.entiryManager.count,
	// 	len(engine.systemManager.signatures),
	// )
	// TODO: separate renderable & updatable?
	for _, system := range engine.systemManager.systems {
		system.Update()
	}
}

func Draw(screen *ebiten.Image) {
	for _, system := range engine.systemManager.systems {
		system.Render(screen)
	}
}

func Layout(outsideWidth, outsideHeight int) (int, int) {
	return ebiten.WindowSize()
}

func Preallocate(entities, components int) {
	MaxEntities = entities
	MaxComponents = components

	engine = NewCoordinator()
}
