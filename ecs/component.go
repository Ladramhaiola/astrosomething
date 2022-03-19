package ecs

import (
	"fmt"
	"log"
	"reflect"
)

// TODO: unit tests
// TODO: better interface
type PackedArray interface {
	Remove(entity Entity)
}

type ComponentManger struct {
	// TODO: strings are kinda bad keys for performance
	// Map from type string pointer to a component type
	componentTypes map[string]ComponentType
	// Map from type string pointer to a component array
	componentArrays map[string]PackedArray
	// The component type to be assigned to the next registered component - starting at 0
	nextComponentType ComponentType
}

func registerComponent[T any](cm *ComponentManger) {
	name := reflect.TypeOf((*T)(nil)).String()

	if _, ok := cm.componentArrays[name]; ok {
		log.Printf("component already registered: %s\n", name)
		return
	}

	// Add this component type to the component type map
	cm.componentTypes[name] = cm.nextComponentType
	// Create a ComponentArray pointer and add it to the component arrays map
	cm.componentArrays[name] = NewComponentArray[T]()
	// Increment the value so that the next component registered will be different
	cm.nextComponentType++
}

func getComponentType[T any](cm *ComponentManger) ComponentType {
	name := reflect.TypeOf((*T)(nil)).String()

	componentType, ok := cm.componentTypes[name]
	if !ok {
		panic(fmt.Sprintf("unknown component type %s", name))
	}
	// Return this component's type - used for creating signatures
	return componentType
}

func addComponent[T any](cm *ComponentManger, entity Entity, component T) {
	// Add a component to the array for an entity
	getComponentArray[T](cm).Insert(entity, component)
}

func removeComponent[T any](cm *ComponentManger, entity Entity) {
	// Remove a component from the array for an entity
	getComponentArray[T](cm).Remove(entity)
}

func getComponent[T any](cm *ComponentManger, entity Entity) T {
	// Get a reference to a component from the array for an entity
	return getComponentArray[T](cm).GetData(entity)
}

func getComponentArray[T any](cm *ComponentManger) *ComponentArray[T] {
	name := reflect.TypeOf((*T)(nil)).String()

	array, ok := cm.componentArrays[name]
	if !ok {
		panic(fmt.Sprintf("call to unregistered component array: %s\n", name))
	}
	return array.(*ComponentArray[T])
}

func (cm *ComponentManger) EntityDestroyed(entity Entity) {
	for _, component := range cm.componentArrays {
		component.Remove(entity)
	}
}

func NewComponentManager() *ComponentManger {
	return &ComponentManger{
		componentTypes:    make(map[string]ComponentType),
		componentArrays:   make(map[string]PackedArray, MaxComponents),
		nextComponentType: 1,
	}
}
