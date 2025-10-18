package game

import (
	"asteroids/internal/astroid"
	"asteroids/internal/config"
	"asteroids/internal/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	texBackground    rl.Texture2D
	Player           player.Player
	initialAsteroids int
	astroids         []astroid.Astroid
	gameOver         bool
}

func New() *Game {
	// Build the GUI window for the game
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Asteroid Shooter")
	rl.SetTargetFPS(60)

	// Build game
	g := &Game{
		texBackground:    rl.LoadTexture("resources/space_background.png"),
		Player:           player.New(config.ScreenWidth/2, config.ScreenHeight/2),
		initialAsteroids: 5,
		astroids:         astroid.CreateMultipleAstroids(5),
		gameOver:         false,
	}
	return g
}

func (g *Game) Update() {
	if !g.gameOver {
		g.Player.Update()
		for i := range g.astroids {
			g.astroids[i].Update()
		}

		// Check for collisions
		g.checkCollisions()
	}
}

func (g *Game) checkCollisions() {
	for i := len(g.astroids) - 1; i >= 0; i-- {
		// Check for collision between player and asteroid
		if rl.CheckCollisionCircles(
			g.Player.Position,
			g.Player.Size.X/4,
			g.astroids[i].Position,
			g.astroids[i].Size.X/4,
		) {
			g.gameOver = true
		}
	}
}

func (g *Game) Draw() {
	// Clear the screen
	rl.BeginDrawing()

	// Draw background
	src := rl.Rectangle{X: 0, Y: 0, Width: float32(g.texBackground.Width), Height: float32(g.texBackground.Height)}
	dst := rl.Rectangle{X: 0, Y: 0, Width: config.ScreenWidth, Height: config.ScreenHeight}
	rl.DrawTexturePro(g.texBackground, src, dst, rl.Vector2{}, 0, rl.White)

	// Draw player
	g.Player.Draw()

	// Draw astroids
	for i := range g.astroids {
		g.astroids[i].Draw()
	}

	if g.gameOver {
		drawCenteredText("Game over", config.ScreenHeight/2, 50, rl.Red)
	}

	// Draw score
	rl.DrawText("Score: 0", 10, 10, 20, rl.Gray)

	rl.EndDrawing()
}

func drawCenteredText(text string, y, fontSize int32, color rl.Color) {
	textWidth := rl.MeasureText(text, fontSize)
	rl.DrawText(text, config.ScreenWidth/2-textWidth/2, y, fontSize, color)
}

func (g *Game) Close() {
	g.Player.Close()
	rl.UnloadTexture(g.texBackground)
	rl.CloseWindow()
}

func (g *Game) WindowShouldClose() bool {
	return rl.WindowShouldClose()
}
