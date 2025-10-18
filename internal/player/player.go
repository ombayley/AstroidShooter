package player

import (
	"asteroids/internal/config"
	"asteroids/internal/util"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func New(x, y float32) Player {

	// Set up the player
	p := Player{
		Position:     rl.Vector2{X: x, Y: y},
		Speed:        rl.Vector2{X: 0, Y: 0},
		Size:         rl.Vector2{X: config.TileSize, Y: config.TileSize},
		Rotation:     0,
		Acceleration: 0,
		texTilesheet: rl.LoadTexture("resources/tilesheet.png"),
	}
	p.SetSprite(2, 2)
	p.SetBoost(5, 7)

	return p
}

func (p *Player) SetSprite(row, col int) {
	ts := float32(config.TileSize)
	p.spriteRec = rl.Rectangle{
		X:     float32(col) * ts,
		Y:     float32(row) * ts,
		Width: ts, Height: ts,
	}
}

func (p *Player) SetBoost(row, col int) {
	ts := float32(config.TileSize)
	p.boostRec = rl.Rectangle{
		X:     float32(col) * ts,
		Y:     float32(row) * ts,
		Width: ts, Height: ts,
	}
}

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

func (p *Player) Draw() {
	dest := rl.Rectangle{X: p.Position.X, Y: p.Position.Y, Width: p.Size.X, Height: p.Size.Y}
	if p.IsBoosting {
		rl.DrawTexturePro(
			p.texTilesheet,
			p.boostRec,
			dest,
			rl.Vector2{X: p.Size.X / 2, Y: p.Size.Y/2 - 40},
			p.Rotation,
			rl.White,
		)
	}
	rl.DrawTexturePro(
		p.texTilesheet,
		p.spriteRec,
		dest,
		rl.Vector2{X: p.Size.X / 2, Y: p.Size.Y / 2},
		p.Rotation,
		rl.White,
	)
}
func (p *Player) Close() {
	rl.UnloadTexture(p.texTilesheet)
}
