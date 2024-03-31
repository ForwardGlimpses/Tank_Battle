package scenes

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	// 导入的图片名应该和变量名对应，尽量和 scenes 代码中的类型也对上
	//go:embed scenes3.png
	Scenes1_png []byte
	ScenesImage *ebiten.Image
)

func init() {
	temp, _, _ := image.Decode(bytes.NewReader(Scenes1_png))
	ScenesImage = ebiten.NewImageFromImage(temp)
}
