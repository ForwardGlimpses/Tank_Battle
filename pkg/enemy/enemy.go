package enemy

import (
	"math/rand"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
)

type Enemy struct {
	Camp           string
	MoveDuration   int
	AttackDuration int
	Direction      direction.Direction
	Attack         bool
	Tank           *tank.Tank
}

func (a *Enemy) Update() {
	if a.MoveDuration > 0 {
		a.MoveDuration--
	} else {
		a.randMove()
	}
	if a.AttackDuration > 0 {
		a.AttackDuration--
	} else {
		a.randAttack()
	}

	// 攻击
	if a.Attack {

	}
	a.Tank.Move(a.Direction)
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
