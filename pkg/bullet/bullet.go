package bullet

import (
	//"fmt"

	//"fmt"

	//"fmt"

	"github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
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

	flag := true
	if check := b.Collider.Check(dx, dy); check != nil {

		for _, obj := range check.Colliders {

			if t , ok := obj.Data.(types.TakeDamage);ok{
				if t.GetCamp() != b.Camp {
					t.TakeDamage(b.Damage)
					if tt ,ok := obj.Data.(types.Obstacle);ok {
						if !tt.BulletIsPassable(){
							flag = false
						}
					}
				}else {
					continue
				}
			}
			
		}
	}
	if flag {
		b.Collider.Move(b.Speed)
	}else{
		b.Collider.Destruction()
		delete(globalBullets, b.Index)
	}

	// 更新自身在网格内的位置
	b.Collider.Update()
	
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(b.Collider.Position.X, b.Collider.Position.Y)
	screen.DrawImage(b.Image, opt)
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
	Direction direction.Direction
	Camp      string
}

var index = 0

func Create(opt *CreateOption) {
	index += 1
	Bullx, Bully := bullet.BulletImage.Bounds().Dx(), bullet.BulletImage.Bounds().Dy()
	bullet := &Bullet{
		Collider:  collision.NewCollider(opt.Position.X+float64(tank.EnemyImage.Bounds().Dx())/2, opt.Position.Y+float64(tank.EnemyImage.Bounds().Dy())/2, float64(Bullx), float64(Bully)),
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
