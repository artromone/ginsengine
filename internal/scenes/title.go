package scenes

import (
	"image/color"

	"github.com/artromone/ginsengine/internal/game"
	"github.com/artromone/ginsengine/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
)

type TitleScene struct {
	titleFont font.Face
}

func NewTitleScene() game.Scene {
	return &TitleScene{
		titleFont: resources.LoadFont("fonts/pressstart2p.ttf", 24),
	}
}

func (s *TitleScene) Update() game.Scene {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return NewDialogScene()
	}
	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 100, 255})

	text := "Press ENTER to start"
	bounds := s.titleFont.Metrics()
	textWidth := font.MeasureString(s.titleFont, text)

	x := (1280 - textWidth.Round()) / 2
	y := 360 + bounds.Height.Round()/2

	textDraw := ebiten.NewImageFromImage(resources.RenderText(s.titleFont, text, color.White))

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(textDraw, op)
}
