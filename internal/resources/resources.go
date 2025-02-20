package resources

import (
	"image/color"
	"log"
	"os"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type ResourceManager struct {
	fontCache sync.Map
	imgCache  sync.Map
}

var (
	instance *ResourceManager
	once     sync.Once
)

func GetInstance() *ResourceManager {
	once.Do(func() {
		instance = &ResourceManager{}
	})
	return instance
}

func (rm *ResourceManager) LoadFont(path string, size float64) font.Face {
	key := path + ":" + string(rune(size))
	if v, ok := rm.fontCache.Load(key); ok {
		return v.(font.Face)
	}

	fontBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to load font %s: %v", path, err)
	}

	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create font face: %v", err)
	}

	rm.fontCache.Store(key, face)
	return face
}

func (rm *ResourceManager) RenderText(face font.Face, str string, clr color.Color) *ebiten.Image {
	bounds := text.BoundString(face, str)
	w := bounds.Dx()
	h := bounds.Dy()

	img := ebiten.NewImage(w, h)
	text.Draw(img, str, face, -bounds.Min.X, -bounds.Min.Y, clr)
	return img
}
