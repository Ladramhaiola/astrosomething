package system

import (
	"asteroids/component"
	"asteroids/ecs"

	"github.com/hajimehoshi/ebiten/v2"
)

type ClipSystem struct {
	*ecs.BaseSystem
	width  float64
	height float64
}

func NewClipSystem(width, height float64) *ClipSystem {
	s := &ClipSystem{
		BaseSystem: &ecs.BaseSystem{},
		width:      width,
		height:     height,
	}

	ecs.RegisterSystem(s)
	// set signature
	ecs.SetSystemSignature(s, ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*component.Transform](),
		ecs.GetComponentType[component.Clipable](),
	))

	return s
}

// TODO: different clip types
func (s *ClipSystem) Update(_ float64) {
	for e := range s.Entities {
		trans := ecs.GetComponent[*component.Transform](e)

		if trans.X < 0 {
			trans.X += s.width
		}
		if trans.Y < 0 {
			trans.Y += s.height
		}
		if trans.X > s.width {
			trans.X -= s.width
		}
		if trans.Y > s.height {
			trans.Y -= s.height
		}
	}
}

func (ClipSystem) Render(_ *ebiten.Image) {}
