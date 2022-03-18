package ecs

// TODO: entity destriyed event system
// TODO: consider component array sizes

type ComponentArray[T any] struct {
	// The packed array of components (of generic type T),
	// set to a specified maximum amount, matching the maximum number
	// of entities allowed to exist simultaneously, so that each entity
	// has a unique spot.
	array []T
	// Map from an entity ID to an array index.
	entityToIndex map[Entity]int
	// Map from an array index to an entity ID.
	indexToEntity map[int]Entity
	// Total size of valid entries in the array.
	size int
}

func (ca *ComponentArray[T]) Insert(entity Entity, component T) {
	// TODO: do nothing on repeating inserts
	index := ca.size

	ca.entityToIndex[entity] = index
	ca.indexToEntity[index] = entity

	ca.array[index] = component

	ca.size++
}

func (ca *ComponentArray[T]) Remove(entity Entity) {
	// Copy element at end into deleted element's place to maintain density
	indexOfRemovedEntity := ca.entityToIndex[entity]
	indexOfLastElement := ca.size - 1

	ca.array[indexOfRemovedEntity] = ca.array[indexOfLastElement]

	// Update map to point to moved spot
	entityOfLastElement := ca.indexToEntity[indexOfLastElement]
	ca.entityToIndex[entityOfLastElement] = indexOfRemovedEntity
	ca.indexToEntity[indexOfRemovedEntity] = entityOfLastElement

	delete(ca.entityToIndex, entity)
	delete(ca.indexToEntity, indexOfLastElement)

	ca.size--
}

func (ca *ComponentArray[T]) Get(entity Entity) T {
	return ca.array[ca.entityToIndex[entity]]
}

func NewComponentArray[T any]() *ComponentArray[T] {
	return &ComponentArray[T]{
		array:         make([]T, MaxEntities),
		entityToIndex: make(map[Entity]int),
		indexToEntity: make(map[int]Entity),
	}
}
