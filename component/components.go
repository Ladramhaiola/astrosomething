package component

import "github.com/hajimehoshi/ebiten/v2"

type CollisionMask uint

// collision masks
const (
	MaskIgnored CollisionMask = iota
	MaskShip
	MaskBullet
	MaskAsteroid
)

// Transform is responsible for object position
type Transform struct {
	X, Y     float64
	Angle    float64
	Rotation float64
}

type Velocity struct {
	X, Y float64
}

// Acceleration speed
type Acceleration int

// Size contains object size & radius
type Size struct {
	Width, Height float64
	Radius        float64
}

type Clipable struct{}

// Sprite contains object renderable part
type Sprite struct {
	Image *ebiten.Image
}

// Lifetime for temporary objects
type Lifetime struct {
	Time float64
}

// OnDestroy hook will be called before object death
type OnDestroy func() error

type Collidable struct {
	Mask CollisionMask
}

type Damageable struct {
	MaxHitPoints int
	CurHitPoints int
}

// UserControl for input handling
type UserControl struct {
	ShootTimer float64
	ShootDelay float64
}
