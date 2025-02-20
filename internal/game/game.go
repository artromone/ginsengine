package game

import (
	"github.com/artromone/ginsengine/internal/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	currentScene Scene
}

type Scene interface {
	Update() Scene
	Draw(screen *ebiten.Image)
}

func NewGame() *Game {
	return &Game{
		currentScene: scenes.NewTitleScene(),
	}
}

func (g *Game) Update() error {
	if nextScene := g.currentScene.Update(); nextScene != nil {
		g.currentScene = nextScene
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1280, 720
}
