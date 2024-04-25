package collision

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Game struct {
	Space       *resolv.Space
	Geometry    []*resolv.Object // 四周围墙
	Bouncers    []*Bouncer       // 所有小球
	MaxBouncers int
}

type Bouncer struct {
	Index  int
	Object *resolv.Object
	Speed  resolv.Vector
}

func NewGame() *Game {
	w := &Game{
		MaxBouncers: 3000,
	}

	w.Init()

	return w
}

func (world *Game) Init() {

	gw ,gh := config.GetWindowSize()

	cellSize := 8

	// gw, gh 为碰撞网格的长和高，cellSize 为碰撞检测的精度（cellSize*cellSize 的方格内就算碰撞）
	world.Space = resolv.NewSpace(gw, gh, cellSize, cellSize)

	world.Geometry = []*resolv.Object{
		resolv.NewObject(0, 0, 16, float64(gh)),
		resolv.NewObject(float64(gw-16), 0, 16, float64(gh)),
		resolv.NewObject(0, 0, float64(gw), 16),
		resolv.NewObject(0, float64(gh-24), float64(gw), 32),
	}

	world.Space.Add(world.Geometry...)

	world.Bouncers = []*Bouncer{}

	for i := 0; i < 10; i++ {
		world.SpawnObject(i)
	}
}

func (world *Game) SpawnObject(index int) {

	bouncer := &Bouncer{
		Index:  index,
		Object: resolv.NewObject(0, 0, 2, 2),
		Speed: resolv.NewVector(
			(rand.Float64()*8)-4,
			(rand.Float64()*8)-4,
		),
	}
	bouncer.Object.Data = bouncer

	world.Space.Add(bouncer.Object)

	// 随机一个未被占有的位置，生成小球
	var c *resolv.Cell
	for c == nil {
		rx := rand.Intn(world.Space.Width())
		ry := rand.Intn(world.Space.Height())
		c = world.Space.Cell(rx, ry)
		if c.Occupied() {
			c = nil
		} else {
			bouncer.Object.Position.X, bouncer.Object.Position.Y = world.Space.SpaceToWorld(c.X, c.Y)
		}
	}

	world.Bouncers = append(world.Bouncers, bouncer)

}

func (world *Game) Update() error {

	// 遍历所有小球处理碰撞关系
	for _, b := range world.Bouncers {

		b.Speed.Y += 0.1

		dx := b.Speed.X
		dy := b.Speed.Y

		// 检测 x 轴是否碰撞，如果碰撞将 x 轴速度反向，下面的 y 轴处理同理
		if check := b.Object.Check(dx, 0); check != nil {
			contact := check.ContactWithCell(check.Cells[0])
			dx = contact.X
			b.Speed.X *= -1

			// 打印发生碰撞的小球编号
			for _, obj := range check.Objects {
				if t, ok := obj.Data.(*Bouncer); ok {
					fmt.Println(b.Index, t.Index)
				}
			}
		}

		b.Object.Position.X += dx

		if check := b.Object.Check(0, dy); check != nil {
			contact := check.ContactWithCell(check.Cells[0])
			dy = contact.Y
			b.Speed.Y *= -1
			for _, obj := range check.Objects {
				if t, ok := obj.Data.(*Bouncer); ok {
					fmt.Println(b.Index, t.Index)
				}
			}
		}

		b.Object.Position.Y += dy

		// 更新自身在网格内的位置
		b.Object.Update()

	}

	return nil
}

func (world *Game) Draw(screen *ebiten.Image) {
	for _, o := range world.Geometry {
		ebitenutil.DrawRect(screen, o.Position.X, o.Position.Y, o.Size.X, o.Size.Y, color.RGBA{60, 60, 60, 255})
	}

	for _, b := range world.Bouncers {
		o := b.Object
		ebitenutil.DrawRect(screen, o.Position.X, o.Position.Y, o.Size.X, o.Size.Y, color.RGBA{0, 80, 255, 255})
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return config.GetWindowSize()
}
