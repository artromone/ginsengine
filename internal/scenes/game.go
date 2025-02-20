package scenes

import (
	"fmt"
	"image/color"

	"github.com/artromone/ginsengine/internal/core"
	"github.com/artromone/ginsengine/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
)

type GameScene struct {
	core.BaseScene
	gameFont  font.Face
	scoreText *ebiten.Image
	score     int
}

func NewGameScene(width, height int) core.Scene {
	s := &GameScene{
		BaseScene: core.NewBaseScene(width, height),
		score:     0,
	}
	return s
}

func (s *GameScene) OnEnter() {
	rm := resources.GetInstance()
	s.gameFont = rm.LoadFont("assets/fonts/pressstart2p.ttf", 20)
	s.updateScoreText()
}

func (s *GameScene) updateScoreText() {
	rm := resources.GetInstance()
	s.scoreText = rm.RenderText(s.gameFont, fmt.Sprintf("Score: %d", s.score), color.White)
}

func (s *GameScene) Update() (core.Scene, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return NewTitleScene(s.Width, s.Height), nil
	}

	// Example game logic: increase score when space is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.score++
		s.updateScoreText()
	}

	return nil, nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 100, 0, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, 10) // Position score in top-left corner
	screen.DrawImage(s.scoreText, op)
}
