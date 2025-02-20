package core

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type Scene struct {
	background    *Background
	dialogueLines []DialogueLine
	currentLine   int
	characters    []Character
	textFont      font.Face
	nameFont      font.Face
}

func NewScene() *Scene {
	scene := &Scene{}
	return scene
}

func (s *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
}
