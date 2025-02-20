package core

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update() (Scene, error)
	Draw(screen *ebiten.Image)
	Layout() (int, int)
	OnEnter()
	OnExit()
}
