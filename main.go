package main

import (
	"fmt"
	"log"
	"sync"

	"asteroids/colors"
	"asteroids/component"
	"asteroids/ecs"
	"asteroids/events"
	"asteroids/system"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const screenWidth, screenHeight = 800, 600

// Game holds game state
type Game struct {
	once  sync.Once
	pause bool
	score int
}

func (g *Game) Update() error {
	g.once.Do(func() {
		// do system's setup here
		w, h := ebiten.WindowSize()
		ecs.SendEvent(events.InitialWindowLoaded{
			Width:  w,
			Height: h,
		})
	})

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.pause = !g.pause
	}

	if g.pause {
		return nil
	}

	dt := 1. / float64(ebiten.MaxTPS())
	ecs.Update(dt)
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
	system.NewUserInputSystem()

	game := &Game{}

	ecs.AddListener(events.TAsteroidDestroyed, func(e ecs.Event) {
		event := e.(events.AsteroidDestroyed)
		game.score += event.Size * 100
		fmt.Println("score:", game.score)
	})

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetInitFocused(true)
	ebiten.SetWindowTitle("AstroPepega")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatalln(err)
	}
}
