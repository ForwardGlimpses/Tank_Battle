package bullet

import (
	//"fmt"

	"github.com/ForwardGlimpses/Tank_Battle/assets/bullet"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/network"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/json"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/vector2"
	//"github.com/ForwardGlimpses/Tank_Battle/pkg/weapon"
)

func init() {
	network.RegisterClient("bullet", &neteworkClient{})
	network.RegisterServer("bullet", &networkServer{})
}

// const (
// 	Weapon int = 0
// )

type bulletMassage struct {
	Dx        int
	Dy        int
	Direction direction.Direction
	Speed     vector2.Vector
	Index     int
	Damage    int
	Camp      string
}

type neteworkClient struct{}

func (a *neteworkClient) Send() string {
	return ""
}

func (a *neteworkClient) Receive(m string) {
	massage := []bulletMassage{}
	json.Unmarshal([]byte(m), &massage)
	for _, bulletmassage := range massage {
		_, ok := globalBullets[bulletmassage.Index]
		if ok {
			//globalBullets[bulletmassage.Index].Direction = bulletmassage.Direction
			globalBullets[bulletmassage.Index].Collider.Position.X = float64(bulletmassage.Dx)
			globalBullets[bulletmassage.Index].Collider.Position.Y = float64(bulletmassage.Dy)
			//globalBullets[bulletmassage.Index].Speed = bulletmassage.Speed
		} else {
			bullet := &Bullet{
				Collider:  collision.NewCollider(float64(bulletmassage.Dx), float64(bulletmassage.Dy), float64(bullet.BulletImage[bulletmassage.Camp].Bounds().Dx()), float64(bullet.BulletImage[bulletmassage.Camp].Bounds().Dy())),
				Direction: bulletmassage.Direction,
				Image:     bullet.BulletImage[bulletmassage.Camp],
				Camp:      bulletmassage.Camp,
				Index:     bulletmassage.Index,
				Damage:    bulletmassage.Damage,
				Speed:     bulletmassage.Speed,
			}
			globalBullets[bullet.Index] = bullet
		}
	}
}

type networkServer struct{}

func (a *networkServer) Send() string {
	massage := []bulletMassage{}
	for _, bullet := range globalBullets {
		massage = append(massage, bulletMassage{
			Index:     bullet.Index,
			Dx:        int(bullet.Collider.Position.X),
			Dy:        int(bullet.Collider.Position.Y),
			Direction: bullet.Direction,
			Camp:      bullet.Camp,
			Damage:    bullet.Damage,
			Speed:     bullet.Speed,
		})
	}
	date := json.MarshalToString(massage)
	return date

}

func (a *networkServer) Receive(m string) {
}
