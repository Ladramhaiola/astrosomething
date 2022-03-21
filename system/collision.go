package system

import (
	"fmt"
	"math"

	"asteroids/component"
	"asteroids/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MaskAsteroid = component.MaskAsteroid
	MaskBullet   = component.MaskBullet
	MaskShip     = component.MaskShip
)

type CollisionSystem struct {
	*ecs.BaseSystem
}

func NewCollisionSystem() *CollisionSystem {
	s := &CollisionSystem{&ecs.BaseSystem{}}
	ecs.RegisterSystem(s)
	// set signature
	ecs.SetSystemSignature(s, ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*component.Collidable](),
		ecs.GetComponentType[*component.Transform](),
		ecs.GetComponentType[*component.Size](),
	))
	return s
}

// TODO: normal collision system :)
func (s *CollisionSystem) Update() {
	for e := range s.Entities {
		c := ecs.GetComponent[*component.Collidable](e)

		if c.Mask == MaskAsteroid {
			var (
				pos  = ecs.GetComponent[*component.Transform](e)
				size = ecs.GetComponent[*component.Size](e)
			)

			for other := range s.Entities {
				otherMask := ecs.GetComponent[*component.Collidable](other)

				if otherMask.Mask != MaskBullet && otherMask.Mask != MaskShip {
					continue
				}

				otherPos := ecs.GetComponent[*component.Transform](other)
				otherSize := ecs.GetComponent[*component.Size](other)
				if !collide(
					pos.X, pos.Y, size.Radius,
					otherPos.X, otherPos.Y, otherSize.Radius,
				) {
					continue
				}

				// collision happens
				if otherMask.Mask == MaskBullet {
					destroyEntity(other)

					health := ecs.GetComponent[*component.Damageable](e)
					health.CurHitPoints -= 1

					if health.CurHitPoints <= 0 {
						destroyEntity(e)
					}
				}

				if otherMask.Mask == MaskShip {
					fmt.Printf("[DEBUG] collision with ship\n")
				}
			}
		}
	}
}

func (CollisionSystem) Render(_ *ebiten.Image) {}

func collide(x1, y1, r1, x2, y2, r2 float64) bool {
	distance := math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)
	rdist := math.Pow(r1+r2, 2)
	return distance < rdist
}
