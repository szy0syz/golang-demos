package apptype

import (
	"fyne.io/fyne/v2"
	"image/color"
)

type BrushType = int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}

type State struct {
	BrushColor     color.Color
	BrushType      int
	FilePath       string
	SwatchSelected int
}

// SetFilePath 相当于原来我们给类的实力赋值
func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
