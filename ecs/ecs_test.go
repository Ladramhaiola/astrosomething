package ecs

import (
	"fmt"
	"testing"
)

type Set struct {
	enitities map[Entity]struct{}
}

func (s *Set) Insert(e Entity) {
	s.enitities[e] = struct{}{}
}

func (s *Set) Remove(e Entity) {
	delete(s.enitities, e)
}

type PhysicsSystem struct {
	*Set
}

type Gravity struct {
	Force float64
}

type RigidBody struct {
	Velocity     [2]float64
	Acceleration [2]float64
}

type Transform struct {
	Position [2]float64
	Rotation [2]float64
	Scale    [2]float64
}

func (s *PhysicsSystem) Update(c *Coordinator) {
	for e := range s.enitities {
		rigidBody := GetComponent[*RigidBody](c, e)
		gravity := GetComponent[*Gravity](c, e)

		fmt.Println(rigidBody)
		fmt.Println(gravity)

		rigidBody.Velocity[1] -= gravity.Force
	}
}

type PositionSystem struct {
	*Set
}

func (s *PositionSystem) Update(c *Coordinator) {
	for e := range s.enitities {
		transform := GetComponent[*Transform](c, e)
		fmt.Println("transform", transform)
	}
}

func TestCoordinatorFlow(t *testing.T) {
	c := NewCoordinator()

	RegisterComponent[*Gravity](c)
	RegisterComponent[*RigidBody](c)
	RegisterComponent[*Transform](c)

	var signature Signature
	signature = SetBit(signature, Signature(GetComponentType[*Gravity](c)))
	signature = SetBit(signature, Signature(GetComponentType[*RigidBody](c)))
	signature = SetBit(signature, Signature(GetComponentType[*Transform](c)))

	physicsSystem := &PhysicsSystem{&Set{enitities: make(map[Entity]struct{})}}
	RegisterSystem(c, physicsSystem)
	SetSystemSignature[*PhysicsSystem](c, signature)

	positionSystem := &PositionSystem{&Set{enitities: make(map[Entity]struct{})}}
	RegisterSystem(c, positionSystem)
	SetSystemSignature[*PositionSystem](c, Signature(GetComponentType[*Transform](c)))

	entities := make([]Entity, 2)

	for i := range entities {
		entity := c.CreateEntity()

		entities[i] = entity

		AddComponent(c, entity, &Gravity{Force: 0.5})
		AddComponent(c, entity, &RigidBody{
			Velocity:     [2]float64{2.0, 2.0},
			Acceleration: [2]float64{3.0, 3.0},
		})
	}

	AddComponent(c, c.CreateEntity(), &Transform{Scale: [2]float64{56, 45}})

	fmt.Println(physicsSystem)
	physicsSystem.Update(c)
	physicsSystem.Update(c)
	positionSystem.Update(c)
}
