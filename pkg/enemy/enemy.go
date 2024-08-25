package enemy

import (
	"math/rand"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	Enemynumers = 0
	Limit       = 5
	index       = 0
)

func init() {
	tankbattle.RegisterUpdate(Update, 10)
}

func Update() {

	if GetCreatEnemy() && Enemynumers < Limit {
		dx := 100
		dy := 100
		t := tank.TankBorn(dx, dy)
		if t.X != dx || t.Y != dy {
			enemy := &Enemy{
				Attack:         false,
				MoveDuration:   0,
				AttackDuration: 0,
				Index:          index,
				Direction:      direction.Up,
				Tank:           tank.New("NPC", t.X, t.Y),
			}
			Enemynumers++
			index++
			globalEnemy[enemy.Index] = enemy
		}
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
	cfg := config.C.Network
	if cfg.Type == "client" {
		return false
	}
	return inpututil.IsKeyJustPressed(ebiten.KeyQ)
}
