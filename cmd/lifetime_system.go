package main

import (
	"asteroids/ecs"
	"log"

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
			destroyEntity(e)
		}
	}
}

func (s *LifetimeSystem) Render(_ *ebiten.Image) {}

func destroyEntity(e ecs.Entity) {
	destroyHook := ecs.GetComponent[OnDestroy](e)
	if destroyHook != nil {
		if err := destroyHook(); err != nil {
			log.Printf("[ERRO] destroy hook failed(%d): %s\n", e, err)
		}
	}

	ecs.DestroyEntity(e)
}
