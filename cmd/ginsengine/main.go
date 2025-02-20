package main

import (
	"log"

	"github.com/artromone/ginsengine/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("ginsEngine")

	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
