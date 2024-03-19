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
	Image image.Image
}


func New() *Bullet {
	return &Bullet{
		dx:    0,
		dy:    0,
		direction: 0,
		Image: bullet.BulletImage,
	}
}
func (b *Bullet) Update(){
	
}
func (b *Bullet) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(float64(b.dx), float64(b.dy))
	screen.DrawImage(ebiten.NewImageFromImage(b.Image), opt)
}
