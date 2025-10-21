package player

import (
	// built-in
	"sync"
	// internal packages
	"asteroids/internal/config"
	"asteroids/internal/util"

	// external packages
	rl "github.com/gen2brain/raylib-go/raylib"
)

// --- Package Variables -----------------

var (
	tileSheet    rl.Texture2D
	loadOnce     sync.Once
	assetsLoaded bool
)

// Init loads the player tilesheet. Call AFTER rl.InitWindow.
func Init() {
	loadOnce.Do(func() {
		tileSheet = rl.LoadTexture("resources/tilesheet.png")
		assetsLoaded = true
	})
}

// Shutdown frees the player tilesheet. Call BEFORE rl.CloseWindow.
func Shutdown() {
	if assetsLoaded {
		rl.UnloadTexture(tileSheet)
		assetsLoaded = false
	}
}

// --- Player -----------------

type Player struct {
	spriteRec    rl.Rectangle
	boostRec     rl.Rectangle
	texTilesheet rl.Texture2D
	Position     rl.Vector2
	Speed        rl.Vector2
	Size         rl.Vector2
	Acceleration float32
	Rotation     float32
	IsBoosting   bool
}

// New creates a new player - TODO add selector to set player, boost and stats
func New() Player {
	p := Player{
		Position:     rl.Vector2{X: config.ScreenWidth / 2, Y: config.ScreenHeight / 2},
		Speed:        rl.Vector2{X: 0, Y: 0},
		Size:         rl.Vector2{X: config.TileSize, Y: config.TileSize},
		Rotation:     0,
		Acceleration: 0,
	}
	p.setSprite(18)
	p.setBoost(1)
	setMaxmima()

	return p
}

// setSprite sets the sprite for the player (1-24)
func (p *Player) setSprite(sprite int) {
	if 1 > sprite || sprite > 24 {
		return // TODO set panic
	}

	frame := config.PlayerSpriteMap[sprite]
	ts := float32(config.TileSize)
	p.spriteRec = rl.Rectangle{
		X:     float32(frame.Col) * ts,
		Y:     float32(frame.Row) * ts,
		Width: ts, Height: ts,
	}
}

// setBoost sets the yellow (1) or purple (2) boost colour
func (p *Player) setBoost(sprite int) {
	if 1 > sprite || sprite > 2 {
		return // TODO set panic
	}

	frame := config.BoostSpriteMap[sprite]
	ts := float32(config.TileSize)
	p.boostRec = rl.Rectangle{
		X:     float32(frame.Col) * ts,
		Y:     float32(frame.Row) * ts,
		Width: ts, Height: ts,
	}
}

// SetMaxmima sets the maxmima for the player
func setMaxmima() {
	config.RotationSpeed = float32(2.0)
	config.PlayerSpeed = float32(6.0)
	config.ShotSpeed = float32(8.0)
	config.MaxShots = int(10)
}

// --- Action Functions -----------------

func (p *Player) Update() {
	// rotate player
	if rl.IsKeyDown(rl.KeyLeft) {
		p.Rotation -= config.RotationSpeed
	}
	if rl.IsKeyDown(rl.KeyRight) {
		p.Rotation += config.RotationSpeed
	}

	// default to not boosting
	p.IsBoosting = false

	// accel/decel player
	if rl.IsKeyDown(rl.KeyUp) {
		if p.Acceleration < 0.9 {
			p.Acceleration += 0.1
		}
		p.IsBoosting = true
	}
	if rl.IsKeyDown(rl.KeyDown) {
		if p.Acceleration > 0 {
			p.Acceleration -= 0.05
		}
		if p.Acceleration < 0 {
			p.Acceleration = 0
		}
	}

	// calculate move
	dir := util.DirectionVector(p.Rotation)
	p.Speed = rl.Vector2Scale(dir, config.PlayerSpeed)

	// move player
	p.Position.X += p.Speed.X * p.Acceleration
	p.Position.Y -= p.Speed.Y * p.Acceleration // screen Y grows downward

	util.WrapPosition(&p.Position, float32(config.TileSize), config.ScreenWidth, config.ScreenHeight)
}

// Draw the player
func (p *Player) Draw() {
	if !assetsLoaded {
		return // TODO set panic
	}

	dest := rl.Rectangle{X: p.Position.X, Y: p.Position.Y, Width: p.Size.X, Height: p.Size.Y}

	// Draw the boost
	if p.IsBoosting {
		// Slightly offset origin so boost draws "behind" the ship.
		rl.DrawTexturePro(
			tileSheet,
			p.boostRec,
			dest,
			rl.Vector2{X: p.Size.X / 2, Y: p.Size.Y/2 - 40},
			p.Rotation,
			rl.White,
		)
	}

	// Draw the ship
	rl.DrawTexturePro(
		tileSheet,
		p.spriteRec,
		dest,
		rl.Vector2{X: p.Size.X / 2, Y: p.Size.Y / 2},
		p.Rotation,
		rl.White,
	)
}
