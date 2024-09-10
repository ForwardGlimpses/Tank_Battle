package bullet

import (
	"github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/scorer"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/types"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/vector2"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	Collider    *collision.Collider
	Direction   direction.Direction
	Speed       vector2.Vector
	Image       *ebiten.Image
	Index       int
	Damage      int
	Camp        string
	PlayerIndex int
}

func init() {
	tankbattle.RegisterUpdate(Update, 30)
	tankbattle.RegisterDraw(Draw, 30)
}

func (b *Bullet) Update() {

	dx := b.Speed.X
	dy := b.Speed.Y

	flag := true
	if check := b.Collider.Check(dx, dy); check != nil {
		for _, obj := range check.Colliders {

			if t, ok := obj.Data.(types.TakeDamage); ok {
				if t.GetCamp() != b.Camp {
					t.TakeDamage(b.Damage)
					if t.GetCamp() == "NPC" {
						scorer.AddPoints(b.PlayerIndex, b.Damage)
					}
					if t.GetCamp() != "bulletIsPassable" {
						flag = false
					}
				}
			}

		}
	}
	if flag {
		b.Collider.Move(b.Speed)
	} else {
		b.Collider.Destruction()
		delete(globalBullets, b.Index)
	}
	b.Collider.Update()
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(b.Collider.Position.X, b.Collider.Position.Y)
	screen.DrawImage(b.Image, opt)
}

var step float64 = 4

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

type CreateOption struct {
	Position    vector2.Vector
	Direction   direction.Direction
	Camp        string
	PlayerIndex int
}

var index = 0

func Create(opt *CreateOption) {
	index += 1
	Bullx, Bully := bullet.BulletImage[opt.Camp].Bounds().Dx(), bullet.BulletImage[opt.Camp].Bounds().Dy()
	bullet := &Bullet{
		Collider:    collision.NewCollider(opt.Position.X+float64(tank.EnemyImage.Bounds().Dx())/2, opt.Position.Y+float64(tank.EnemyImage.Bounds().Dy())/2, float64(Bullx), float64(Bully)),
		Direction:   opt.Direction,
		Speed:       opt.Direction.DirectionVector2().MulScale(step),
		Image:       bullet.BulletImage[opt.Camp],
		Index:       index,
		Damage:      50,
		Camp:        opt.Camp,
		PlayerIndex: opt.PlayerIndex,
	}

	bullet.Collider.Data = bullet
	globalBullets[bullet.Index] = bullet
}
