package enemy

import (
	//"fmt"
	//"fmt"
	"math/rand"

	//"time"
	//"sync"
	tankImage "github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	Up int = iota
	Down
	Left
	Right
)

type Enemy struct {
	MoveDuration   int
	AttackDuration int
	Direction      direction.Direction
	Attack         bool
	Tank           *tank.Tank
	Index          int
}

var (
	globalEnemy = make(map[int]*Enemy)
	index       = 0
	Enemynumers = 0
	Limit       = 5
)

func init() {
	tankbattle.RegisterUpdate(Update, 1)
}

func Update() {

	if GetCreatEnemy() && Enemynumers < Limit {
		Enemynumers++
		hx, hy := config.GetWindowSize()
		var dx, dy int
		for {
			dx = rand.Intn(hx)
			dy = rand.Intn(hy)
			t := collision.NewCollider(float64(dx), float64(dy), float64(tankImage.EnemyImage.Bounds().Dx()), float64(tankImage.EnemyImage.Bounds().Dy()))
			if check := t.Check(0, 0); check == nil {
				break
			}
		}
		index += 1
		enemy := &Enemy{
			Attack:         false,
			MoveDuration:   0,
			AttackDuration: 0,
			Index:          index,
			Direction:      direction.Direction(Up),
			Tank:           tank.New("NPC", dx, dy),
		}

		globalEnemy[enemy.Index] = enemy
	}

	var Destroyed []Enemy

	for _, enemy := range globalEnemy {
		if enemy.Tank.Hp <= 0 {
			Destroyed = append(Destroyed, *enemy)
		} else {
			enemy.Update()
		}
	}

	for _, enemy := range Destroyed {
		delete(globalEnemy, enemy.Index)
		Enemynumers--
		delete(tank.GlobalTanks, enemy.Tank.Index)
		enemy.Tank.Collider.Destruction()
	}

}

func (a *Enemy) Update() {

	if a.MoveDuration > 0 {
		a.MoveDuration--
	} else {
		a.randMove()
		a.IsMove()
	}
	if a.AttackDuration > 0 {
		a.AttackDuration--
	} else {
		a.randAttack()
	}

	// 攻击
	if a.Attack {
		a.Tank.Attack = true
		a.Attack = false
	}

	a.Tank.Direction = a.Direction
	tank.GlobalTanks[a.Tank.Index] = a.Tank
}

var baseDuration = 60

func (a *Enemy) randAttack() {
	randDuration := rand.Intn(60)
	a.AttackDuration = baseDuration + randDuration

	x := rand.Intn(2)
	if x == 1 {
		a.Attack = true
	} else {
		a.Attack = false
	}
}

func (a *Enemy) randMove() {
	randDuration := rand.Intn(60)
	a.MoveDuration = baseDuration + randDuration
	dir := rand.Intn(4)
	a.Direction = direction.Direction(dir)
}

func (a *Enemy) IsMove() {
	stop := rand.Intn(2)
	if stop%2 == 1 {
		a.Tank.Move = false
	} else {
		a.Tank.Move = true
	}
}

func GetCreatEnemy() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyQ)
}
