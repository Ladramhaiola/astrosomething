package system

import (
	"math"

	"asteroids/component"
	"asteroids/ecs"
	"asteroids/entity"
	"asteroids/events"

	"github.com/hajimehoshi/ebiten/v2"
)

type UserInputSystem struct {
	*ecs.BaseSystem
	userEntityID ecs.Entity
}

func NewUserInputSystem() *UserInputSystem {
	s := &UserInputSystem{
		BaseSystem: &ecs.BaseSystem{},
	}

	// wait for window initial size
	ecs.AddListener(events.TInitialWindowLoaded, func(e ecs.Event) {
		event := e.(events.InitialWindowLoaded)

		player := entity.NewShipEntity(
			float64(event.Width)/2,
			float64(event.Height)/2,
		)
		// set player entity
		s.userEntityID = player

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
	})

	return s
}

func (s *UserInputSystem) Update(dt float64) {
	var (
		trans    = ecs.GetComponent[*component.Transform](s.userEntityID)
		accel    = ecs.GetComponent[component.Acceleration](s.userEntityID)
		control  = ecs.GetComponent[*component.UserControl](s.userEntityID)
		velocity = ecs.GetComponent[*component.Velocity](s.userEntityID)
		size     = ecs.GetComponent[*component.Size](s.userEntityID)
	)

	control.ShootTimer -= dt

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
		trans.Angle += trans.Rotation * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		trans.Angle -= trans.Rotation * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		// acceleration vector
		vxDt := math.Cos(trans.Angle) * float64(accel) * dt
		vyDt := math.Sin(trans.Angle) * float64(accel) * dt

		velocity.X += vxDt
		velocity.Y += vyDt
	}

	// slowdown
	velocity.X -= velocity.X * dt
	velocity.Y -= velocity.Y * dt
}

func (UserInputSystem) Render(_ *ebiten.Image) {}
