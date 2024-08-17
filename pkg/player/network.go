package player

import (
	"github.com/ForwardGlimpses/Tank_Battle/pkg/network"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tank"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/json"
	"github.com/google/uuid"
)

func init() {
	network.RegisterClient("player", &networkClient{})
	network.RegisterServer("player", &networkServer{})
}

type playerMassage struct {
	PlayerUuid string
	Index      string
	Action     Action
	TankIndex  int
}

type networkClient struct{}

var (
	Uuid string = uuid.New().String()
)

func (a *networkClient) Send() string {
	massage := []playerMassage{}
	for _, player := range globalPlayer {
		massage = append(massage, playerMassage{
			PlayerUuid: Uuid,
			Index:      player.Index,
			Action:     player.Action,
		})
	}
	date := json.MarshalToString(massage)
	//fmt.Println("发送数据：",date)
	return date
}

func (a *networkClient) Receive(m string) {}

type networkServer struct{}

func (a *networkServer) Send() string {
	return ""
}

func (a *networkServer) Receive(m string) {
	massage := []playerMassage{}
	json.Unmarshal([]byte(m), &massage)

	for _, playermassage := range massage {
		player, ok := globalPlayer[playermassage.Index]
		if ok {
			player.Action = playermassage.Action
			player.NetworkCount = 10
			player.TankIndex = playermassage.TankIndex
		} else {
			player := &Player{
				TankIndex:    tank.New("Player", 100, 100).Index,
				Index:        playermassage.Index,
				Action:       playermassage.Action,
				NetworkCount: 10,
			}
			globalPlayer[playermassage.Index] = player
		}
	}
	// 10轮未接收数据，清除玩家数据
	var deletaPlayer []Player
	for _, player := range globalPlayer {
		player.NetworkCount--
		if player.NetworkCount <= 0 {
			deletaPlayer = append(deletaPlayer, *player)
		}
	}

	for _, player := range deletaPlayer {

		delete(globalPlayer, player.Index)
	}
}
