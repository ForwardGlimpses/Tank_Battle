package weapon

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/vector2"
)

var (
	Weapons = make(map[int]Weapon)
)

type Weapon interface {
	Fight(position vector2.Vector, direction direction.Direction, camp string, playerIndex int)
	Cooling()
}

// 当前武器是一个抽象概念，不需要实际的图片
type DefaultWeapon struct {
	Damage       int
	CoolingCount int
}

func (d *DefaultWeapon) Fight(position vector2.Vector, direction direction.Direction, camp string, playerIndex int) {
	// TODO: 客户端冷却数值不准确
	if d.CoolingCount > 0 {
		return
	}

	opt := &bullet.CreateOption{
		Position:    position,
		Direction:   direction,
		Camp:        camp,
		PlayerIndex: playerIndex,
	}
	bullet.Create(opt)
	d.CoolingCount = 60
}

func (d *DefaultWeapon) Cooling() {
	d.CoolingCount--
}

func GetWeapon(Type int) Weapon {
	return Weapons[Type]
}

func init() {
	Weapons[0] = &DefaultWeapon{
		Damage:       50,
		CoolingCount: 0,
	}
}
