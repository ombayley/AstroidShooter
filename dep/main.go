package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const (
	screenWidth   = 800
	screenHeight  = 400
	tileSize      = 64
	rotationSpeed = 2.0
	playerSpeed   = 6.0
)

var (
	texTilesheet  rl.Texture2D
	texBackground rl.Texture2D
	spriteRec     rl.Rectangle
	boostRec      rl.Rectangle
	player        Player
	spriteRow     int
	spriteCol     int
)

func init() {
	// Initialize raylib window
	rl.InitWindow(screenWidth, screenHeight, "Astroid Shooter")
	rl.SetTargetFPS(60)

	// Select Sprite
	setSprite(2, 2)

	// Load textures
	texTilesheet = rl.LoadTexture("resources/tilesheet.png")
	texBackground = rl.LoadTexture("resources/space_background.png")

	// init game
	initGame()
}

func (p *Player) Update() {
	// Player movement
	if rl.IsKeyDown(rl.KeyLeft) {
		p.rotation -= rotationSpeed
	}
	if rl.IsKeyDown(rl.KeyRight) {
		p.rotation += rotationSpeed
	}
	// Accelerate the player with up
	if rl.IsKeyDown(rl.KeyUp) {
		if p.acceleration < 0.9 {
			p.acceleration += 0.1
		}
	}
	// Decellerate the player with down
	if rl.IsKeyDown(rl.KeyDown) {
		if p.acceleration > 0 {
			p.acceleration -= 0.05
		}
		if p.acceleration < 0 {
			p.acceleration = 0
		}
	}
	// Get the direction the sprite is pointing
	direction := getDirectionVector(player.rotation)

	// Start to move to the direction
	player.speed = rl.Vector2Scale(direction, playerSpeed)

	// Accelerate in that direction
	player.position.X += player.speed.X * player.acceleration
	player.position.Y -= player.speed.Y * player.acceleration

	// To void losing our ship, we wrap around the screen
	wrapPosition(&p.position, tileSize)
}

func getDirectionVector(rotation float32) rl.Vector2 {
	// Convert the rotation to radians
	radians := float64(rotation) * rl.Deg2rad

	// Return the vector of the direction we are pointing at
	return rl.Vector2{
		X: float32(math.Sin(radians)),
		Y: float32(math.Cos(radians)),
	}
}

func wrapPosition(pos *rl.Vector2, objectSize float32) {
	// If we go off the left side of the screen
	if pos.X > screenWidth+objectSize {
		pos.X = -objectSize
	}
	// If we go off the right side of the screen
	if pos.X < -objectSize {
		pos.X = screenWidth + objectSize
	}
	// If we go off the bottom of the screen
	if pos.Y > screenHeight+objectSize {
		pos.Y = -objectSize
	}
	// If we go off the top of the screen
	if pos.Y < -objectSize {
		pos.Y = screenHeight + objectSize
	}
}

func setSprite(row, col int) {
	spriteRow = row
	spriteCol = col
	spriteRec = rl.Rectangle{X: float32(tileSize) * float32(spriteCol), Y: float32(tileSize) * float32(spriteRow), Width: float32(tileSize), Height: float32(tileSize)}
}

type Player struct {
	position     rl.Vector2
	speed        rl.Vector2
	size         rl.Vector2
	acceleration float32
	rotation     float32
	isBoosting   bool
}

func (p *Player) Draw() {
	destTexture := rl.Rectangle{X: p.position.X, Y: p.position.Y, Width: p.size.X, Height: p.size.Y}
	rl.DrawTexturePro(
		texTilesheet,
		spriteRec,
		destTexture,
		rl.Vector2{X: p.size.X / 2, Y: p.size.Y / 2},
		p.rotation,
		rl.White,
	)
}

func initGame() {
	player = Player{
		position:     rl.Vector2{X: 400, Y: 200},
		speed:        rl.Vector2{X: 0.0, Y: 0.0},
		size:         rl.Vector2{X: tileSize, Y: tileSize},
		rotation:     0.0,
		acceleration: 0.0,
		isBoosting:   false,
	}
}

func draw() {
	// Clear the screen
	rl.BeginDrawing()

	// Set the background to a nebula
	bgSource := rl.Rectangle{X: 0, Y: 0, Width: float32(texBackground.Width), Height: float32(texBackground.Height)}
	bgDest := rl.Rectangle{X: 0, Y: 0, Width: screenWidth, Height: screenHeight}
	rl.DrawTexturePro(texBackground, bgSource, bgDest, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

	//Draw the player
	player.Draw()

	// Draw score on the screen
	rl.DrawText("Score: 0", 10, 10, 20, rl.Gray)

	// End
	rl.EndDrawing()
}

func update() {
	// Update game logic
	player.Update()
	// End
}

func deinit() {
	// Deinitialize raylib window
	rl.CloseWindow()

	// Unload textures when the game closes
	rl.UnloadTexture(texTilesheet)
	rl.UnloadTexture(texBackground)
}

func main() {
	defer deinit()

	for !rl.WindowShouldClose() {
		draw()
		update()
	}
}
