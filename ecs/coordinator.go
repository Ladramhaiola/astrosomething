package ecs

import (
	bitmap "asteroids/bits"
)

// Coordinator is responsible for game world
type Coordinator struct {
	componentManager *ComponentManger
	entiryManager    *EntityManager
	systemManager    *SystemManager
}

func CreateEntity() Entity {
	return engine.entiryManager.CreateEntity()
}

func DestroyEntity(entity Entity) {
	engine.entiryManager.DestroyEntity(entity)
	engine.componentManager.EntityDestroyed(entity)
	engine.systemManager.EntityDestroyed(entity)
}

func RegisterComponent[T any]() {
	registerComponent[T](engine.componentManager)
}

func AddComponent[T any](entity Entity, component T) {
	addComponent(engine.componentManager, entity, component)

	signature := engine.entiryManager.GetSignature(entity)

	// calculate updated signature
	signature = bitmap.SetBit(
		signature,
		Signature(getComponentType[T](engine.componentManager)),
	)
	engine.entiryManager.SetSignature(entity, signature)

	// notify systems about changed signature
	engine.systemManager.EntitySignatureChanged(entity, signature)
}

func RemoveComponent[T any](entity Entity) {
	removeComponent[T](engine.componentManager, entity)

	signature := engine.entiryManager.GetSignature(entity)

	// calculate updated signature
	signature = bitmap.SetBit(
		signature,
		Signature(getComponentType[T](engine.componentManager)),
	)
	engine.entiryManager.SetSignature(entity, signature)

	// notify systems about changed signature
	engine.systemManager.EntitySignatureChanged(entity, signature)
}

func GetComponent[T any](entity Entity) T {
	return getComponent[T](engine.componentManager, entity)
}

func GetComponentType[T any]() ComponentType {
	return getComponentType[T](engine.componentManager)
}

func RegisterSystem(s System) {
	engine.systemManager.RegisterSystem(s)
}

func SetSystemSignature(s System, signature Signature) {
	engine.systemManager.SetSystemSignature(s, signature)
}

func NewCoordinator() *Coordinator {
	return &Coordinator{
		componentManager: NewComponentManager(),
		entiryManager:    NewEntityManager(),
		systemManager:    NewSystemManager(),
	}
}
