package player

import (
	// internal packages
	"asteroids/internal/config"
	// external packages
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Shot struct for player shots
type Shot struct {
	Position rl.Vector2
	Speed    rl.Vector2
	Radius   float32
	Active   bool
}

// Draw the shot to screen
func (s *Shot) Draw() {
	if s.Active {
		rl.DrawCircleV(s.Position, s.Radius, rl.Yellow)
	}
}

// Update the shot position
func (s *Shot) Update() {
	if !s.Active {
		return
	}
	s.Position.X += s.Speed.X
	s.Position.Y -= s.Speed.Y
	if s.Position.X < 0 || s.Position.X > config.ScreenWidth || s.Position.Y < 0 || s.Position.Y > config.ScreenHeight {
		s.Active = false
	}
}
