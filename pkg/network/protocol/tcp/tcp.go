package tcp

import (
	"fmt"
	"net"
	"sync"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/network"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/types"
)

func init() {
	network.RegisterProtocol("tcp", &factory{})
}

type factory struct{}

var (
	send    = make(chan []byte, 10)
	receive = make(chan []byte, 10)
)

func (a *factory) Server(ip string, port int) (types.SendChan, types.ReceiveChan, error) {
	var wait sync.WaitGroup
	var err error
	wait.Add(1)
	go func() {
		var listener net.Listener
		var conns []net.Conn
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
		wait.Done() // 同步 err 信息
		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Accept() failed, err: ", err)
					continue
				}
				conns = append(conns, conn)
			}
		}()

		for {
			data := <-send
			for _, conn := range conns {
				conn.Write(data)
			}
		}
	}()
	wait.Wait()
	return send, receive, err
}

func (a *factory) Client(ip string, port int) (types.SendChan, types.ReceiveChan, error) {
	go func() {

	}()
	return send, receive, nil
}
