package main

import (
	"asteroids/ecs"
	"math"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewBullet(x, y, angle float64) ecs.Entity {
	const speed = 300

	bullet := ecs.CreateEntity()

	ecs.AddComponent(bullet, &Transform{
		X:     x,
		Y:     y,
		Angle: angle,
	})
	ecs.AddComponent(bullet, &Velocity{
		X: math.Cos(angle) * speed,
		Y: math.Sin(angle) * speed,
	})
	ecs.AddComponent(bullet, &Size{Radius: 3})
	ecs.AddComponent(bullet, &Collidable{Mask: MaskBullet})

	// render bullet sprite
	dc := gg.NewContext(6, 6)
	dc.DrawCircle(3, 3, 3)
	dc.SetRGB(1, 1, 1)
	dc.Fill()

	ecs.AddComponent(bullet, &Sprite{
		Image: ebiten.NewImageFromImage(dc.Image()),
	})

	// suppose this should be configurable
	ecs.AddComponent(bullet, &Lifetime{Time: 5})

	return bullet
}
