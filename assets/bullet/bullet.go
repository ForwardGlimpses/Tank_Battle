package bullet

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed bullet.png
	Bullet_png  []byte
	BulletImage *ebiten.Image
)

func init() {
	temp, _, _ := image.Decode(bytes.NewReader(Bullet_png))
	BulletImage = ebiten.NewImageFromImage(temp)
}

