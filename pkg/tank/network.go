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
	Dx        int
	Dy        int
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
	//fmt.Println("接收数据:", massage)
	Survived := map[int]bool{}
	for _, tank := range GlobalTanks {
		Survived[tank.Index] = false
	}
	for _, tankmassage := range massage {
		tankTank, ok := GlobalTanks[tankmassage.Index]
		Survived[tankmassage.Index] = true
		if ok {
			tankTank.Hp = tankmassage.Hp
			tankTank.Direction = tankmassage.Direction
			tankTank.Attack = tankmassage.Attack
			tankTank.Move = tankmassage.Move
			tankTank.Collider.Position.X = float64(tankmassage.Dx)
			tankTank.Collider.Position.Y = float64(tankmassage.Dy)
		} else {
			tank := &Tank{
				Hp:        tankmassage.Hp,
				Collider:  collision.NewCollider(float64(tankmassage.Dx), float64(tankmassage.Dy), float64(tank.PlayerImage.Bounds().Dx()), float64(tank.PlayerImage.Bounds().Dy())),
				Direction: tankmassage.Direction,
				weapon:    weapon.GetWeapon(Weapon),
				Image:     tank.TankImage[tankmassage.Camp],
				Attack:    tankmassage.Attack,
				Move:      tankmassage.Move,
				Camp:      tankmassage.Camp,
				Index:     tankmassage.Index,
			}
			GlobalTanks[tank.Index] = tank
		}
	}
	for tankindex, flag := range Survived {
		if !flag {
			GlobalTanks[tankindex].Collider.Destruction()
			delete(GlobalTanks, tankindex)
		}
	}
	//fmt.Println("-------------")
}

type networkServer struct{}

func (a *networkServer) Send() string {
	massage := []tankMassage{}
	for _, tank := range GlobalTanks {
		massage = append(massage, tankMassage{
			Index:     tank.Index,
			Hp:        tank.Hp,
			Dx:        int(tank.Collider.Position.X),
			Dy:        int(tank.Collider.Position.Y),
			Direction: tank.Direction,
			weapon:    Weapon,
			Attack:    tank.Attack,
			Move:      tank.Move,
			Camp:      tank.Camp,
		})
	}
	date := json.MarshalToString(massage)
	return date

}

func (a *networkServer) Receive(m string) {
	//fmt.Println("--------------")
}
