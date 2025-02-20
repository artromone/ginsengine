package main

import (
	"log"

	"github.com/artromone/ginsengine/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.NewGame(1280, 720)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
