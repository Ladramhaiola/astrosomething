package events

import "asteroids/ecs"

const (
	TAsteroidDestroyed = iota
	TAsteroidSpawned
)

type AsteroidDestroyed struct {
	Size int
}

func (AsteroidDestroyed) Type() ecs.EventType {
	return TAsteroidDestroyed
}

type AsteroidSpawned struct {
	Size int
}

func (AsteroidSpawned) Type() ecs.EventType {
	return TAsteroidSpawned
}
