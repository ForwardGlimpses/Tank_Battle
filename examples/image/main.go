package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	m image.Image
)

func main() {
	// 需要导入 image/png 解析 png 格式图片
	m, _, _ = image.Decode(bytes.NewReader(tank.Tank0_png))
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Options 可以设置绘制参数
	// 设置位置：op.GeoM.Translate(x, y)
	screen.DrawImage(ebiten.NewImageFromImage(m), &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(w, h int) (int, int) {
	return 100, 100
}
