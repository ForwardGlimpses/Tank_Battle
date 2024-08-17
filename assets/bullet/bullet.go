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
	Bullet1_png []byte
	//go:embed bullet2.png
	Bullet2_png []byte
	BulletImage = map[string]*ebiten.Image{}
)

func init() {
	bulletPlayer, _, _ := image.Decode(bytes.NewReader(Bullet1_png))
	BulletPlayer := ebiten.NewImageFromImage(bulletPlayer)
	bulletEnemy, _, _ := image.Decode(bytes.NewReader(Bullet2_png))
	BulletEnemy := ebiten.NewImageFromImage(bulletEnemy)
	BulletImage = map[string]*ebiten.Image{
		"Player": BulletPlayer,
		"NPC":    BulletEnemy,
	}
}
