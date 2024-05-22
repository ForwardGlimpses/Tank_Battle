package tank

import (
	//"fmt"
	"fmt"
	"image"
	_ "image/png"
	"math"

	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/weapon"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Up int = iota
	Down
	Left
	Right
	step float64 = 3
)

type Tank struct {
	Hp       int
	Collider *collision.Collider
	//Position  vector2.Vector
	direction direction.Direction
	weapon    weapon.Weapon
	Image     image.Image
}

func New() *Tank {
	return &Tank{
		Collider: collision.NewCollider(60, 60, 20, 20),
		Hp:       100,
		weapon:   &weapon.DefaultWeapon{},
		Image:    tank.PlayerImage,
	}
}

func (t *Tank) Move(direction direction.Direction) {
	t.direction = direction
	increment := direction.DirectionVector2().MulScale(step)
	dx := increment.X
	dy := increment.Y
	if check := t.Collider.Check(dx, dy); check != nil {
		// TODO: 这里需要判断是否碰到障碍物，如果没碰到，正常移动
		for _, obj := range check.Colliders {
			if _, ok := obj.Data.(*Tank); ok {
				fmt.Print(t.Hp)
			}
		}
	} else {
		t.Collider.Position = t.Collider.Position.Add(direction.DirectionVector2().MulScale(step))
	}
	// // 更新自身在网格内的位置
	t.Collider.Update()
}

func (t *Tank) Fight() {
	// TODO: 计算子弹发射位置（坦克正前方）
	t.weapon.Fight(t.Collider.Position, t.direction)
}

func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(55)/2, -float64(49)/2)
	opt.GeoM.Rotate(t.direction.Theta() * 2 * math.Pi / 360)
	opt.GeoM.Translate(t.Collider.Position.X, t.Collider.Position.Y)
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)
}
