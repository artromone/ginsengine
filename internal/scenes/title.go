package scenes

import (
	"image/color"

	"github.com/artromone/ginsengine/internal/core"
	"github.com/artromone/ginsengine/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
)

type TitleScene struct {
	core.BaseScene
	titleFont font.Face
	titleText *ebiten.Image
}

func NewTitleScene(width, height int) core.Scene {
	s := &TitleScene{
		BaseScene: core.NewBaseScene(width, height),
	}
	return s
}

func (s *TitleScene) OnEnter() {
	rm := resources.GetInstance()
	s.titleFont = rm.LoadFont("assets/fonts/pressstart2p.ttf", 24)
	s.titleText = rm.RenderText(s.titleFont, "Press ENTER to start", color.White)
}

func (s *TitleScene) Update() (core.Scene, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return NewGameScene(s.Width, s.Height), nil
	}
	return nil, nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 100, 255})

	op := &ebiten.DrawImageOptions{}
	x := float64(s.Width-s.titleText.Bounds().Dx()) / 2
	y := float64(s.Height-s.titleText.Bounds().Dy()) / 2
	op.GeoM.Translate(x, y)
	screen.DrawImage(s.titleText, op)
}
