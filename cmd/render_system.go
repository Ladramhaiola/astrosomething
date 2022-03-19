package main

import (
	"asteroids/ecs"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type RenderSystem struct{ *ecs.BaseSystem }

func NewRenderSystem() *RenderSystem {
	s := &RenderSystem{&ecs.BaseSystem{}}
	ecs.RegisterSystem(s)
	// set signature
	ecs.SetSystemSignature(s, ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*Transform](),
		ecs.GetComponentType[*Sprite](),
	))
	return s
}

func (s *RenderSystem) Update() {}

func (s *RenderSystem) Render(screen *ebiten.Image) {
	for e := range s.Entities {
		pos := ecs.GetComponent[*Transform](e)
		sprite := ecs.GetComponent[*Sprite](e)

		op := ebiten.DrawImageOptions{}

		width, height := sprite.Image.Size()

		// translate to image center
		op.GeoM.Translate(-float64(width)/2, -float64(height)/2)
		// rotate image
		op.GeoM.Rotate(pos.Angle + math.Pi/2)
		// move to current ship position in game world
		op.GeoM.Translate(pos.X, pos.Y)
		op.Filter = ebiten.FilterLinear

		screen.DrawImage(sprite.Image, &op)
	}
}
