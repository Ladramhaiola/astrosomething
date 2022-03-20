package main

import (
	"asteroids/ecs"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewAsteroid(x, y float64, size int) ecs.Entity {
	a := ecs.CreateEntity()

	if size == 0 {
		size = 3
	}
	radius := size * 15

	ecs.AddComponent(a, &Transform{X: x, Y: y})
	ecs.AddComponent(a, &Size{Radius: float64(radius)})
	ecs.AddComponent(a, &Collidable{Mask: MaskAsteroid})
	ecs.AddComponent(a, &Damageable{
		MaxHitPoints: size,
		CurHitPoints: size,
	})

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	vx := r.Intn(21) - r.Intn(21)
	vy := r.Intn(21) - r.Intn(21)

	ecs.AddComponent(a, &Velocity{
		X: float64(vx),
		Y: float64(vy),
	})

	// draw sprite
	dc := gg.NewContext(radius*2, radius*2)
	dc.DrawCircle(float64(radius), float64(radius), float64(radius-3))
	dc.SetRGB(r.Float64(), r.Float64(), r.Float64())
	dc.SetLineWidth(3)
	dc.Stroke()

	ecs.AddComponent(a, &Sprite{
		Image: ebiten.NewImageFromImage(dc.Image()),
	})

	return a
}
