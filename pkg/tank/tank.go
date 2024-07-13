package tank

import (
	//"fmt"
	"image"
	_ "image/png"
	"math"

	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	//"github.com/ForwardGlimpses/Tank_Battle/pkg/scenes"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
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

var GlobalTanks = make(map[int]*Tank)

var TankIndex = 0

type Tank struct {
	Hp       int
	Collider *collision.Collider
	Direction direction.Direction
	weapon    weapon.Weapon
	Image     image.Image
	Attack    bool
	Move      bool
	Camp      string
	Index     int
}

func New(camp string,tankx int,tanky int) *Tank {
	tank := &Tank{
		Collider: collision.NewCollider(float64(tankx), float64(tanky), float64(tank.PlayerImage.Bounds().Dx()), float64(tank.PlayerImage.Bounds().Dy())),
		Hp:       100,
		weapon:   &weapon.DefaultWeapon{},
		Image:    tank.TankImage[camp],
		Camp:     camp,
		Index:    TankIndex,
	}
	tank.Collider.Data = tank
	GlobalTanks[tank.Index] = tank
	TankIndex ++
	return tank
}

func init() {
	tankbattle.RegisterDraw(Draw,1)
	tankbattle.RegisterUpdate(Update,3)
}

func (t *Tank) Update(direction direction.Direction) {
	t.Direction = direction
	increment := direction.DirectionVector2().MulScale(step)
	dx := increment.X
	dy := increment.Y
	stop := false
	if check := t.Collider.Check(dx, dy); check != nil {
		// TODO: 这里需要判断是否碰到障碍物，如果没碰到，正常移动
		for _, obj := range check.Colliders {
			if _, ok := obj.Data.(*Tank); ok {
				stop = true
			}
			if  tt , ok := obj.Data.(types.Obstacle); ok {
				if !tt.TankIsPassable() {
					stop = true
				}
			}
		}
	}
	if !stop {
		t.Collider.Position = t.Collider.Position.Add(direction.DirectionVector2().MulScale(step))
	}
	 // 更新自身在网格内的位置
	t.Collider.Update()
}

func Update() {
	for _,tank := range GlobalTanks{
		if tank.Move {
			tank.Update(tank.Direction)
		}
	}
	for _,tank := range GlobalTanks{
		if tank.Attack {
			tank.Fight()
			tank.Attack = false
		}
	}
}

func (t *Tank) Fight() {
	// TODO: 计算子弹发射位置（坦克正前方）
	t.weapon.Fight(t.Collider.Position, t.Direction, t.Camp)
}


func (t *Tank) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	tranX := float64(t.Image.Bounds().Dx()) / 2
	tranY := float64(t.Image.Bounds().Dy()) / 2
	opt.GeoM.Translate(-tranX, -tranY)
	opt.GeoM.Rotate(t.Direction.Theta() * 2 * math.Pi / 360)
	opt.GeoM.Translate(t.Collider.Position.X+tranX, t.Collider.Position.Y+tranY)
	//screen.DrawImage(t.Image, opt)
	screen.DrawImage(ebiten.NewImageFromImage(t.Image), opt)

}

func Draw(screen *ebiten.Image) {
	for _,tank := range GlobalTanks{
		tank.Draw(screen)
	}
}

func (t *Tank) Obstacle() {

}

func (t *Tank) TankIsPassable() bool {
		return false
}

func (t *Tank) BulletIsPassable() bool {
		return false
}

func (t *Tank) GetCamp() string {
	return t.Camp
}

func (t *Tank) TakeDamage(damage int) {
	t.Hp -= damage
}