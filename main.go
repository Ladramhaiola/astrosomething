package main

import (
	"fmt"
	"log"

	"asteroids/colors"
	"asteroids/component"
	"asteroids/ecs"
	"asteroids/entity"
	"asteroids/events"
	"asteroids/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const screenWidth, screenHeight = 800, 600

// Game holds game state
type Game struct {
	score int
}

func (Game) Update() error {
	ecs.Update()
	return nil
}

func (Game) Draw(screen *ebiten.Image) {
	screen.Fill(colors.BackgroundColor)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
	ecs.Draw(screen)
}

func (Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ecs.Preallocate(1000, 100)

	ecs.RegisterComponent[*component.Transform]()
	ecs.RegisterComponent[*component.Velocity]()
	ecs.RegisterComponent[component.Clipable]()
	ecs.RegisterComponent[component.Acceleration]()
	ecs.RegisterComponent[*component.Size]()
	ecs.RegisterComponent[*component.UserControl]()
	ecs.RegisterComponent[*component.Sprite]()
	ecs.RegisterComponent[*component.Lifetime]()
	ecs.RegisterComponent[*component.Collidable]()
	ecs.RegisterComponent[*component.Damageable]()
	ecs.RegisterComponent[component.OnDestroy]()

	system.NewMovementSystem()
	system.NewClipSystem(screenWidth, screenHeight)
	system.NewCollisionSystem()
	system.NewLifetimeSystem()
	system.NewAsteroidSpawnerSystem(screenHeight, 3, 20)
	system.NewRenderSystem()

	player := entity.NewShipEntity(screenWidth/2, screenHeight/2)
	system.NewUserInputSystem(ecs.Entity(player))

	game := &Game{}

	ecs.AddListener(events.TAsteroidDestroyed, func(e ecs.Event) {
		event := e.(events.AsteroidDestroyed)
		game.score += event.Size * 100
		fmt.Println("score:", game.score)
	})

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("AstroPepega")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalln(err)
	}
}
