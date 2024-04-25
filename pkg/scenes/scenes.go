package scenes

import (
	"github.com/ForwardGlimpses/Tank_Battle/assets/scenes"
	//"github.com/ForwardGlimpses/Tank_Battle/pkg/vector2"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
)

type ScenesType int

const (
	Space ScenesType = iota
	Brick
	Steel
	Grass
)

type Scenes struct {
	Position resolv.Vector
	Image    *ebiten.Image
	index    int
	Type     ScenesType
}

var (
	globalScenes = make(map[int]*Scenes)
	scenesImages = []*ebiten.Image{nil, scenes.BrickImage, scenes.SteelImage, scenes.GrassImage}
	Key int = 0
)

func Init() {
	for y, line := range defMap {
		for x, t := range line {
			var position resolv.Vector = resolv.NewVector(float64(x*60), float64(y*60))
			if t != Space {
				ins := New(position,t)
				globalScenes[ins.index] = ins
			}
		}
	}
}

func New(position resolv.Vector, t ScenesType) *Scenes {
	Key++
	return &Scenes{
		Position: position,
		Image:    scenesImages[t],
		index:    Key,
		Type:     t,
	}
}

func (s *Scenes) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(s.Position.X,s.Position.Y)
	screen.DrawImage(s.Image, opt)
}

func Draw(screen *ebiten.Image) {
	for _, scenes := range globalScenes {
		scenes.Draw(screen)
	}
}
