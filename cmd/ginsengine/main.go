package main

import (
	"log"

	"github.com/artromone/ginsengine/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 720
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Visual Novel Engine")

	game := game.NewGame(ScreenWidth, ScreenHeight)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
