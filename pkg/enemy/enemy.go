package enemy

import (
	//"fmt"
	//"fmt"
	"math/rand"

	//"time"
	//"sync"
	//"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/player"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/inpututil"
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

var globalEnemy = make(map[int]*Enemy)

var index = 0

func Update() {
	if player.GetCreatEnemy() {

		hx, hy := config.GetWindowSize()
		var dx, dy int
		var t *collision.Collider
		for {
			dx = rand.Intn(hx)
			dy = rand.Intn(hy)
			if check := t.Check(float64(dx), float64(dy)); check != nil {

			} else {
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

	for _, enemy := range globalEnemy {
		enemy.Update()
	}
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
		a.Tank.Fight()
		a.Attack = false
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

func (p *Enemy) Draw(screen *ebiten.Image) {
	p.Tank.Draw(screen)
}

func Draw(screen *ebiten.Image) {
	for _, enemy := range globalEnemy {
		enemy.Draw(screen)
	}
}

func Obstacle() {

}
