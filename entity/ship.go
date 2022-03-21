package entity

import (
	"math"

	"asteroids/colors"
	"asteroids/component"
	"asteroids/ecs"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

var shipImage = func() *ebiten.Image {
	// render ship
	dc := gg.NewContext(60, 60)
	dc.MoveTo(30, 0)
	dc.LineTo(60, 60)
	dc.LineTo(30, 40)
	dc.LineTo(0, 60)
	dc.ClosePath()
	dc.SetColor(colors.DefaultColor)
	dc.SetLineWidth(2)
	dc.Stroke()

	return ebiten.NewImageFromImage(dc.Image())
}()

func NewShipEntity(x, y float64) ecs.Entity {
	ship := ecs.CreateEntity()

	ecs.AddComponent(ship, &component.Transform{
		X:        x,
		Y:        y,
		Angle:    0,
		Rotation: math.Pi,
	})
	ecs.AddComponent(ship, &component.Velocity{})
	ecs.AddComponent(ship, component.Acceleration(300))
	ecs.AddComponent(ship, component.Clipable{})
	ecs.AddComponent(ship, &component.Size{
		Width:  60,
		Height: 60,
		Radius: 30,
	})
	ecs.AddComponent(ship, &component.Collidable{Mask: component.MaskShip})

	ecs.AddComponent(ship, &component.UserControl{
		ShootTimer: 0,
		ShootDelay: 0.3,
	})

	ecs.AddComponent(ship, &component.Sprite{Image: shipImage})

	return ship
}
