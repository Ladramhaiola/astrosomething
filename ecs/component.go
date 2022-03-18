package ecs

import (
	"reflect"
)

// TODO: unit tests

type ComponentManger struct {
	// TODO: strings are kinda bad keys for performance
	// Map from type string pointer to a component type
	componentTypes map[string]ComponentType
	// Map from type string pointer to a component array
	componentArrays map[string]any
	// The component type to be assigned to the next registered component - starting at 0
	nextComponentType ComponentType
}

func RegisterComponent[T any](cm *ComponentManger, t T) {
	name := reflect.TypeOf(t).String()

	if _, ok := cm.componentArrays[name]; ok {
		return
	}

	// Add this component type to the component type map
	cm.componentTypes[name] = cm.nextComponentType
	// Create a ComponentArray pointer and add it to the component arrays map
	cm.componentArrays[name] = NewComponentArray[T]()
	// Increment the value so that the next component registered will be different
	cm.nextComponentType++
}

func GetComponentType[T any](cm *ComponentManger, t T) ComponentType {
	name := reflect.TypeOf(t).String()

	componentType, ok := cm.componentTypes[name]
	if !ok {
		panic("unknown component type")
	}
	// Return this component's type - used for creating signatures
	return componentType
}

func AddComponent[T any](cm *ComponentManger, entity Entity, component T) {
	// Add a component to the array for an entity
	GetComponentArray(cm, component).Insert(entity, component)
}

func RemoveComponent[T any](cm *ComponentManger, component T, entity Entity) {
	// Remove a component from the array for an entity
	GetComponentArray(cm, component).Remove(entity)
}

func GetComponent[T any](cm *ComponentManger, component T, entity Entity) T {
	// Get a reference to a component from the array for an entity
	return GetComponentArray(cm, component).Get(entity)
}

func GetComponentArray[T any](cm *ComponentManger, t T) *ComponentArray[T] {
	name := reflect.TypeOf(t).String()

	array, ok := cm.componentArrays[name]
	if !ok {
		panic("call to unregistered component array")
	}
	return array.(*ComponentArray[T])
}

// TODO: EntityDestroyed

func NewComponentManager() *ComponentManger {
	return &ComponentManger{
		componentTypes:  make(map[string]ComponentType),
		componentArrays: make(map[string]any, MaxComponents),
	}
}
