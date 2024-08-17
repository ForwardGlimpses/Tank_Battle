package network

import (
	"encoding/json"
	"errors"
	"fmt"

	//"github.com/ForwardGlimpses/Tank_Battle/pkg/config"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/config"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/tankbattle"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/types"
)

var (
	serverManagers = map[string]types.NetworkManager{}
	clientManagers = map[string]types.NetworkManager{}
	sendCh         types.SendChan
	receiveCh      types.ReceiveChan
	duration       = 10
	model          int
)

const (
	clientModel = iota
	serverModel
	noneModel
)

func RegisterClient(key string, manager types.NetworkManager) {
	clientManagers[key] = manager
}

func RegisterServer(key string, manager types.NetworkManager) {
	serverManagers[key] = manager
}

func init() {
	tankbattle.RegisterInit(Init, 4)
	tankbattle.RegisterUpdate(Update, 4)
}

func Init() (err error) {
	cfg := config.C.Network
	switch cfg.Type {
	case "client":
		protocol := protocolsFactorys[cfg.Protocol]
		if protocol == nil {
			err = errors.New("unknow protocol")
			return
		}
		sendCh, receiveCh, err = protocol.Client(cfg.IP, cfg.Port)
		if err != nil {
			return err
		}
		model = clientModel
	case "server":
		protocol := protocolsFactorys[cfg.Protocol]
		if protocol == nil {
			err = errors.New("unknow protocol")
			return
		}
		sendCh, receiveCh, err = protocol.Server(cfg.IP, cfg.Port)
		if err != nil {
			return err
		}
		model = serverModel
	default:
		model = noneModel
	}
	fmt.Println(model)
	return
}

func Update() {
	if model == noneModel {
		return
	}

	if duration > 0 {
		duration--
		return
	} else {
		duration = 10
	}

	sendMassage()
	receiveMassage()
}

func sendMassage() {
	massage := map[string]string{}
	if model == clientModel {
		for key, client := range clientManagers {
			m := client.Send()
			if m != "" {
				massage[key] = m
			}
		}
	}
	if model == serverModel {
		for key, server := range serverManagers {
			m := server.Send()
			if m != "nil" {
				massage[key] = m
			}
		}
	}

	data, _ := json.Marshal(massage)
	sendCh <- data
}

func receiveMassage() {
	var data []byte

Out:
	for {
		select {
		case data = <-receiveCh:
		default:
			break Out
		}
	}
	if len(data) == 0 {
		return
	}

	massage := map[string]string{}
	json.Unmarshal(data, &massage)

	if model == clientModel {
		for key, m := range massage {
			clientManagers[key].Receive(m)
		}
	}
	if model == serverModel {
		for key, m := range massage {
			serverManagers[key].Receive(m)
		}
	}
}
