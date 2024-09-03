package network

import "github.com/ForwardGlimpses/Tank_Battle/pkg/types"

var (
	protocolsFactorys = map[string]types.ProtocolFactory{}
)

func RegisterProtocol(protocol string, factory types.ProtocolFactory) {
	protocolsFactorys[protocol] = factory
}
