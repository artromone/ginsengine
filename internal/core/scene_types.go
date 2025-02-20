package core

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png" // Needed for PNG support
	"os"

	"github.com/artromone/ginsengine/internal/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	DialogueBoxHeight  = 150
	DialogueBoxPadding = 20
	NameBoxWidth       = 200
	NameBoxHeight      = 40
)

// DialogueType represents different types of dialogue presentation
type DialogueType int

const (
	CharacterDialogue DialogueType = iota
	NarratorDialogue
	CenterScreenText
)

// Character represents a character in the visual novel
type Character struct {
	Name      string
	ImagePath string
}

// DialogueLine represents a single line of dialogue
type DialogueLine struct {
	Type      DialogueType
	Character *Character // nil for narrator or center text
	Text      string
	Position  Position // position of character sprite
}

// Position represents character sprite position
type Position struct {
	X float64
	Y float64
}

// Background represents scene background
type Background struct {
	ImagePath string
}

// Добавляем структуру для хранения загруженных изображений
type ImageCache struct {
	characters  map[string]*ebiten.Image
	backgrounds map[string]*ebiten.Image
}

// Расширяем NovelScene
type NovelScene struct {
	BaseScene
	background    *Background
	dialogueLines []DialogueLine
	currentLine   int
	characters    map[string]*Character
	textFont      font.Face
	nameFont      font.Face
	imageCache    ImageCache
	dialogueBox   *ebiten.Image
	nameBox       *ebiten.Image
}

func NewNovelScene(width, height int) *NovelScene {
	scene := &NovelScene{
		BaseScene:   NewBaseScene(width, height),
		characters:  make(map[string]*Character),
		currentLine: 0,
		imageCache: ImageCache{
			characters:  make(map[string]*ebiten.Image),
			backgrounds: make(map[string]*ebiten.Image),
		},
	}

	// Инициализация шрифтов
	rm := resources.GetInstance()
	scene.textFont = rm.LoadFont("assets/fonts/pressstart2p.ttf", 20)
	scene.nameFont = rm.LoadFont("assets/fonts/pressstart2p.ttf", 24)

	// Создание диалогового окна и окна имени
	scene.createDialogueBoxes()

	return scene
}

func (s *NovelScene) createDialogueBoxes() {
	// Создаем диалоговое окно
	s.dialogueBox = ebiten.NewImage(s.Width, DialogueBoxHeight)
	s.dialogueBox.Fill(color.RGBA{0, 0, 0, 200})

	// Создаем окно имени
	s.nameBox = ebiten.NewImage(NameBoxWidth, NameBoxHeight)
	s.nameBox.Fill(color.RGBA{0, 0, 0, 230})
}

func (s *NovelScene) loadImage(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open image %s: %v", path, err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image %s: %v", path, err)
	}

	return ebiten.NewImageFromImage(img), nil
}

func (s *NovelScene) LoadResources() error {
	// Загрузка фона
	if s.background != nil {
		bgImage, err := s.loadImage(s.background.ImagePath)
		if err != nil {
			return err
		}
		s.imageCache.backgrounds[s.background.ImagePath] = bgImage
	}

	// Загрузка спрайтов персонажей
	for _, char := range s.characters {
		if _, exists := s.imageCache.characters[char.ImagePath]; !exists {
			charImage, err := s.loadImage(char.ImagePath)
			if err != nil {
				return err
			}
			s.imageCache.characters[char.ImagePath] = charImage
		}
	}

	return nil
}

func (s *NovelScene) Update() (Scene, error) {
	if s.currentLine >= len(s.dialogueLines) {
		// Ждем явного подтверждения для возврата в меню
		// if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// 	return NewTitleScene(s.Width, s.Height), nil
		// }
		return nil, nil
	}

	// Обработка перехода между диалогами
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.currentLine++
	}
	return nil, nil
}

func (s *NovelScene) Draw(screen *ebiten.Image) {
	// Отрисовка фона
	if bg, exists := s.imageCache.backgrounds[s.background.ImagePath]; exists {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(bg, op)
	}

	if s.currentLine >= len(s.dialogueLines) {
		return
	}

	line := s.dialogueLines[s.currentLine]

	// Отрисовка спрайта персонажа
	if line.Character != nil {
		if sprite, exists := s.imageCache.characters[line.Character.ImagePath]; exists {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(line.Position.X, line.Position.Y)
			screen.DrawImage(sprite, op)
		}
	}

	// Отрисовка диалога в зависимости от типа
	switch line.Type {
	case CharacterDialogue:
		s.drawCharacterDialogue(screen, line)
	case NarratorDialogue:
		s.drawNarratorDialogue(screen, line)
	case CenterScreenText:
		s.drawCenterText(screen, line)
	}
}

func (s *NovelScene) drawCharacterDialogue(screen *ebiten.Image, line DialogueLine) {
	// Отрисовка окна диалога
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(s.Height-DialogueBoxHeight))
	screen.DrawImage(s.dialogueBox, op)

	// Отрисовка имени персонажа
	if line.Character != nil {
		nameOp := &ebiten.DrawImageOptions{}
		nameOp.GeoM.Translate(DialogueBoxPadding, float64(s.Height-DialogueBoxHeight-NameBoxHeight))
		screen.DrawImage(s.nameBox, nameOp)

		text.Draw(screen,
			line.Character.Name,
			s.nameFont,
			DialogueBoxPadding+10,
			s.Height-DialogueBoxHeight-NameBoxHeight/2,
			color.White)
	}

	// Отрисовка текста диалога
	text.Draw(screen,
		line.Text,
		s.textFont,
		DialogueBoxPadding,
		s.Height-DialogueBoxHeight+DialogueBoxPadding,
		color.White)
}

func (s *NovelScene) drawNarratorDialogue(screen *ebiten.Image, line DialogueLine) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(s.Height-DialogueBoxHeight))
	screen.DrawImage(s.dialogueBox, op)

	text.Draw(screen,
		line.Text,
		s.textFont,
		DialogueBoxPadding,
		s.Height-DialogueBoxHeight+DialogueBoxPadding,
		color.White)
}

func (s *NovelScene) drawCenterText(screen *ebiten.Image, line DialogueLine) {
	bounds := text.BoundString(s.textFont, line.Text)
	x := (s.Width - bounds.Dx()) / 2
	y := s.Height / 2

	text.Draw(screen,
		line.Text,
		s.textFont,
		x,
		y,
		color.White)
}

func (s *NovelScene) OnEnter() {
	err := s.LoadResources()
	if err != nil {
		// В реальном приложении здесь должна быть proper обработка ошибок
		panic(err)
	}
}

func (s *NovelScene) OnExit() {
	// Очистка ресурсов при необходимости
}
