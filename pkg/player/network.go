package player

import (
	"fmt"

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
	Uuid    string
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
	massage := []playerMassage{}
	json.Unmarshal([]byte(m), &massage)

}
