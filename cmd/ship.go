package main

import (
	"asteroids/ecs"
	"math"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewShipEntity() ecs.Entity {
	ship := ecs.CreateEntity()

	ecs.AddComponent(ship, &Transform{
		X:        300,
		Y:        300,
		Angle:    0,
		Rotation: math.Pi,
	})
	ecs.AddComponent(ship, &Velocity{})
	ecs.AddComponent(ship, Acceleration(300))
	ecs.AddComponent(ship, &Size{
		Width:  60,
		Height: 60,
		Radius: 30,
	})
	ecs.AddComponent(ship, &Collidable{Mask: MaskShip})

	ecs.AddComponent(ship, &UserControl{
		ShootDelay: 0.3,
	})

	// render ship
	dc := gg.NewContext(60, 60)
	dc.MoveTo(30, 0)
	dc.LineTo(60, 60)
	dc.LineTo(30, 40)
	dc.LineTo(0, 60)
	dc.ClosePath()
	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(2)
	dc.Stroke()

	ecs.AddComponent(ship, &Sprite{
		Image: ebiten.NewImageFromImage(dc.Image()),
	})

	return ship
}
