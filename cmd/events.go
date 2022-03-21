package main

import "asteroids/ecs"

const (
	EventTypeAsteroidDestroyed = iota
	EventTypeAsteroidSpawned
)

type AsteroidDestroyedEvent struct {
	Size int
}

func (AsteroidDestroyedEvent) Type() ecs.EventType {
	return EventTypeAsteroidDestroyed
}

type AsteroidSpawnedEvent struct{}

func (AsteroidSpawnedEvent) Type() ecs.EventType {
	return EventTypeAsteroidSpawned
}
