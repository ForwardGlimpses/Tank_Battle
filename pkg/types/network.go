package types

type NetworkManager interface {
	Send() string
	Receive(string)
}

type SendChan chan<- []byte
type ReceiveChan <-chan []byte

type ProtocolFactory interface {
	Server(ip string, port int) (SendChan, ReceiveChan, error)
	Client(ip string, port int) (SendChan, ReceiveChan, error)
}
