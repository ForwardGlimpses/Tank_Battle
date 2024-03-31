package scenes

import (
	"github.com/ForwardGlimpses/Tank_Battle/assets/scenes"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/vector2"
	"github.com/hajimehoshi/ebiten/v2"
)

type ScenesType int

const (
	// Space 不使用，仅作为地图空白标识
	Space ScenesType = iota
	Brick
	// TODO: 补充地形
)

type Scenes struct {
	Position *vector2.Vector2
	Image    *ebiten.Image
}

var globalScenes = make(map[*Scenes]*Scenes)

func Init() {
	// TODO: 遍历地图，生成各个块，加到 globalScenes 里面
	// for y, line := range defMap {
	// 	for x, t := range line {

	// 	}
	// }
}

func New(position *vector2.Vector2, t ScenesType) *Scenes {
	// TODO: 根据类型使用不同的图片
	return &Scenes{
		Position: position,
		Image:    scenes.ScenesImage,
	}
}

func (s *Scenes) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(s.Position.ValueFloat64())
	screen.DrawImage(s.Image, opt)
}

func Draw(screen *ebiten.Image) {
	for _, scenes := range globalScenes {
		scenes.Draw(screen)
	}
}
