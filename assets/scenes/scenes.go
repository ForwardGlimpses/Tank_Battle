package scenes

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed brick.png
	brick_png []byte
	BrickImage *ebiten.Image
	//go:embed steel.png
	steel_png []byte
	SteelImage *ebiten.Image
	//go:embed grass.png
	grass_png []byte
	GrassImage *ebiten.Image
)

func init() {
	Brick, _, _ := image.Decode(bytes.NewReader(brick_png))
	BrickImage = ebiten.NewImageFromImage(Brick)
	Steel, _, _ := image.Decode(bytes.NewReader(steel_png))
	SteelImage = ebiten.NewImageFromImage(Steel)
	Grass, _, _ := image.Decode(bytes.NewReader(grass_png))
	GrassImage = ebiten.NewImageFromImage(Grass)
}
