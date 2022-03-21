package main

import (
	"asteroids/ecs"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

var bangImage = func() *ebiten.Image {
	// draw sprite
	dc := gg.NewContext(8, 8)
	dc.DrawRectangle(0, 0, 8, 8)
	dc.SetColor(HPColor)
	dc.Fill()

	return ebiten.NewImageFromImage(dc.Image())
}()

func NewBang(x, y float64) ecs.Entity {
	bang := ecs.CreateEntity()

	ecs.AddComponent(bang, &Transform{
		X: x,
		Y: y,
	})
	ecs.AddComponent(bang, &Lifetime{Time: 0.1})
	ecs.AddComponent(bang, &Sprite{Image: bangImage})

	return bang
}
