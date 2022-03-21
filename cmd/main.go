package main

import (
	"asteroids/ecs"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth, screenHeight = 800, 600
)

type CollisionMask uint

const (
	MaskIgnored CollisionMask = iota
	MaskShip
	MaskBullet
	MaskAsteroid
)

var (
	DefaultColor    = color.RGBA{R: 222, G: 222, B: 222, A: 255}
	BackgroundColor = color.RGBA{R: 16, G: 16, B: 16, A: 255}
	AmmoColor       = color.RGBA{R: 123, G: 200, B: 14, A: 255}
	BoostColor      = color.RGBA{R: 76, G: 195, B: 217, A: 255}
	HPColor         = color.RGBA{R: 241, G: 103, B: 69, A: 255}
	SkillPointColor = color.RGBA{R: 255, G: 198, B: 193, A: 255}
)

// Components
type (
	Transform struct {
		X, Y     float64
		Angle    float64
		Rotation float64
	}

	Velocity struct {
		X, Y float64
	}

	Acceleration int

	Size struct {
		Width, Height float64
		Radius        float64
	}

	UserControl struct {
		ShootTimer float64
		ShootDelay float64
	}

	Sprite struct{ Image *ebiten.Image }

	Lifetime  struct{ Time float64 }
	OnDestroy func() error

	// TODO: classify entites somehow
	Collidable struct{ Mask CollisionMask }

	Damageable struct {
		MaxHitPoints int
		CurHitPoints int
	}
)

// game state
// TODO: in a normal way :)
var playerEntity ecs.Entity

type Game struct {
	score int
}

func (Game) Update() error {
	ecs.Update()
	return nil
}

func (Game) Draw(screen *ebiten.Image) {
	screen.Fill(BackgroundColor)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
	ecs.Draw(screen)
}

func (Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ecs.Preallocate(1000, 100)

	ecs.RegisterComponent[*Transform]()
	ecs.RegisterComponent[*Velocity]()
	ecs.RegisterComponent[Acceleration]()
	ecs.RegisterComponent[*Size]()
	ecs.RegisterComponent[*UserControl]()
	ecs.RegisterComponent[*Sprite]()
	ecs.RegisterComponent[*Lifetime]()
	ecs.RegisterComponent[*Collidable]()
	ecs.RegisterComponent[*Damageable]()
	ecs.RegisterComponent[OnDestroy]()

	NewPositionSystem()
	NewUserInputSystem()
	NewRenderSystem()
	NewLifetimeSystem()
	NewCollisionSystem()
	NewAsteroidSpawnerSystem(5, 30)

	playerEntity = NewShipEntity(300, 300)

	game := &Game{}

	ecs.AddListener(EventTypeAsteroidDestroyed, func(e ecs.Event) {
		event := e.(AsteroidDestroyedEvent)
		game.score += event.Size * 100
		fmt.Println("score:", game.score)
	})

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalln(err)
	}
}
