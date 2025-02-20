package main

import (
	"log"

	"github.com/artromone/ginsengine/internal/core"
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

	novelScene := core.NewNovelScene(ScreenWidth, ScreenHeight)
	err := novelScene.LoadScript("assets/scripts/story.txt")
	if err != nil {
		log.Fatalf("Failed to load script: %v", err)
	}

	game.SetScene(novelScene)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
