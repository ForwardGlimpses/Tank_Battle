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
	Fight(position vector2.Vector, direction direction.Direction,camp string)
	IsCooling() bool
}

// 当前武器是一个抽象概念，不需要实际的图片
type DefaultWeapon struct {
	Damage int
	Cooling int
	CoolingCount int
}

func (d *DefaultWeapon) Fight(position vector2.Vector, direction direction.Direction,camp string) {
	if d.IsCooling(){
		return
	}
	opt := &bullet.CreateOption{
		Position:  position,
		Direction: direction,
		Camp: camp,
	}
	bullet.Create(opt)
	d.CoolingCount = d.Cooling
}

func (d *DefaultWeapon) IsCooling() bool {
	if d.CoolingCount > 0 {
		d.CoolingCount --
		return true
	}
	return false
}

func GetWeapon(Type int) Weapon {
    return Weapons[Type]
}

func init() {
    Weapons[0] = &DefaultWeapon{
		Damage: 50,
		Cooling: 10,
		CoolingCount: 0,
	}
}