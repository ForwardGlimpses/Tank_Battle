package bullet

import (
	//"fmt"
	//"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/direction"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

// type Game struct {
// 	Space       *resolv.Space
// 	Geometry    []*resolv.Object // 四周围墙
// 	Bullets     []*Bullet     // 所有小球
// 	MaxBouncers int
// }


type Bullet struct {
	Object    resolv.Object
	Position  resolv.Vector
	Direction direction.Direction
	Speed     resolv.Vector
	Image     *ebiten.Image
	Index     int
}

// func NewGame() *Game {
// 	w := &Game{
// 		MaxBouncers: 3000,
// 	}

// 	w.Init()

// 	return w
// }

// func (world *Game) Init() {

// 	gw ,gh := config.GetWindowSize()

// 	cellSize := 8

// 	// gw, gh 为碰撞网格的长和高，cellSize 为碰撞检测的精度（cellSize*cellSize 的方格内就算碰撞）
// 	world.Space = resolv.NewSpace(gw, gh, cellSize, cellSize)

// 	world.Geometry = []*resolv.Object{
// 		resolv.NewObject(0, 0, 16, float64(gh)),
// 		resolv.NewObject(float64(gw-16), 0, 16, float64(gh)),
// 		resolv.NewObject(0, 0, float64(gw), 16),
// 		resolv.NewObject(0, float64(gh-24), float64(gw), 32),
// 	}

// 	world.Space.Add(world.Geometry...)

// 	world.Bullets = []*Bullet{}

// 	// for i := 0; i < 10; i++ {
// 	// 	world.SpawnObject(i)
// 	// }
// }



func (b *Bullet) Update() {
	b.Position = b.Position.Add(b.Speed)

	// b.Object.Position.X += b.Speed.X
	// b.Object.Position.Y += b.Speed.Y
	// // b.Speed.Y += 0.1

	// 	dx := b.Speed.X
	// 	dy := b.Speed.Y

	// 	// 检测 x 轴是否碰撞，如果碰撞将 x 轴速度反向，下面的 y 轴处理同理
	// 	if check := b.Object.Check(dx, 0); check != nil {
	// 		// contact := check.ContactWithCell(check.Cells[0])
	// 		// dx = contact.X
	// 		//b.Speed.X *= -1

	// 		// 打印发生碰撞的小球编号
	// 		for _, obj := range check.Objects {
	// 			if t, ok := obj.Data.(*Bullet); ok {
	// 				fmt.Println(b.Index, t.Index)
	// 			}
	// 		}
	// 	}

	// 	//b.Object.Position.X += dx

	// 	if check := b.Object.Check(0, dy); check != nil {
	// 		// contact := check.ContactWithCell(check.Cells[0])
	// 		// dy = contact.Y
	// 		//b.Speed.Y *= -1
	// 		for _, obj := range check.Objects {
	// 			if t, ok := obj.Data.(*Bullet); ok {
	// 				fmt.Println(b.Index, t.Index)
	// 			}
	// 		}
	// 	}

	// 	//b.Object.Position.Y += dy

	// 	// 更新自身在网格内的位置
	// 	b.Object.Update()

}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(b.Position.X,b.Position.Y)
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
	//Object *resolv.Object
	Position  resolv.Vector
	// Speed     *vector2.Vector2    // 子弹速度如果都相同，可以通过方向计算出来
	Direction direction.Direction
}

var index = 0

func Create(opt *CreateOption) {
	index += 1
	bullet := &Bullet {
		Position:  opt.Position,
		Direction: opt.Direction,
		Speed:     opt.Direction.DirectionVector2().Scale(step),
		Image:     bullet.BulletImage,
		Index:     index,
	}
	//  TODO: 设置碰撞器
	
	globalBullets[bullet.Index] = bullet
}
