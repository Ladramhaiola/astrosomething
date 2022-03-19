package main

import (
	"asteroids/ecs"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type UserInputSystem struct{ *ecs.BaseSystem }

func NewUserInputSystem() *UserInputSystem {
	s := &UserInputSystem{&ecs.BaseSystem{}}
	ecs.RegisterSystem(s)
	// set signature
	ecs.SetSystemSignature(s, ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*Transform](),
		ecs.GetComponentType[*Velocity](),
		ecs.GetComponentType[Acceleration](),
		ecs.GetComponentType[*Size](),
		ecs.GetComponentType[*UserControl](),
		ecs.GetComponentType[*Sprite](),
	))
	return s
}

func (s *UserInputSystem) Update() {
	var (
		trans    = ecs.GetComponent[*Transform](playerEntity)
		accel    = ecs.GetComponent[Acceleration](playerEntity)
		control  = ecs.GetComponent[*UserControl](playerEntity)
		velocity = ecs.GetComponent[*Velocity](playerEntity)
		size     = ecs.GetComponent[*Size](playerEntity)
	)

	control.ShootTimer -= 1 / 60. // TODO: find how delta time in ebiten works

	// shooting
	if ebiten.IsKeyPressed(ebiten.KeySpace) && control.ShootTimer <= 0 {
		// position relative to ship weapon mount with according rotation
		x := trans.X + math.Cos(trans.Angle)*size.Radius
		y := trans.Y + math.Sin(trans.Angle)*size.Radius
		// spawn new buller
		NewBullet(x, y, trans.Angle)
		control.ShootTimer = control.ShootDelay
	}

	// movement
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		trans.Angle += trans.Rotation / 60.
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		trans.Angle -= trans.Rotation / 60.
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		// acceleration vector
		vxDt := math.Cos(trans.Angle) * float64(accel/60)
		vyDt := math.Sin(trans.Angle) * float64(accel/60)

		velocity.X += vxDt
		velocity.Y += vyDt
	}

	// slowdown
	velocity.X -= velocity.X / 60.
	velocity.Y -= velocity.Y / 60.
}

func (s *UserInputSystem) Render(screen *ebiten.Image) {}
