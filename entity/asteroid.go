package entity

import (
	"math/rand"
	"time"

	"asteroids/component"
	"asteroids/ecs"
	"asteroids/events"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewAsteroid(x, y float64, size int) ecs.Entity {
	a := ecs.CreateEntity()

	if size == 0 {
		size = 3
	}
	radius := size * 15

	ecs.AddComponent(a, &component.Transform{X: x, Y: y})
	ecs.AddComponent(a, component.Clipable{})
	ecs.AddComponent(a, &component.Size{Radius: float64(radius)})
	ecs.AddComponent(a, &component.Collidable{Mask: component.MaskAsteroid})
	ecs.AddComponent(a, &component.Damageable{
		MaxHitPoints: size,
		CurHitPoints: size,
	})

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	vx := r.Intn(21) - r.Intn(21)
	vy := r.Intn(21) - r.Intn(21)

	ecs.AddComponent(a, &component.Velocity{
		X: float64(vx),
		Y: float64(vy),
	})

	// TODO: randomize asteroid sprite
	// draw sprite
	dc := gg.NewContext(radius*2, radius*2)
	dc.DrawCircle(float64(radius), float64(radius), float64(radius-3))
	dc.SetRGB(r.Float64(), r.Float64(), r.Float64())
	dc.SetLineWidth(3)
	dc.Stroke()

	ecs.AddComponent(a, &component.Sprite{
		Image: ebiten.NewImageFromImage(dc.Image()),
	})

	// spawn smaller asteroids
	ecs.AddComponent(a, component.OnDestroy(func() error {
		ecs.SendEvent(events.AsteroidDestroyed{Size: size})

		trans := ecs.GetComponent[*component.Transform](a)

		if size > 1 {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for i := 0; i <= r.Intn(3); i++ {
				// random x, y position shift
				rx := r.Float64() - r.Float64()
				ry := r.Float64() - r.Float64()
				rx *= (float64(size) * 3)
				ry *= (float64(size) * 3)

				NewAsteroid(trans.X+rx, trans.Y+ry, size-1)
			}
		}
		return nil
	}))

	ecs.SendEvent(events.AsteroidSpawned{})

	return a
}
