package scenes

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed scenes0.png
	Scenes1_png  []byte
	ScenesImage *ebiten.Image
)

func init() {
	temp, _, _ := image.Decode(bytes.NewReader(Scenes1_png))
	ScenesImage = ebiten.NewImageFromImage(temp)
}

