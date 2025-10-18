package main

import "asteroids/internal/game"

func main() {
	g := game.New()
	defer g.Close()

	for !gorlWindowShouldClose(g) { // helper to avoid importing rl in main
		g.Update()
		g.Draw()
	}
}

func gorlWindowShouldClose(game *game.Game) bool {
	return game.WindowShouldClose()
}
