package main

import (
	"asteroids/ecs"
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type AsteroidSpawnerSystem struct {
	*ecs.BaseSystem
	maxAllowedCount int
	currentCount    int
	spawnTime       float64 // current time
	spawnDelay      float64 // between-spawn delay
}

func NewAsteroidSpawnerSystem(spawnDelay float64, maxAllowedCount int) *AsteroidSpawnerSystem {
	s := &AsteroidSpawnerSystem{
		BaseSystem:      &ecs.BaseSystem{},
		maxAllowedCount: maxAllowedCount,
		currentCount:    0,
		spawnTime:       0,
		spawnDelay:      2,
	}

	ecs.RegisterSystem(s)
	// listen for events
	ecs.AddListener(EventTypeAsteroidDestroyed, func(_ ecs.Event) {
		s.currentCount--
	})
	ecs.AddListener(EventTypeAsteroidSpawned, func(_ ecs.Event) {
		s.currentCount++
	})

	// TODO: normal load system
	for i := 0; i < 3; i++ {
		s.spawn()
	}

	return s
}

func (s *AsteroidSpawnerSystem) Update() {
	s.spawnTime -= 1 / 60.

	if s.spawnTime <= 0 && s.currentCount < s.maxAllowedCount {
		s.spawnTime = s.spawnDelay
		s.spawn()
	}
}

// spawn single asteroid
func (s *AsteroidSpawnerSystem) spawn() {
	fmt.Println(s.currentCount)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	y := r.Float64() * float64(screenHeight)

	// random size
	size := r.Intn(5)
	if size <= 2 {
		size = 2
	}

	NewAsteroid(0, y, size)
}

func (s *AsteroidSpawnerSystem) Render(_ *ebiten.Image) {}
