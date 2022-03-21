package system

import (
	"log"

	"asteroids/component"
	"asteroids/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

type LifetimeSystem struct {
	*ecs.BaseSystem
}

func NewLifetimeSystem() *LifetimeSystem {
	s := &LifetimeSystem{&ecs.BaseSystem{}}
	ecs.RegisterSystem(s)
	// set signature
	ecs.SetSystemSignature(s, ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*component.Lifetime]()),
	)
	return s
}

// TODO: dt
func (s *LifetimeSystem) Update(dt float64) {
	for e := range s.Entities {
		lifetime := ecs.GetComponent[*component.Lifetime](e)
		lifetime.Time -= dt

		if lifetime.Time <= 0 {
			destroyEntity(e)
		}
	}
}

func (LifetimeSystem) Render(_ *ebiten.Image) {}

// TODO: think about normal destruction
func destroyEntity(e ecs.Entity) {
	destroyHook := ecs.GetComponent[component.OnDestroy](e)
	if destroyHook != nil {
		if err := destroyHook(); err != nil {
			log.Printf("[ERRO] destroy hook failed(%d): %s\n", e, err)
		}
	}

	ecs.DestroyEntity(e)
}
