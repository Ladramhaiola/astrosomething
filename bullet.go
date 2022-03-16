package main

import (
	"math"

	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
)

var bulletImage *ebiten.Image

// TODO: check ebiten's vector graphics
func init() {
	dc := gg.NewContext(6, 6)
	dc.DrawCircle(3, 3, 3)
	dc.SetRGB(1, 1, 1)
	dc.Fill()

	bulletImage = ebiten.NewImageFromImage(dc.Image())
}

type Bullet struct {
	X, Y   float64
	VX, VY float64

	Radius   int
	Speed    int
	Lifetime float32

	ObjectType ObjectType
}

func NewBullet(x, y, angle float64) *Bullet {
	const speed = 300

	return &Bullet{
		X:        x,
		Y:        y,
		Radius:   3,
		Lifetime: 5,
		Speed:    speed,
		VX:       math.Cos(angle) * speed,
		VY:       math.Sin(angle) * speed,
	}
}

func (b *Bullet) Update() {
	b.Lifetime -= 1. / 60.

	if b.Lifetime <= 0 {
		// destroy self
	}

	b.X += b.VX / 60.
	b.Y += b.VY / 60.

	// TODO: generic clipper
	if b.X < 0 {
		b.X += screenWidth
	}
	if b.Y < 0 {
		b.Y += screenHeight
	}
	if b.X > screenWidth {
		b.X -= screenWidth
	}
	if b.Y > screenHeight {
		b.Y -= screenHeight
	}
}

func (b *Bullet) Draw(dst *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)
	dst.DrawImage(bulletImage, &op)
}
