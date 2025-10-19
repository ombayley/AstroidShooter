package game

import (
	"asteroids/internal/astroid"
	"asteroids/internal/config"
	"asteroids/internal/player"
	"asteroids/internal/util"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	texBackground     rl.Texture2D
	Player            player.Player
	shots             []player.Shot
	initialAsteroids  int
	astroids          []astroid.Astroid
	gameOver          bool
	paused            bool
	victory           bool
	astriodsDestroyed int
}

func New() *Game {
	// Build the GUI window for the game
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Asteroid Shooter")
	rl.SetTargetFPS(60)

	// Build game
	g := &Game{
		texBackground:     rl.LoadTexture("resources/space_background.png"),
		Player:            player.New(config.ScreenWidth/2, config.ScreenHeight/2),
		shots:             make([]player.Shot, config.MaxShots),
		initialAsteroids:  5,
		astroids:          astroid.CreateMultipleAstroids(5),
		gameOver:          false,
		victory:           false,
		paused:            false,
		astriodsDestroyed: 0,
	}

	// Set shots to inactive
	for i := range g.shots {
		g.shots[i].Active = false
	}

	return g
}

func (g *Game) Update() {
	// If there are no asteroids left, we in
	if len(g.astroids) == 0 {
		g.victory = true
	}

	// Toggle paused
	if rl.IsKeyDown('P') {
		g.paused = !g.paused
	}

	// Restart the game
	if (g.gameOver || g.victory) && rl.IsKeyPressed('R') {
		g = New()
	}

	// If it is not game over, update the frame
	if !g.paused && !g.victory && !g.gameOver {
		// Update the player
		g.Player.Update()

		// Update the astroids
		for i := range g.astroids {
			g.astroids[i].Update()
		}

		// Update the shots
		for i := range g.shots {
			g.shots[i].Update()
		}

		// Fire the lasers
		if rl.IsKeyPressed(rl.KeySpace) {
			g.fireShot()
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

		// Check for a collision between shots and the asteroid
		for j := range g.shots {
			// Loop through all the active shots
			if g.shots[j].Active {
				// If it has collided with an asteroid
				if rl.CheckCollisionCircles(
					g.shots[j].Position,
					g.shots[j].Radius,
					g.astroids[i].Position,
					g.astroids[i].Size.X/2,
				) {
					// Destroy the shot and split the asteroid
					g.shots[j].Active = false

					// The asteroid shot split according to our rules
					newAstroids := astroid.SplitAsteroid(g.astroids[i])
					g.astroids = append(g.astroids, newAstroids...)

					// Remove the original asteroid from the slice
					g.astroids = append(g.astroids[:i], g.astroids[i+1:]...)

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

	// Draw astroids
	for i := range g.astroids {
		g.astroids[i].Draw()
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
