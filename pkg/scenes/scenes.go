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
	Steel
	Grass
	// TODO: 补充地形
)

type Scenes struct {
	Position *vector2.Vector2
	Image    *ebiten.Image
	Index     ScenesType
}

var globalScenes = make(map[ScenesType]*Scenes)

var Key ScenesType = 1

func Init() {
	// TODO: 遍历地图，生成各个块，加到 globalScenes 里面
	for y, line := range defMap {
		for x, t := range line {
			Key ++
			var position *vector2.Vector2 = vector2.New(x*60,y*60)
			globalScenes[Key]=New(position,t)
		}
	}
}

func New(position *vector2.Vector2, t ScenesType) *Scenes {
	// TODO: 根据类型使用不同的图片
	switch t {
	case Brick: return &Scenes{
					Position: position,
					Image:    scenes.BrickImage,
					Index:    t,
				}
	case Steel: return &Scenes{
					Position: position,
					Image:    scenes.SteelImage,
					Index:    t,
				}
	case Grass: return &Scenes{
					Position: position,
					Image:    scenes.GrassImage,
					Index:    t,
				}
    }
	return &Scenes{
		Position: position,
		Image:    scenes.BrickImage,
		Index:    t,
	}

}

func (s *Scenes) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(s.Position.ValueFloat64())
	screen.DrawImage(s.Image, opt)
}

func Draw(screen *ebiten.Image) {
	for _ , scenes := range globalScenes {
		if scenes.Index != Space{
		scenes.Draw(screen)
	    }
	}
}
