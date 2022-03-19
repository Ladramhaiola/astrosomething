package main

import (
	"asteroids/ecs"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth, screenHeight = 800, 600
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

	Sprite struct {
		Image *ebiten.Image
	}

	Lifetime struct {
		Time float64
	}
)

// game state
// TODO: in a normal way :)
var playerEntity ecs.Entity

type Game struct{}

func (Game) Update() error {
	ecs.Update()
	return nil
}

func (Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
	ecs.Draw(screen)
}

func (Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ecs.RegisterComponent[*Transform]()
	ecs.RegisterComponent[*Velocity]()
	ecs.RegisterComponent[Acceleration]()
	ecs.RegisterComponent[*Size]()
	ecs.RegisterComponent[*UserControl]()
	ecs.RegisterComponent[*Sprite]()
	ecs.RegisterComponent[*Lifetime]()

	NewPositionSystem()
	NewUserInputSystem()
	NewRenderSystem()
	NewLifetimeSystem()

	playerEntity = NewShipEntity()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatalln(err)
	}
}
