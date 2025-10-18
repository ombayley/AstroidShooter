package astroid

import (
	"asteroids/internal/config"
	"asteroids/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Enum for astroid sizes
type AstroidSize int

const (
	Large AstroidSize = iota
	Medium
	Small
)

type Astroid struct {
	Position     rl.Vector2
	Speed        rl.Vector2
	Size         rl.Vector2
	texTilesheet rl.Texture2D
	asteroidRec  rl.Rectangle
}

func (a *Astroid) Draw() {
	// Draw the asteroid to the screen
	destTexture := rl.Rectangle{X: a.Position.X, Y: a.Position.Y, Width: a.Size.X, Height: a.Size.Y}
	rl.DrawTexturePro(
		a.texTilesheet,
		a.asteroidRec,
		destTexture,
		rl.Vector2{X: a.Size.X / 2, Y: a.Size.Y / 2},
		0.0,
		rl.White,
	)
}

func (a *Astroid) Update() {
	// Move the asteroid in its direction
	a.Position = rl.Vector2Add(a.Position, a.Speed)

	// Wrap the position, so they are always on screen
	util.WrapPosition(&a.Position, a.Size.X, config.ScreenWidth, config.ScreenHeight)
}

// Asteroid helper functions
func createLargeAsteroid() Astroid {

	// Generate a random edge of the screen to spawn
	randomEdge := rl.GetRandomValue(0, 3)
	var position rl.Vector2

	// Generate a random position on screen
	randomX := float32(rl.GetRandomValue(0, config.ScreenWidth))
	randomY := float32(rl.GetRandomValue(0, config.ScreenHeight))

	switch randomEdge {
	case 0:
		position = rl.Vector2{X: randomX, Y: +config.TileSize}
	case 1:
		position = rl.Vector2{X: config.ScreenWidth + config.TileSize, Y: randomY}
	case 2:
		position = rl.Vector2{X: randomX, Y: config.ScreenHeight + config.TileSize}
	case 3:
		position = rl.Vector2{X: -config.TileSize, Y: randomY}
	}

	// Generate a random speed and direction for the asteroid
	speed := rl.Vector2{
		X: float32(rl.GetRandomValue(-10, 10)) / 10,
		Y: float32(rl.GetRandomValue(-10, 10)) / 10,
	}

	// Create the large asteroid
	return createAsteroid(Large, position, speed)
}

func createAsteroid(asteroidSize AstroidSize, position, speed rl.Vector2) Astroid {

	// Scale the image of the asteroid based on the asteroidSize
	var size rl.Vector2
	switch asteroidSize {
	case Large:
		size = rl.Vector2{X: config.TileSize * 1.0, Y: config.TileSize * 1.0}
	case Medium:
		size = rl.Vector2{X: config.TileSize * 0.7, Y: config.TileSize * 0.7}
	case Small:
		size = rl.Vector2{X: config.TileSize * 0.4, Y: config.TileSize * 0.4}
	}

	// Create the asteroid
	return Astroid{
		Position:     position,
		Speed:        speed,
		Size:         size,
		texTilesheet: rl.LoadTexture("resources/tilesheet.png"),
		asteroidRec: rl.Rectangle{
			X:      0 * config.TileSize,
			Y:      4 * config.TileSize,
			Width:  config.TileSize,
			Height: config.TileSize,
		},
	}
}

func CreateMultipleAstroids(count int) []Astroid {
	astroids := make([]Astroid, count)
	for i := range astroids {
		astroids[i] = createLargeAsteroid()
	}
	return astroids
}
