package game

import (
	"fmt"

	"github.com/artromone/ginsengine/internal/core"
	"github.com/artromone/ginsengine/internal/scenes"
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
	g.currentScene = scenes.NewTitleScene(width, height)
	g.currentScene.OnEnter()
	return g
}

func (g *Game) Update() error {
	if g.nextScene != nil {
		g.currentScene.OnExit()
		g.currentScene = g.nextScene
		g.currentScene.OnEnter()
		g.nextScene = nil
	}

	nextScene, err := g.currentScene.Update()
	if err != nil {
		return fmt.Errorf("scene update error: %w", err)
	}

	if nextScene != nil {
		g.nextScene = nextScene
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
