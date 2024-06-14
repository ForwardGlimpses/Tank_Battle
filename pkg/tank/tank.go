package tank

import (
	"fmt"
	"image"
	_ "image/png"
	"math"

	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/types"

	//"github.com/ForwardGlimpses/Tank_Battle/pkg/scenes"
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

// var globalBullets = make(map[int]*Tank)

type Tank struct {
	Hp       int
	Collider *collision.Collider
	//Position  vector2.Vector
	direction direction.Direction
	weapon    weapon.Weapon
	Image     image.Image
	Camp      string
}

func New(camp string,tankx int,tanky int) *Tank {
	return &Tank{
		Collider: collision.NewCollider(float64(tankx), float64(tanky), float64(tank.PlayerImage.Bounds().Dx()), float64(tank.PlayerImage.Bounds().Dy())),
		Hp:       100,
		weapon:   &weapon.DefaultWeapon{},
		Image:    tank.TankImage[camp],
		Camp:     camp,
	}
}

func (t *Tank) Move(direction direction.Direction) {
	t.direction = direction
	increment := direction.DirectionVector2().MulScale(step)
	dx := increment.X
	dy := increment.Y
	stop := false
	if check := t.Collider.Check(dx, dy); check != nil {
		// TODO: 这里需要判断是否碰到障碍物，如果没碰到，正常移动
		for _, obj := range check.Colliders {
			if _, ok := obj.Data.(*Tank); ok {
				fmt.Print(t.Hp)
			}
			if  _ , ok := obj.Data.(types.Obstacle); ok {
				stop = true
			}
		}
	}
	if !stop {
		t.Collider.Position = t.Collider.Position.Add(direction.DirectionVector2().MulScale(step))
	}
	 // 更新自身在网格内的位置
	t.Collider.Update()
}

func (t *Tank) Fight() {
	// TODO: 计算子弹发射位置（坦克正前方）
	t.weapon.Fight(t.Collider.Position, t.direction, t.Camp)
}

func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	tranX := float64(t.Image.Bounds().Dx()) / 2
	tranY := float64(t.Image.Bounds().Dy()) / 2
	opt.GeoM.Translate(-tranX, -tranY)
	opt.GeoM.Rotate(t.direction.Theta() * 2 * math.Pi / 360)
	opt.GeoM.Translate(t.Collider.Position.X+tranX, t.Collider.Position.Y+tranY)
	//screen.DrawImage(t.Image, opt)
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)

}

func (t *Tank) Obstacle() {

}