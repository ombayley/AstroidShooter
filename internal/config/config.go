package config

// Basic struct contaning row and col
type frame struct{ Row, Col int }

// Sprite map for player selection
var PlayerSpriteMap = map[int]frame{
	1:  {0, 0},
	2:  {0, 1},
	3:  {0, 2},
	4:  {0, 3},
	5:  {0, 4},
	6:  {0, 5},
	7:  {0, 6},
	8:  {0, 7},
	9:  {1, 0},
	10: {1, 1},
	11: {1, 2},
	12: {1, 3},
	13: {1, 4},
	14: {1, 5},
	15: {1, 6},
	16: {1, 7},
	17: {2, 0},
	18: {2, 1},
	19: {2, 2},
	20: {2, 3},
	21: {2, 4},
	22: {2, 5},
	23: {2, 6},
	24: {2, 7},
}

// Sprite map for boost selection
var BoostSpriteMap = map[int]frame{
	1: {5, 6},
	2: {5, 7},
}

// Vars detrmining player maximums
var (
	RotationSpeed = float32(2.0)
	PlayerSpeed   = float32(6.0)
	ShotSpeed     = float32(8.0)
	MaxShots      = int(10)
)

// Source for all constants used in-game
const (
	ScreenWidth  = 1600
	ScreenHeight = 800
	TileSize     = 64
)
