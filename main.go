package main

import "asteroids/internal/game"

// main
func main() {
	g := game.New()
	defer g.Close()

	// Game loop
	for !g.WindowShouldClose() {
		g.Update()
		g.Draw()
	}
}
