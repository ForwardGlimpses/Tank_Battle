package player

import (
	"fmt"
	"strconv"

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
	Operate    Operate
}

type networkClient struct{}

var (
	Uuid          string = uuid.New().String()
	networkDetect        = map[string]int{}
)

func (a *networkClient) Send() string {
	massage := []playerMassage{}
	for _, player := range globalPlayer {
		massage = append(massage, playerMassage{
			PlayerUuid: Uuid,
			Index:      player.Index,
			Operate:    player.Operate,
		})
	}
	date:=json.MarshalToString(massage)
	fmt.Println("发送数据：",date)
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
	fmt.Println("接收数据：",massage)
	for _, playermassage := range massage {
		CombinedKey := fmt.Sprintf("%s%s", playermassage.PlayerUuid, playermassage.Index)
		networkDetect[CombinedKey] = 10
		_, ok := globalPlayer[CombinedKey]
		if ok {
			globalPlayer[CombinedKey].Operate = playermassage.Operate
		} else {
			dx, _ := strconv.Atoi(playermassage.Index)
			dy, _ := strconv.Atoi(playermassage.Index)
			player := &Player{
				Tank:       tank.New("Player", (dx+2)*100, (dy+2)*100),
				PlayerUuid: playermassage.PlayerUuid,
				Index:      playermassage.Index,
				Operate:    playermassage.Operate,
			}
			globalPlayer[CombinedKey] = player
		}
	}
	// 10轮未接收数据，清除玩家数据
	var deletaPlayer []Player
	for _, player := range globalPlayer {
		CombinedKey := fmt.Sprintf("%s%s", player.PlayerUuid, player.Index)
		networkDetect[CombinedKey]--
		if networkDetect[CombinedKey] == 0 {
			deletaPlayer = append(deletaPlayer, *player)
		}
	}

	for _, player := range deletaPlayer {
		CombinedKey := fmt.Sprintf("%s%s", player.PlayerUuid, player.Index)
		delete(globalPlayer, CombinedKey)
	}
}
