package core

type BaseScene struct {
	Width  int
	Height int
}

func NewBaseScene(width, height int) BaseScene {
	return BaseScene{
		Width:  width,
		Height: height,
	}
}

func (b *BaseScene) Layout() (int, int) {
	return b.Width, b.Height
}

func (b *BaseScene) OnEnter() {}
func (b *BaseScene) OnExit()  {}
