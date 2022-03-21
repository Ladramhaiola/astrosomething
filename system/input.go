package system

import (
	"math"

	"asteroids/component"
	"asteroids/ecs"
	"asteroids/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type UserInputSystem struct {
	*ecs.BaseSystem
	userEntityID ecs.Entity
}

func NewUserInputSystem(userID ecs.Entity) *UserInputSystem {
	s := &UserInputSystem{
		BaseSystem:   &ecs.BaseSystem{},
		userEntityID: userID,
	}
	ecs.RegisterSystem(s)
	// set signature
	ecs.SetSystemSignature(s, ecs.SignatureFromComponentTypes(
		ecs.GetComponentType[*component.Transform](),
		ecs.GetComponentType[*component.Velocity](),
		ecs.GetComponentType[component.Acceleration](),
		ecs.GetComponentType[*component.Size](),
		ecs.GetComponentType[*component.UserControl](),
		ecs.GetComponentType[*component.Sprite](),
	))
	return s
}

func (s *UserInputSystem) Update() {
	var (
		trans    = ecs.GetComponent[*component.Transform](s.userEntityID)
		accel    = ecs.GetComponent[component.Acceleration](s.userEntityID)
		control  = ecs.GetComponent[*component.UserControl](s.userEntityID)
		velocity = ecs.GetComponent[*component.Velocity](s.userEntityID)
		size     = ecs.GetComponent[*component.Size](s.userEntityID)
	)

	control.ShootTimer -= 1 / 60.

	// shooting
	if control.ShootTimer <= 0 {
		control.ShootTimer = control.ShootDelay

		// position relative to ship weapon mount with according rotation
		x := trans.X + math.Cos(trans.Angle)*size.Radius
		y := trans.Y + math.Sin(trans.Angle)*size.Radius
		// spawn new buller
		entity.NewBullet(x, y, trans.Angle)
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

func (UserInputSystem) Render(_ *ebiten.Image) {}
