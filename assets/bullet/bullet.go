package bullet

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
)

var (
	//go:embed bullet0.png
	Bullet0_png   []byte
	BulletImage image.Image
)

func init() {
	BulletImage, _, _ = image.Decode(bytes.NewReader(Bullet0_png))
}
