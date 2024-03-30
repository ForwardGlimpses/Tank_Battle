package scenes

import (

	"github.com/ForwardGlimpses/Tank_Battle/assets/scenes"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/vector2"
	"github.com/hajimehoshi/ebiten/v2"
)


type Scenes struct {
	Position *vector2.Vector2
	Image *ebiten.Image
}

var globalScenes = make(map[*Scenes]*Scenes)

func New() *Scenes{
	return &Scenes{
		Position: vector2.New(100,100),
		Image: scenes.ScenesImage,
	}
}

func (s *Scenes) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(s.Position.ValueFloat64())
	screen.DrawImage(s.Image, opt)
}

func Draw(screen *ebiten.Image){
	for _, scenes := range globalScenes{
		scenes.Draw(screen)
	}
}