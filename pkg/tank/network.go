package tank

import (
	"github.com/ForwardGlimpses/Tank_Battle/assets/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/network"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/collision"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/direction"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/json"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/weapon"
)

func init() {
	network.RegisterClient("tank", &neteworkClient{})
	network.RegisterServer("tank", &networkServer{})
}

const (
	Weapon int = 0
)

type tankMassage struct {
	Index     int
	Hp        int
	dx        int
	dy        int
	Direction direction.Direction
	weapon    int
	Attack    bool
	Move      bool
	Camp      string
}

type neteworkClient struct{}

func (a *neteworkClient) Send() string {
	return ""
}

func (a *neteworkClient) Receive(m string) {
	massage := []tankMassage{}
	json.Unmarshal([]byte(m), &massage)
	for _, tankmassage := range massage {
		tank := &Tank{
			Hp:        tankmassage.Hp,
			Collider:  collision.NewCollider(float64(tankmassage.dx),float64(tankmassage.dy),float64(tank.PlayerImage.Bounds().Dx()),float64(tank.PlayerImage.Bounds().Dy())),
			Direction: tankmassage.Direction,
			weapon:    weapon.GetWeapon(Weapon),
			Image:     tank.TankImage[tankmassage.Camp],
			Attack:    tankmassage.Attack,
			Move:      tankmassage.Move,
			Camp:      tankmassage.Camp,
			Index:     TankIndex,
		}
		GlobalTanks[tankmassage.Index] = tank
	}
}

type networkServer struct{}

func (a *networkServer) Send() string {
	massage := []tankMassage{}
	for _, tank := range GlobalTanks {
		massage = append(massage, tankMassage{
			Index:     TankIndex,
			Hp:        tank.Hp,
			dx:        int(tank.Collider.Position.X),
			dy:        int(tank.Collider.Position.Y),
			Direction: tank.Direction,
			weapon:    Weapon,
			Attack:    tank.Attack,
			Move:      tank.Move,
			Camp:      tank.Camp,
		})
	}
	return json.MarshalToString(massage)
}

func (a *networkServer) Receive(m string) {

}
