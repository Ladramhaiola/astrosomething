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
