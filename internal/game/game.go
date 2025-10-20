package game

import (
	// built-ins
	"fmt"

	// internal packages
	"asteroids/internal/asteroid"
	"asteroids/internal/config"
	"asteroids/internal/player"
	"asteroids/internal/util"

	// external packages
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game struct holding the game state variables
type Game struct {
	texBackground     rl.Texture2D
	Player            player.Player
	shots             []player.Shot
	asteroids         []asteroid.Asteroid
	initialAsteroids  int
	astriodsDestroyed int
	gameOver          bool
	paused            bool
	victory           bool
}

// Build a new game
func New() *Game {
	// Build the GUI window for the game
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "asteroid Shooter")
	rl.SetTargetFPS(60)

	// Init the asteroids package
	asteroid.Init()

	// Build game struct
	g := &Game{
		texBackground:    rl.LoadTexture("resources/space_background.png"),
		initialAsteroids: 5,
	}

	// Set game to initial game state
	g.initState()

	return g
}

func (g *Game) initState() {
	// Create initial player state
	g.Player = player.New(config.ScreenWidth/2, config.ScreenHeight/2)

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
	g.astriodsDestroyed = 0
}

// Reset the game
func (g *Game) Reset() {
	g.initState()
}

func (g *Game) Update() {
	// If there are no asteroids left, we in
	if !g.gameOver && !g.victory && len(g.asteroids) == 0 {
		g.victory = true
	}

	// Toggle paused
	if rl.IsKeyPressed(rl.KeyP) {
		g.paused = !g.paused
	}

	// Restart the game
	if (g.gameOver || g.victory) && rl.IsKeyPressed(rl.KeyR) {
		g.Reset()
		return
	}

	// If it is not game over, update the frame
	if !g.paused && !g.victory && !g.gameOver {
		// Update the player
		g.Player.Update()

		// Update the asteroids
		for i := range g.asteroids {
			g.asteroids[i].Update()
		}

		// Update the shots
		for i := range g.shots {
			g.shots[i].Update()
		}

		// Fire shot
		if rl.IsKeyPressed(rl.KeySpace) {
			g.fireShot()
		}

		// Check for collisions
		g.checkCollisions()
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

	// Draw shots
	for i := range g.shots {
		g.shots[i].Draw()
	}

	// Draw asteroids
	for i := range g.asteroids {
		g.asteroids[i].Draw()
	}

	if g.gameOver {
		drawCenteredText("Game over", config.ScreenHeight/2, 50, rl.Red)
		drawCenteredText("Press R to restart", config.ScreenHeight/2+60, 20, rl.DarkGray)
	}

	if g.victory {
		drawCenteredText("YOU WIN!", config.ScreenHeight/2, 50, rl.Gray)
		drawCenteredText("Press R to restart", config.ScreenHeight/2+60, 20, rl.RayWhite)
	}

	// Draw score
	rl.DrawText(fmt.Sprintf("Score %d", g.astriodsDestroyed), 10, 10, 20, rl.Gray)
	pauseTextSize := rl.MeasureText("[P]ause", 20)
	rl.DrawText("[P]ause", config.ScreenWidth-pauseTextSize-10, 10, 20, rl.Gray)
	rl.EndDrawing()
}

func (g *Game) Close() {
	g.Player.Close()
	asteroid.Shutdown()
	rl.UnloadTexture(g.texBackground)
	rl.CloseWindow()
}

func (g *Game) checkCollisions() {
	for i := len(g.asteroids) - 1; i >= 0; i-- {
		// Check for collision between player and asteroid
		if rl.CheckCollisionCircles(
			g.Player.Position,
			g.Player.Size.X/4,
			g.asteroids[i].Position,
			g.asteroids[i].Size.X/4,
		) {
			g.gameOver = true
		}

		// Check for a collision between shots and the asteroid
		for j := range g.shots {
			// Loop through all the active shots
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
					g.astriodsDestroyed++
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

func drawCenteredText(text string, y, fontSize int32, color rl.Color) {
	textWidth := rl.MeasureText(text, fontSize)
	rl.DrawText(text, config.ScreenWidth/2-textWidth/2, y, fontSize, color)
}

func (g *Game) WindowShouldClose() bool {
	return rl.WindowShouldClose()
}
