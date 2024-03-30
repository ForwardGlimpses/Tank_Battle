package bullet

import (
	"github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/vector2"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	Position  *vector2.Vector2
	Direction direction.Direction
	Speed     *vector2.Vector2
	Image     *ebiten.Image
}

func (b *Bullet) Update() {
	b.Position = b.Position.Add(b.Speed)
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(b.Position.ValueFloat64())
	screen.DrawImage(b.Image, opt)
}

var step int = 5

// 全局子弹列表
var globalBullets []*Bullet

func Update() {
	for _, bullet := range globalBullets {
		bullet.Update()
	}
}

func Draw(screen *ebiten.Image) {
	for _, bullet := range globalBullets {
		bullet.Draw(screen)
	}
}

// TODO: 还需要伤害，创建者之类的信息
type CreateOption struct {
	Position *vector2.Vector2
	// Speed     *vector2.Vector2    // 子弹速度如果都相同，可以通过方向计算出来
	Direction direction.Direction
}

func Create(opt *CreateOption) {
	bullet := &Bullet{
		Position:  opt.Position,
		Direction: opt.Direction,
		Speed:     opt.Direction.DirectionVector2().MulScalar(step),
		Image:     bullet.BulletImage,
	}
	//  TODO: 设置碰撞器

	globalBullets = append(globalBullets, bullet)
}