package util

import (
	// built-in
	"math"

	// external packages
	rl "github.com/gen2brain/raylib-go/raylib"
)

// DirectionVector returns a unit vector in the direction of the rotation
func DirectionVector(rotationDeg float32) rl.Vector2 {
	r := float64(rotationDeg) * float64(rl.Deg2rad)
	return rl.Vector2{
		X: float32(math.Sin(r)),
		Y: float32(math.Cos(r)),
	}
}

// WrapPosition wraps the position of an object around the screen
func WrapPosition(pos *rl.Vector2, objectSize float32, screenW, screenH float32) {
	if pos.X > screenW+objectSize {
		pos.X = -objectSize
	}
	if pos.X < -objectSize {
		pos.X = screenW + objectSize
	}
	if pos.Y > screenH+objectSize {
		pos.Y = -objectSize
	}
	if pos.Y < -objectSize {
		pos.Y = screenH + objectSize
	}
}
