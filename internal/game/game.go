package game

import (
	"github.com/artromone/ginsengine/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	currentScene core.Scene
	nextScene    core.Scene
	width        int
	height       int
}

func NewGame(width, height int) *Game {
	g := &Game{
		width:  width,
		height: height,
	}
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
