package player

import (
	"fmt"
	"strconv"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/network"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/utils/json"
	"github.com/google/uuid"
)

func init() {
	network.RegisterClient("player", &networkClient{})
	network.RegisterServer("player", &networkServer{})
}

type playerMassage struct {
	Index   string
	Operate Operate
}

type networkClient struct{}

var Uuid string = uuid.New().String()

func (a *networkClient) Send() string {
	massage := []playerMassage{}
	for _, player := range globalPlayer {
		massage = append(massage, playerMassage{
			Index:   fmt.Sprintf("%s%d", Uuid, player.Index),
			Operate: player.Operate,
		})
	}
	return json.MarshalToString(massage)
}

func (a *networkClient) Receive(m string) {}

type networkServer struct{}

func (a *networkServer) Send() string {
	return ""
}

func (a *networkServer) Receive(m string) {

	//add 10轮未接收数据，清除玩家数据
	massage := []playerMassage{}
	json.Unmarshal([]byte(m), &massage)
	for _, playermassage := range massage {
		index, _ := strconv.Atoi(playermassage.Index)
		globalPlayer[index].Operate = playermassage.Operate
	}
}
