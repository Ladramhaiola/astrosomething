package ecs_test

import (
	"asteroids/ecs"
	"fmt"
	"testing"
)

// TODO: API improvements & Rendering API & priority queue & parallel systems
// and complex querying

type Position struct {
	X, Y float64
}

type Velocity struct {
	X, Y float64
}

type Acceleration struct {
	Angle float64
	Speed int
}

type Sprite struct{}

type PhysicsSystem struct{ *ecs.BaseSystem }

func (s *PhysicsSystem) Update() {
	for entity := range s.BaseSystem.Entities {
		position := ecs.GetComponent[*Position](entity)
		velocity := ecs.GetComponent[*Velocity](entity)

		position.X += velocity.X / 60.
		position.Y += velocity.Y / 60.
	}
}

type RenderSystem struct{ *ecs.BaseSystem }

func (s *RenderSystem) Update() {
	for entity := range s.BaseSystem.Entities {
		position := ecs.GetComponent[*Position](entity)
		_ = ecs.GetComponent[*Sprite](entity)

		fmt.Printf("drawing sprite at %v\n", position)
	}
}

func TestEngine(t *testing.T) {
	ecs.RegisterComponent[*Position]()
	ecs.RegisterComponent[*Velocity]()
	ecs.RegisterComponent[*Acceleration]()
	ecs.RegisterComponent[*Sprite]()

	signature := ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*Position](),
		ecs.GetComponentType[*Velocity](),
	)

	physicsSystem := &PhysicsSystem{&ecs.BaseSystem{}}
	ecs.RegisterSystem(physicsSystem)
	ecs.SetSystemSignature(physicsSystem, signature)

	renderSystem := &RenderSystem{&ecs.BaseSystem{}}
	ecs.RegisterSystem(renderSystem)
	ecs.SetSystemSignature(renderSystem, ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*Position](),
		ecs.GetComponentType[*Sprite](),
	))

	entities := make([]ecs.Entity, 10)

	for i := range entities {
		entity := ecs.CreateEntity()

		entities[i] = entity

		ecs.AddComponent(entity, &Position{X: 2, Y: 10})
		ecs.AddComponent(entity, &Velocity{X: 3, Y: 4})
		ecs.AddComponent(entity, &Acceleration{Angle: 32, Speed: 200})

		if i%2 == 0 {
			ecs.AddComponent(entity, &Sprite{})
		}
	}

	physicsSystem.Update()
	renderSystem.Update()
}
