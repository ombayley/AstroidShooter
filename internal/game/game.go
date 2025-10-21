package game

import (
	// built-ins
	"fmt"
	"sync"

	// internal packages
	"asteroids/internal/asteroid"
	"asteroids/internal/config"
	"asteroids/internal/player"
	"asteroids/internal/util"

	// external packages
	rl "github.com/gen2brain/raylib-go/raylib"
)

// --- Package Variables -----------------

var (
	texBackground rl.Texture2D
	loadOnce      sync.Once
	assetsLoaded  bool
)

// Init loads the background tilesheet. Call AFTER rl.InitWindow.
func Init() {
	loadOnce.Do(func() {
		texBackground = rl.LoadTexture("resources/space_background.png")
		assetsLoaded = true
	})
}

// Shutdown frees the background tilesheet. Call BEFORE rl.CloseWindow.
func Shutdown() {
	if assetsLoaded {
		rl.UnloadTexture(texBackground)
		assetsLoaded = false
	}
}

//
// --- Game ---------------------------------------------------------------------
//

// Game struct holding the game state variables
type Game struct {
	// Game elements
	Player    player.Player
	shots     []player.Shot
	asteroids []asteroid.Asteroid
	// Game state
	asteriodsDestroyed int
	gameOver           bool
	paused             bool
	victory            bool
}

// --- Lifecycle ----------------------------------------------------------------

// New builds a new game (window + assets + initial state).
func New() *Game {
	// Build Window
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "asteroid Shooter")
	rl.SetTargetFPS(60)

	// Init package assets (package textures)
	Init()
	asteroid.Init()
	player.Init()

	// Build game struct
	g := &Game{}

	// Set game to initial game state
	g.initState()

	return g
}

// Close all textures then window
func (g *Game) Close() {
	player.Shutdown()
	asteroid.Shutdown()
	Shutdown()
	rl.CloseWindow()
}

// --- State --------------------------------------------------------------------

// Setup initial game state
func (g *Game) initState() {
	// Create initial player state
	g.Player = player.New()

	// Setup shots
	g.shots = make([]player.Shot, config.MaxShots)
	for i := range g.shots {
		g.shots[i].Active = false
	}

	// Setup asteroids
	g.asteroids = asteroid.CreateAsteroids(5)

	// Reset game state
	g.gameOver = false
	g.victory = false
	g.paused = false
	g.asteriodsDestroyed = 0
}

// Reset the game
func (g *Game) Reset() {
	g.initState()
}

// --- Update / Draw ------------------------------------------------------------

func (g *Game) Update() {
	// Win condition
	if !g.gameOver && !g.victory && len(g.asteroids) == 0 {
		g.victory = true
	}

	// Pause
	if rl.IsKeyPressed(rl.KeyP) {
		g.paused = !g.paused
	}

	// Restart
	if (g.gameOver || g.victory) && rl.IsKeyPressed(rl.KeyR) {
		g.Reset()
		return
	}

	// Live update
	if !g.paused && !g.victory && !g.gameOver {
		// Update player
		g.Player.Update()

		// Update asteroids
		for i := range g.asteroids {
			g.asteroids[i].Update()
		}

		// Fire shot
		if rl.IsKeyPressed(rl.KeySpace) {
			g.fireShot()
		}

		// Update shots
		for i := range g.shots {
			g.shots[i].Update()
		}

		// Check collisions
		g.checkCollisions()
	}
}

func (g *Game) Draw() {
	// Clear the screen
	rl.BeginDrawing()

	// Background
	g.drawBackground()

	// Player
	g.Player.Draw()

	// Shots
	for i := range g.shots {
		g.shots[i].Draw()
	}

	// Asteroids
	for i := range g.asteroids {
		g.asteroids[i].Draw()
	}

	// Game over overlay
	if g.gameOver {
		drawCenteredText("Game over", config.ScreenHeight/2, 50, rl.Red)
		drawCenteredText("Press R to restart", config.ScreenHeight/2+60, 20, rl.DarkGray)
	}
	// Victory overlay
	if g.victory {
		drawCenteredText("YOU WIN!", config.ScreenHeight/2, 50, rl.Gray)
		drawCenteredText("Press R to restart", config.ScreenHeight/2+60, 20, rl.RayWhite)
	}

	// HUD
	rl.DrawText(fmt.Sprintf("Score %d", g.asteriodsDestroyed), 10, 10, 20, rl.Gray)
	pauseTextSize := rl.MeasureText("[P]ause", 20)
	rl.DrawText("[P]ause", config.ScreenWidth-pauseTextSize-10, 10, 20, rl.Gray)

	rl.EndDrawing()
}

// --- Internals ----------------------------------------------------------------

// Check for collisions between game elements
func (g *Game) checkCollisions() {
	for i := len(g.asteroids) - 1; i >= 0; i-- {
		// Check collision between player and asteroid
		if rl.CheckCollisionCircles(
			g.Player.Position,
			g.Player.Size.X/4,
			g.asteroids[i].Position,
			g.asteroids[i].Size.X/4,
		) {
			g.gameOver = true
		}

		// Check collision between shots and asteroids
		for j := range g.shots {
			if g.shots[j].Active {
				// If it has collided with an asteroid
				if rl.CheckCollisionCircles(
					g.shots[j].Position,
					g.shots[j].Radius,
					g.asteroids[i].Position,
					g.asteroids[i].Size.X/2,
				) {
					// Destroy the shot and split the asteroid
					g.shots[j].Active = false

					// The asteroid shot split according to our rules
					newAsteroids := asteroid.Break(g.asteroids[i])
					g.asteroids = append(g.asteroids, newAsteroids...)

					// Remove the original asteroid from the slice
					g.asteroids = append(g.asteroids[:i], g.asteroids[i+1:]...)

					// Increase our score
					g.asteriodsDestroyed++
					break
				}
			}
		}
	}
}

func (g *Game) fireShot() {
	for i := range g.shots {
		// Find the first inactive shot
		if !g.shots[i].Active {
			// Start at the players position
			g.shots[i].Position = g.Player.Position
			g.shots[i].Active = true

			// Get the players direction
			shotDirection := util.DirectionVector(g.Player.Rotation)

			// Get the initial velocity
			shotVelocity := rl.Vector2Scale(shotDirection, config.ShotSpeed)
			// Account for the players speed
			playerVelocity := rl.Vector2Scale(g.Player.Speed, g.Player.Acceleration)

			// Fire the shot, realative to the players speed
			g.shots[i].Speed = rl.Vector2Add(playerVelocity, shotVelocity)

			g.shots[i].Radius = 2
			// Break after one shot
			break
		}
	}
}

// --- Helpers ----------------------------------------------------------------

// Draw centered text
func drawCenteredText(text string, y, fontSize int32, color rl.Color) {
	textWidth := rl.MeasureText(text, fontSize)
	rl.DrawText(text, config.ScreenWidth/2-textWidth/2, y, fontSize, color)
}

// Draw the background
func (g *Game) drawBackground() {
	src := rl.Rectangle{X: 0, Y: 0, Width: float32(texBackground.Width), Height: float32(texBackground.Height)}
	dst := rl.Rectangle{X: 0, Y: 0, Width: config.ScreenWidth, Height: config.ScreenHeight}
	rl.DrawTexturePro(texBackground, src, dst, rl.Vector2{}, 0, rl.White)
}

// Check if the window should close (remove rl import req. in main)
func (g *Game) WindowShouldClose() bool {
	return rl.WindowShouldClose()
}
