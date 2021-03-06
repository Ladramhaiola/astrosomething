package system

import (
	"math/rand"
	"time"

	"asteroids/ecs"
	"asteroids/entity"
	"asteroids/events"

	"github.com/hajimehoshi/ebiten/v2"
)

type AsteroidSpawnerSystem struct {
	*ecs.BaseSystem
	maxAllowedCount int
	currentCount    int
	spawnTime       float64 // current time
	spawnDelay      float64 // between-spawn delay
	screenHeight    float64
}

func NewAsteroidSpawnerSystem(screenHeight, spawnDelay float64, maxAllowedCount int) *AsteroidSpawnerSystem {
	s := &AsteroidSpawnerSystem{
		BaseSystem:      &ecs.BaseSystem{},
		maxAllowedCount: maxAllowedCount,
		currentCount:    0,
		spawnTime:       0,
		spawnDelay:      2,
		screenHeight:    screenHeight,
	}

	ecs.RegisterSystem(s)
	// listen for events
	ecs.AddListener(events.TAsteroidDestroyed, func(_ ecs.Event) {
		s.currentCount--
	})
	ecs.AddListener(events.TAsteroidSpawned, func(_ ecs.Event) {
		s.currentCount++
	})

	// TODO: normal load system
	for i := 0; i < 3; i++ {
		s.spawn()
	}

	return s
}

func (s *AsteroidSpawnerSystem) Update(dt float64) {
	s.spawnTime -= dt

	if s.spawnTime <= 0 && s.currentCount < s.maxAllowedCount {
		s.spawnTime = s.spawnDelay
		s.spawn()
	}
}

// spawn single asteroid
func (s *AsteroidSpawnerSystem) spawn() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	y := r.Float64() * s.screenHeight

	// random size
	size := r.Intn(5)
	if size <= 2 {
		size = 2
	}

	entity.NewAsteroid(0, y, size)
}

func (s *AsteroidSpawnerSystem) Render(_ *ebiten.Image) {}
