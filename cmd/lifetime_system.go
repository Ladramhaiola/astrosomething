package main

import (
	"asteroids/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

type LifetimeSystem struct{ *ecs.BaseSystem }

func NewLifetimeSystem() *LifetimeSystem {
	s := &LifetimeSystem{&ecs.BaseSystem{}}
	ecs.RegisterSystem(s)
	// set signature
	ecs.SetSystemSignature(s, ecs.SignatureFromComponentTypes(ecs.GetComponentType[*Lifetime]()))
	return s
}

func (s *LifetimeSystem) Update() {
	for e := range s.Entities {
		lifetime := ecs.GetComponent[*Lifetime](e)
		lifetime.Time -= 1 / 60.

		if lifetime.Time <= 0 {
			// TODO: destroy hook or something
			ecs.DestroyEntity(e)
		}
	}
}

func (s *LifetimeSystem) Render(_ *ebiten.Image) {}
