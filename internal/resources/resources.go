package resources

import (
	"image"
	"image/color"
	"log"
	"sync"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var fontCache sync.Map

func LoadFont(path string, size float64) font.Face {
	key := path + ":" + string(size)
	if v, ok := fontCache.Load(key); ok {
		return v.(font.Face)
	}

	tt, err := text.OpenFont(path)
	if err != nil {
		log.Fatal(err)
	}

	face, err := text.NewGoTextFace(tt, text.GoTextFaceOptions{
		Size: size,
	})
	if err != nil {
		log.Fatal(err)
	}

	fontCache.Store(key, face)
	return face
}

func RenderText(face font.Face, text string, clr color.Color) *ebiten.Image {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	text.Draw(img, text, face, image.Point{}, color.Black, text.AlignmentLeft)

	ebitenImage := ebiten.NewImageFromImage(img)
	return ebitenImage
}
