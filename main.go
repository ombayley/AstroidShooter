package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1600
	screenHeight = 800
)

var (
	texBackground rl.Texture2D
)

func init() {
	// Initialize raylib window
	rl.InitWindow(screenWidth, screenHeight, "Astroid Shooter")
	rl.SetTargetFPS(60)

	// Load textures
	texBackground = rl.LoadTexture("resources/space_background.png")
}

func draw() {
	// Clear the screen
	rl.BeginDrawing()

	// Set the background to a nebula
	bgSource := rl.Rectangle{X: 0, Y: 0, Width: float32(texBackground.Width), Height: float32(texBackground.Height)}
	bgDest := rl.Rectangle{X: 0, Y: 0, Width: screenWidth, Height: screenHeight}
	rl.DrawTexturePro(texBackground, bgSource, bgDest, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

	//Set the background color rto black
	// rl.ClearBackground(rl.Black)

	// Draw score on the screen
	rl.DrawText("Score: 0", 10, 10, 20, rl.Gray)

	// End
	rl.EndDrawing()
}

func update() {
	// Update game logic

	// End
}

func deinit() {
	// Deinitialize raylib window
	rl.CloseWindow()

	// Unload textures when the game closes
	rl.UnloadTexture(texBackground)
}

func main() {
	defer deinit()

	for !rl.WindowShouldClose() {
		draw()
		update()
	}
}
