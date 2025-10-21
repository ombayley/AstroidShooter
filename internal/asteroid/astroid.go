package asteroid

import (
	//internal packages
	"asteroids/internal/config"
	"asteroids/internal/util"
	"sync"

	// external packages
	rl "github.com/gen2brain/raylib-go/raylib"
)

// --- Variables -----------------
var (
	tileSheet   rl.Texture2D
	asteroidSrc = rl.Rectangle{ // source rect on the tilesheet
		X:      float32(0 * config.TileSize),
		Y:      float32(4 * config.TileSize),
		Width:  float32(config.TileSize),
		Height: float32(config.TileSize),
	}
	loadOnce     sync.Once
	assetsLoaded bool
)

// Init loads package-local assets. Call after rl.InitWindow.
func Init() {
	loadOnce.Do(func() {
		tileSheet = rl.LoadTexture("resources/tilesheet.png")
		assetsLoaded = true
	})
}

// Shutdown frees package-local assets. Call before rl.CloseWindow.
func Shutdown() {
	if assetsLoaded {
		rl.UnloadTexture(tileSheet)
		assetsLoaded = false
	}
}

// --- Sizing -----------------

// type to store the size of the Asteroid
type AsteroidSize int

// Enum for Asteroid sizes
const (
	Large AsteroidSize = iota
	Medium
	Small
)

// Scale the image of the asteroid based on the asteroidSize
func (s AsteroidSize) scale() float32 {
	switch s {
	case Large:
		return 1.0
	case Medium:
		return 0.7
	case Small:
		return 0.4
	default:
		return 1.0
	}
}

// Get the size of the Asteroid
func (s AsteroidSize) SizeVec() rl.Vector2 {
	ts := float32(config.TileSize)
	k := s.scale()
	return rl.Vector2{X: ts * k, Y: ts * k}
}

// --- Asteroid -----------------

// Asteroid struct containing Asteroid state
type Asteroid struct {
	Position     rl.Vector2
	Speed        rl.Vector2
	Size         rl.Vector2
	asteroidSize AsteroidSize
}

// Draw the asteroid to the screen
func (a *Asteroid) Draw() {
	destTexture := rl.Rectangle{X: a.Position.X, Y: a.Position.Y, Width: a.Size.X, Height: a.Size.Y}
	rl.DrawTexturePro(
		tileSheet,
		asteroidSrc,
		destTexture,
		rl.Vector2{X: a.Size.X / 2, Y: a.Size.Y / 2},
		0.0,
		rl.White,
	)
}

// Update the asteroid position
func (a *Asteroid) Update() {
	// Move the asteroid in its direction
	a.Position = rl.Vector2Add(a.Position, a.Speed)

	// Wrap the position, so they are always on screen
	util.WrapPosition(&a.Position, a.Size.X, config.ScreenWidth, config.ScreenHeight)
}

// Split the asteroid into smaller asteroids
func Break(asteroid Asteroid) []Asteroid {
	// return nil if the asteroid is small (small asteroids dont generate more)
	if asteroid.asteroidSize == Small {
		return nil
	}

	// calculate # of splits to do
	var newSize AsteroidSize
	var split int
	switch asteroid.asteroidSize {
	case Large:
		newSize, split = Medium, 2
	case Medium:
		newSize, split = Small, 4
	default:
		return nil
	}

	// Create the new smaller asteroids
	children := make([]Asteroid, 0, split)
	for i := 0; i < split; i++ {
		// Generate a random direction to go
		angle := float32(rl.GetRandomValue(0, 360))
		direction := util.DirectionVector(angle)
		speed := rl.Vector2Scale(direction, 2.0)
		// Create the new asteroid and add it to the list
		child := createAsteroid(newSize, asteroid.Position, speed)
		children = append(children, child)
	}

	// return list of asteroid structs
	return children
}

// --- Spawn -----------------

// Create multiple starting Asteroids (large)
func CreateAsteroids(count int) []Asteroid {
	asteroids := make([]Asteroid, count)
	for i := range asteroids {
		asteroids[i] = createRandomLargeAsteroid()
	}
	return asteroids
}

// Create a single large Asteroid at a random position
func createRandomLargeAsteroid() Asteroid {

	// Generate a random edge of the screen to spawn
	randomEdge := rl.GetRandomValue(0, 3)
	var position rl.Vector2

	// Generate a random position on screen
	randomX := float32(rl.GetRandomValue(0, config.ScreenWidth))
	randomY := float32(rl.GetRandomValue(0, config.ScreenHeight))

	// Set the position based on the random edge given
	switch randomEdge {
	case 0: // top
		position = rl.Vector2{X: randomX, Y: -config.TileSize}
	case 1: // right
		position = rl.Vector2{X: config.ScreenWidth + config.TileSize, Y: randomY}
	case 2: // bottom
		position = rl.Vector2{X: randomX, Y: config.ScreenHeight + config.TileSize}
	case 3: //left
		position = rl.Vector2{X: -config.TileSize, Y: randomY}
	}

	// Generate a random speed and direction for the asteroid
	speed := rl.Vector2{
		X: float32(rl.GetRandomValue(-10, 10)) / 10,
		Y: float32(rl.GetRandomValue(-10, 10)) / 10,
	}

	// Create the asteroid
	return createAsteroid(Large, position, speed)
}

// Create a single Asteroid of specific size at a specific position
func createAsteroid(asteroidSize AsteroidSize, position, speed rl.Vector2) Asteroid {

	// Get the size of the asteroid from the asteroidSize enum
	size := asteroidSize.SizeVec()

	// Create the asteroid
	return Asteroid{
		Position:     position,
		Speed:        speed,
		Size:         size,
		asteroidSize: asteroidSize,
	}
}
