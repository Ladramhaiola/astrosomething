package events

import "asteroids/ecs"

const (
	TAsteroidDestroyed = iota
	TAsteroidSpawned
	TWindowSizeChanged
	TInitialWindowLoaded
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

type WindowSizeChanged struct {
	Width, Height int
}

func (WindowSizeChanged) Type() ecs.EventType {
	return TWindowSizeChanged
}

type InitialWindowLoaded struct {
	Width, Height int
}

func (InitialWindowLoaded) Type() ecs.EventType {
	return TInitialWindowLoaded
}
