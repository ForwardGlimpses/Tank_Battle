package bullet

import (
	"github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/scenes"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/types"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/vector2"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	Collider  *collision.Collider
	Direction direction.Direction
	Speed     vector2.Vector
	Image     *ebiten.Image
	Index     int
}

func (b *Bullet) Update() {

	dx := b.Speed.X
	dy := b.Speed.Y

	// 检测 x 轴是否碰撞，如果碰撞将 x 轴速度反向，下面的 y 轴处理同理
	if check := b.Collider.Check(dx, dy); check != nil {

		// 打印发生碰撞的小球编号
		for _, obj := range check.Colliders {

			if t, ok := obj.Data.(types.TakeDamage); ok {
				t.GetCamp()
			}
			if t, ok := obj.Data.(*scenes.Scenes); ok {
				if t.Type == scenes.Steel {
					b.Collider.Destruction()
					delete(globalBullets, b.Index)
				} else if t.Type == scenes.Brick {
					b.Collider.Destruction()
					t.Collider.Destruction()
					delete(globalBullets, b.Index)
					scenes.Delete(t)
				} else {
					b.Collider.Position = b.Collider.Position.Add(b.Speed)
				}
			}
		}
	} else {
		b.Collider.Move(b.Speed)
	}

	// 更新自身在网格内的位置
	b.Collider.Update()

}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(b.Collider.Position.X, b.Collider.Position.Y)
	screen.DrawImage(b.Image, opt)
}

var step float64 = 5

// 全局子弹列表
var globalBullets = make(map[int]*Bullet)

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
	Position vector2.Vector
	// Speed     *vector2.Vector2    // 子弹速度如果都相同，可以通过方向计算出来
	Direction direction.Direction
}

var index = 0

func Create(opt *CreateOption) {
	index += 1
	bullet := &Bullet{
		Collider:  collision.NewCollider(opt.Position.X, opt.Position.Y, 3, 3),
		Direction: opt.Direction,
		Speed:     opt.Direction.DirectionVector2().MulScale(step),
		Image:     bullet.BulletImage,
		Index:     index,
	}
	//  TODO: 设置碰撞器

	bullet.Collider.Data = bullet
	globalBullets[bullet.Index] = bullet
}
