package scenes

import (
	"github.com/ForwardGlimpses/Tank_Battle/assets/scenes"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
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
	Rivers
)

type Scenes struct {
	Collider *collision.Collider
	Image *ebiten.Image
	index int
	Type  ScenesType
	Hp int
}

var (
	globalScenes     = make(map[int]*Scenes)
	scenesImages     = []*ebiten.Image{nil, scenes.BrickImage, scenes.SteelImage, scenes.GrassImage,scenes.Rivers_Image}
	Key          int = 0
)

func init() {
	tankbattle.RegisterInit(Init,2)
	tankbattle.RegisterDraw(Draw,2)
}


func Init() error {
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
	return nil
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
		if t.Hp <= 0 {
			delete(globalScenes, t.index)
			t.Collider.Destruction()
	        t.Collider.Update()
		}
	}
}


func (t *Scenes) TankIsPassable() bool {
	if t.Type == Grass {
		return true
	}else{
		return false
	}
}

func (t *Scenes) BulletIsPassable() bool {
	if t.Type == Grass || t.Type == Rivers{
		return true
	}else{
		return false
	}
}
