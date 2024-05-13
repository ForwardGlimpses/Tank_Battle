package bullet

import (
	"fmt"
	//"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/direction"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/scenes"
)


type Bullet struct {
	Collider    *resolv.Object
	Direction direction.Direction
	Speed     resolv.Vector
	Image     *ebiten.Image
	Index     int
}


func Init() {

	gw ,gh := 100, 100

	cellSize := 8

	// gw, gh 为碰撞网格的长和高，cellSize 为碰撞检测的精度（cellSize*cellSize 的方格内就算碰撞）
	scenes.World.Space = resolv.NewSpace(gw, gh, cellSize, cellSize)

	world.Geometry = []*resolv.Object{
		resolv.NewObject(0, 0, 16, float64(gh)),
		resolv.NewObject(float64(gw-16), 0, 16, float64(gh)),
		resolv.NewObject(0, 0, float64(gw), 16),
		resolv.NewObject(0, float64(gh-24), float64(gw), 32),
	}

	world.Space.Add(world.Geometry...)

	world.Bullets = []*Bullet{}

}


func (b *Bullet) Update() {

		dx := b.Speed.X + b.Collider.Position.X
		dy := b.Speed.Y + b.Collider.Position.Y

		//检测 x 轴是否碰撞，如果碰撞将 x 轴速度反向，下面的 y 轴处理同理
		if check := b.Collider.Check(dx, dy); check != nil {
			// 打印发生碰撞的小球编号
			for _, obj := range check.Objects {
				if t, ok := obj.Data.(*Bullet); ok {
					fmt.Println(b.Index, t.Index)
				}
			}
		 } else {
			b.Collider.Position = b.Collider.Position.Add(b.Speed)
		}

		// 更新自身在网格内的位置
		b.Collider.Update()

}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(b.Collider.Position.X,b.Collider.Position.Y)
	screen.DrawImage(b.Image, opt)
}

var step float64 = 1

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
	//Object *resolv.Object
	Position  resolv.Vector
	// Speed     *vector2.Vector2    // 子弹速度如果都相同，可以通过方向计算出来
	Direction direction.Direction
}

var index = 0

func Create(opt *CreateOption) {
	index += 1
	bullet := &Bullet {
		Collider:    resolv.NewObject(opt.Position.X-10, opt.Position.Y-10, 10, 10),
		Direction: opt.Direction,
		Speed:     opt.Direction.DirectionVector2().Scale(step),
		Image:     bullet.BulletImage,
		Index:     index,
	}
	bullet.Collider.Data = bullet
	//  TODO: 设置碰撞器
	world.Space.Add(bullet.Collider)

	globalBullets[bullet.Index] = bullet
}

