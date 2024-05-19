package tank

import (
	//"fmt"
	"fmt"
	"image"
	_ "image/png"
	"math"

	//"github.com/ForwardGlimpses/Tank_Battle/pkg/scenes"
	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/direction"

	//	"github.com/ForwardGlimpses/Tank_Battle/pkg/vector2"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/weapon"
	"github.com/hajimehoshi/ebiten/v2"

	//	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
)

const (
	Up int = iota
	Down
	Left
	Right
	step float64 = 3
)

type Tank struct {
	Hp        int
	Collider  *resolv.Object
	//Position  resolv.Vector
	direction direction.Direction
	weapon    weapon.Weapon
	Image     image.Image
}

func New() *Tank {
	return &Tank{
		//Position: resolv.NewVector(28, 25),
		Collider: resolv.NewObject(60,60,20,20),
		Hp:       100,
		weapon:   &weapon.DefaultWeapon{},
		Image:    tank.PlayerImage,
	}
}

func (t *Tank) Move(direction direction.Direction) {
	t.direction = direction
	increment:= direction.DirectionVector2().Scale(step)
	dx := increment.X
	dy := increment.Y
	if check := t.Collider.Check(dx, dy); check != nil {
		// 打印发生碰撞的小球编号
		for _, obj := range check.Objects {
			
			if _, ok := obj.Data.(*Tank); ok {
				fmt.Print(t.Hp)
				//scenes.SpaceRemove(t.Collider)
			}
		}
	} else {
		t.Collider.Position = t.Collider.Position.Add(direction.DirectionVector2().Scale(step))
	}
	// // 更新自身在网格内的位置
	t.Collider.Update()

	// TODO: 通过碰撞检测限制移动
	
}

func (t *Tank) Fight() {
	// TODO: 计算子弹发射位置（坦克正前方）
	t.weapon.Fight(t.Collider.Position, t.direction)
}

func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(55)/2, -float64(49)/2)
	opt.GeoM.Rotate(t.direction.Theta() * 2 * math.Pi / 360)
	opt.GeoM.Translate(t.Collider.Position.X,t.Collider.Position.Y)
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)
}
