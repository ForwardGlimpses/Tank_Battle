package bullet

import (
	"fmt"
	//"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Bullet struct {
	Collider  *resolv.Object
	Direction direction.Direction
	Speed     resolv.Vector
	Image     *ebiten.Image
	Index     int
}

func (b *Bullet) Update() {

		dx := b.Speed.X
		dy := b.Speed.Y

		// 检测 x 轴是否碰撞，如果碰撞将 x 轴速度反向，下面的 y 轴处理同理
		if check := b.Collider.Check(dx, dy); check != nil {

			// 打印发生碰撞的小球编号
			for _, obj := range check.Objects {
				
				if t, ok := obj.Data.(*Bullet); ok {
					fmt.Println(b.Index, t.Index)
					scenes.SpaceRemove(b.Collider)
					scenes.SpaceRemove(t.Collider)
					delete(globalBullets,b.Index)
					delete(globalBullets,t.Index)
				}
				if t, ok := obj.Data.(*scenes.Scenes); ok {
					if t.Type == scenes.Steel {
						scenes.SpaceRemove(b.Collider)
						delete(globalBullets,b.Index)
					} else if t.Type == scenes.Brick {
						scenes.SpaceRemove(b.Collider)
						scenes.SpaceRemove(t.Collider)
						delete(globalBullets,b.Index)
						scenes.Delete(t)
					} else {
						b.Collider.Position=b.Collider.Position.Add(b.Speed)
					}
				}
			}
		} else {
			b.Collider.Position=b.Collider.Position.Add(b.Speed)
		}
		
		// 更新自身在网格内的位置
		b.Collider.Update()
		
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(b.Collider.Position.X,b.Collider.Position.Y)
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
	Position  resolv.Vector
	// Speed     *vector2.Vector2    // 子弹速度如果都相同，可以通过方向计算出来
	Direction direction.Direction
}

var index = 0

func Create(opt *CreateOption) {
	index += 1
	bullet := &Bullet {
		Collider:  resolv.NewObject(opt.Position.X,opt.Position.Y,3,3),
		Direction: opt.Direction,
		Speed:     opt.Direction.DirectionVector2().Scale(step),
		Image:     bullet.BulletImage,
		Index:     index,
	}
	//  TODO: 设置碰撞器

	bullet.Collider.Data = bullet
	scenes.SpaceAdd(bullet.Collider)
	globalBullets[bullet.Index] = bullet
}
