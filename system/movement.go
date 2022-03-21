package system

import (
	"asteroids/component"
	"asteroids/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

type MovementSystem struct {
	*ecs.BaseSystem
}

func NewMovementSystem() *MovementSystem {
	s := &MovementSystem{&ecs.BaseSystem{}}
	ecs.RegisterSystem(s)
	// set signature
	ecs.SetSystemSignature(s, ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*component.Transform](),
		ecs.GetComponentType[*component.Velocity](),
	))
	return s
}

func (s *MovementSystem) Update() {
	for e := range s.Entities {
		position := ecs.GetComponent[*component.Transform](e)
		velocity := ecs.GetComponent[*component.Velocity](e)

		position.X += velocity.X / 60.
		position.Y += velocity.Y / 60.
	}
}

func (MovementSystem) Render(_ *ebiten.Image) {}
