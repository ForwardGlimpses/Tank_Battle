package weapon

import "image"

type Weapon interface {
	Fight(direction int, dx int, dy int)
}

type DefaultWeapon struct {
	Damage  int
	Picture image.Image
}

func (D *DefaultWeapon)Fight(direction int, dx int, dy int) {
	

}
