package entity

import (
	"math"

	"asteroids/colors"
	"asteroids/component"
	"asteroids/ecs"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

var bulletImage = func() *ebiten.Image {
	// render bullet sprite
	dc := gg.NewContext(6, 6)
	dc.DrawCircle(3, 3, 3)
	dc.SetColor(colors.DefaultColor)
	dc.Fill()

	return ebiten.NewImageFromImage(dc.Image())
}()

func NewBullet(x, y, angle float64) ecs.Entity {
	const speed = 300

	bullet := ecs.CreateEntity()

	ecs.AddComponent(bullet, &component.Transform{
		X:     x,
		Y:     y,
		Angle: angle,
	})
	ecs.AddComponent(bullet, &component.Velocity{
		X: math.Cos(angle) * speed,
		Y: math.Sin(angle) * speed,
	})
	ecs.AddComponent(bullet, component.Clipable{})
	ecs.AddComponent(bullet, &component.Size{Radius: 3})
	ecs.AddComponent(bullet, &component.Collidable{Mask: component.MaskBullet})
	ecs.AddComponent(bullet, &component.Sprite{Image: bulletImage})

	// suppose this should be configurable
	ecs.AddComponent(bullet, &component.Lifetime{Time: 5})
	ecs.AddComponent(bullet, component.OnDestroy(func() error {
		trans := ecs.GetComponent[*component.Transform](bullet)
		// spawn bang effect
		NewBang(trans.X, trans.Y)
		return nil
	}))

	return bullet
}
