package player

import (
	"asteroids/internal/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shot struct {
	Position rl.Vector2
	Speed    rl.Vector2
	Radius   float32
	Active   bool
}

func (s *Shot) Draw() {
	if s.Active {
		rl.DrawCircleV(s.Position, s.Radius, rl.Yellow)
	}
}
func (s *Shot) Update() {
	if s.Active {
		s.Position.X += s.Speed.X
		s.Position.Y -= s.Speed.Y
		if s.Position.X < 0 || s.Position.X > config.ScreenWidth || s.Position.Y < 0 || s.Position.Y > config.ScreenHeight {
			s.Active = false
		}
	}
}
