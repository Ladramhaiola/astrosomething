package main

import (
	"asteroids/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

type PositionSystem struct{ *ecs.BaseSystem }

func NewPositionSystem() *PositionSystem {
	s := &PositionSystem{&ecs.BaseSystem{}}
	ecs.RegisterSystem(s)
	// set signature
	ecs.SetSystemSignature(s, ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*Transform](),
		ecs.GetComponentType[*Velocity](),
	))
	return s
}

func (s *PositionSystem) Update() {
	for e := range s.Entities {
		position := ecs.GetComponent[*Transform](e)
		velocity := ecs.GetComponent[*Velocity](e)

		position.X += velocity.X / 60.
		position.Y += velocity.Y / 60.

		// clip position
		if position.X < 0 {
			position.X += screenWidth
		}
		if position.Y < 0 {
			position.Y += screenHeight
		}
		if position.X > screenWidth {
			position.X -= screenWidth
		}
		if position.Y > screenHeight {
			position.Y -= screenHeight
		}
	}
}

func (s *PositionSystem) Render(_ *ebiten.Image) {}
