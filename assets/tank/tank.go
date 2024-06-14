package tank

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
)

var (
	//go:embed tank0.png
	Tank0_png   []byte
	//go:embed tank1.png
	Tank1_png   []byte
	PlayerImage image.Image
	EnemyImage  image.Image
	TankImage = make(map[string]image.Image)
)
func init() {
	PlayerImage, _, _ = image.Decode(bytes.NewReader(Tank0_png))
	EnemyImage, _, _ = image.Decode(bytes.NewReader(Tank1_png))
	TankImage =map[string]image.Image{
		"Player":PlayerImage,
		"NPC"   :EnemyImage,
	}
}
