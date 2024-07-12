package bullet

import (
	//"fmt"

	//"fmt"

	"github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/scenes"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	//"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
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
	Damage    int
	Camp      string
}

func init() {
	tankbattle.RegisterUpdate(Update,3)
	tankbattle.RegisterDraw(Draw,3)
}

func (b *Bullet) Update() {

	dx := b.Speed.X
	dy := b.Speed.Y

	// 检测 x 或 y 轴是否碰撞，如果碰撞将子弹销毁
	if check := b.Collider.Check(dx, dy); check != nil {

		// 打印发生碰撞的小球编号
		for _, obj := range check.Colliders {

			if t, ok := obj.Data.(types.TakeDamage); ok {
				if t.GetCamp() != b.Camp {
					t.TakeDamage(b.Damage)
					if tt, ok := obj.Data.(*scenes.Scenes); ok {
						if tt.Type == scenes.Rivers || tt.Type == scenes.Grass {
					 		b.Collider.Move(b.Speed)
					        break ;
						} else if tt.Type == scenes.Brick {
							b.Collider.Destruction()
							delete(globalBullets, b.Index)
							if tt.Hp <= 0 {
								scenes.Delete(tt)
								tt.Collider.Destruction()
							}
						} else if tt.Type == scenes.Steel{
							b.Collider.Destruction()
							delete(globalBullets, b.Index)
						}
					} else {
						b.Collider.Destruction()
						delete(globalBullets, b.Index)
					}
				} else {
					b.Collider.Move(b.Speed)
					break ;
				}
			} else {
				b.Collider.Move(b.Speed)
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

func (b *Bullet) Obstacle() {

}

var step float64 = 4

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
	Camp      string
}

var index = 0

func Create(opt *CreateOption) {
	index += 1
	dirx := opt.Direction.DirectionVector2().X
	diry := opt.Direction.DirectionVector2().Y
	dxx := 0
	dyy := 0
	if dirx < 0 {
		dxx = -5
	}
	if diry < 0 {
		dyy = -5
	}
	Tankx, Tanky := tank.PlayerImage.Bounds().Dx()/2*int(dirx)+tank.PlayerImage.Bounds().Dx()/2+dxx, tank.PlayerImage.Bounds().Dy()/2*int(diry)+tank.PlayerImage.Bounds().Dy()/2+dyy
	Bullx, Bully := bullet.BulletImage.Bounds().Dx(), bullet.BulletImage.Bounds().Dy()
	bullet := &Bullet{
		Collider:  collision.NewCollider(opt.Position.X+float64(Tankx), opt.Position.Y+float64(Tanky), float64(Bullx), float64(Bully)),
		//Collider:  collision.NewCollider(opt.Position.X+float64(tank.EnemyImage.Bounds().Dx()), opt.Position.Y+float64(tank.EnemyImage.Bounds().Dy()), float64(Bullx), float64(Bully)),
		Direction: opt.Direction,
		Speed:     opt.Direction.DirectionVector2().MulScale(step),
		Image:     bullet.BulletImage,
		Index:     index,
		Damage:    50,
		Camp:      opt.Camp,
	}
	//  TODO: 设置碰撞器

	bullet.Collider.Data = bullet
	globalBullets[bullet.Index] = bullet
}
