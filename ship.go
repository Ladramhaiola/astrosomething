package main

import (
	"math"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var shipImage *ebiten.Image

const (
	width  = 60
	height = 60

	maxAngle = 360
)

func init() {
	dc := gg.NewContext(width, height)
	dc.MoveTo(30, 0)
	dc.LineTo(60, 60)
	dc.LineTo(30, 40)
	dc.LineTo(0, 60)
	dc.ClosePath()
	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(2)
	dc.Stroke()

	shipImage = ebiten.NewImageFromImage(dc.Image())
}

type Ship struct {
	// ship position
	X, Y float64
	// angle represents the ship's angle in XY plane
	Angle    float64
	Rotation float64
	// movement vectors
	VX, VY float64

	Acceleration int
}

func NewShip(x, y float64, acceleration int) *Ship {
	return &Ship{
		X:            x,
		Y:            y,
		Acceleration: acceleration,
		Rotation:     math.Pi,
	}
}

func (s *Ship) Update() {
	if inpututil.KeyPressDuration(ebiten.KeyD) >= 1 {
		s.Angle += s.Rotation / 60.
	}
	if inpututil.KeyPressDuration(ebiten.KeyA) >= 1 {
		s.Angle -= s.Rotation / 60.
	}
	if inpututil.KeyPressDuration(ebiten.KeyW) >= 1 {
		// acceleration vector
		vxDt := math.Cos(s.Angle) * float64(s.Acceleration/60)
		vyDt := math.Sin(s.Angle) * float64(s.Acceleration/60)

		s.VX += vxDt
		s.VY += vyDt
	}

	s.X += s.VX / 60.
	s.Y += s.VY / 60.

	s.VX -= s.VX / 60.
	s.VY -= s.VY / 60.

	// clip ship
	if s.X < 0 {
		s.X += screenWidth
	}
	if s.Y < 0 {
		s.Y += screenHeight
	}
	if s.X > screenWidth {
		s.X -= screenWidth
	}
	if s.Y > screenHeight {
		s.Y -= screenHeight
	}
}

func (s *Ship) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}

	op.GeoM.Translate(-width/2, -height/2)
	op.GeoM.Rotate(s.Angle + math.Pi/2)
	op.GeoM.Translate(s.X, s.Y)
	op.Filter = ebiten.FilterLinear

	screen.DrawImage(shipImage, &op)
}
