package bullet

import (
	"image"

	"github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct{
	dx int
	dy int
	direction int 
	speed int
	Image image.Image
}


func New() *Bullet {
	return &Bullet{
		dx:    0,
		dy:    0,
		direction: 0,
		speed:  10,
		Image: bullet.BulletImage,
	}
}
func (b *Bullet) Move(Tank_dx int,Tank_dy int){
	b.dx = Tank_dx
	b.dy = Tank_dy
	

}
func (b *Bullet) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(b.dx), float64(b.dy))
	screen.DrawImage(ebiten.NewImageFromImage(b.Image), opt)
}
