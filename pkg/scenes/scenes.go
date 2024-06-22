package scenes

import (
	"github.com/ForwardGlimpses/Tank_Battle/assets/scenes"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/vector2"
	"github.com/hajimehoshi/ebiten/v2"
)

type ScenesType int

const (
	Space ScenesType = iota
	Brick
	Steel
	Grass
)

type Scenes struct {
	Collider *collision.Collider
	//Position vector2.Vector
	Image *ebiten.Image
	index int
	Type  ScenesType
	Hp int
}

var (
	globalScenes     = make(map[int]*Scenes)
	scenesImages     = []*ebiten.Image{nil, scenes.BrickImage, scenes.SteelImage, scenes.GrassImage}
	Key          int = 0
)

func Init() {
	for y, line := range defMap {
		for x, t := range line {
			var position vector2.Vector = vector2.NewVector(float64(x*60), float64(y*60))
			if t != Space {
				ins := New(position, t)
				ins.Collider.Data = ins
				globalScenes[ins.index] = ins
			}
		}
	}
}

func New(position vector2.Vector, t ScenesType) *Scenes {
	Key++
	return &Scenes{
		Collider: collision.NewCollider(position.X, position.Y, float64(scenesImages[t].Bounds().Dy()), float64(scenesImages[t].Bounds().Dx())),
		Image:    scenesImages[t],
		index:    Key,
		Type:     t,
		Hp:       100,
	}
}

func Delete(s *Scenes) {
	delete(globalScenes, s.index)
	s.Collider.Update()
}

func (s *Scenes) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(s.Collider.Position.X, s.Collider.Position.Y)
	screen.DrawImage(s.Image, opt)
}

func Draw(screen *ebiten.Image) {
	for _, scenes := range globalScenes {
		scenes.Draw(screen)
	}
}

func (t *Scenes) GetCamp() string {
	return ""
}

func (t *Scenes) TakeDamage(damage int) {
	if t.Type == Brick {
		t.Hp -= damage
	}
}

// 此函数暂时无逻辑，仅标识此结构为障碍物
func (t *Scenes) Obstacle() {

}
