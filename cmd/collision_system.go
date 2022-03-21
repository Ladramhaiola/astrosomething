package main

import (
	"asteroids/ecs"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionSystem struct{ *ecs.BaseSystem }

func NewCollisionSystem() *CollisionSystem {
	s := &CollisionSystem{&ecs.BaseSystem{}}
	ecs.RegisterSystem(s)
	// set signature
	ecs.SetSystemSignature(s, ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*Collidable](),
		ecs.GetComponentType[*Transform](),
		ecs.GetComponentType[*Size](),
	))
	return s
}

// TODO: normal collision system :)
func (s *CollisionSystem) Update() {
	for e := range s.Entities {
		c := ecs.GetComponent[*Collidable](e)

		if c.Mask == MaskAsteroid {
			var (
				pos  = ecs.GetComponent[*Transform](e)
				size = ecs.GetComponent[*Size](e)
			)

			for other := range s.Entities {
				otherMask := ecs.GetComponent[*Collidable](other)

				if otherMask.Mask != MaskBullet && otherMask.Mask != MaskShip {
					continue
				}

				otherPos := ecs.GetComponent[*Transform](other)
				otherSize := ecs.GetComponent[*Size](other)
				if !collide(
					pos.X, pos.Y, size.Radius,
					otherPos.X, otherPos.Y, otherSize.Radius,
				) {
					continue
				}

				// collision happens
				if otherMask.Mask == MaskBullet {
					destroyEntity(other)

					health := ecs.GetComponent[*Damageable](e)
					health.CurHitPoints -= 1

					if health.CurHitPoints <= 0 {
						destroyEntity(e)
					}
				}
			}
		}
	}
}

func (s *CollisionSystem) Render(_ *ebiten.Image) {}

func collide(x1, y1, r1, x2, y2, r2 float64) bool {
	distance := math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)
	rdist := math.Pow(r1+r2, 2)
	return distance < rdist
}
