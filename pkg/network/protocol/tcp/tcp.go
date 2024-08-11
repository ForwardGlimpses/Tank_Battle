package tcp

import (
	"bufio"
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
		conns := make([]net.Conn, 0)

		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
		wait.Done()
		if err != nil {
			fmt.Println("Listen error: ", err)
			return
		}

		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Accept error: ", err)
					continue
				}
				conns = append(conns, conn)
				go func(conn net.Conn) {
					defer conn.Close()
					reader := bufio.NewReader(conn)
					for {
						message, err := reader.ReadString('\n')
						if err != nil {
							fmt.Println("Read error: ", err)
							// 删除断开的连接
							for i, c := range conns {
								if c == conn {
									conns = append(conns[:i], conns[i+1:]...)
									break
								}
							}
							return
						}
						receive <- []byte(message)
					}
				}(conn)
			}
		}()
		go func() {
			for data := range send {
				for _, conn := range conns {
					_, err := conn.Write(append(data, '\n'))
					if err != nil {
						fmt.Println("Write error: ", err)
						conn.Close()
						// 删除断开的连接
						for i, c := range conns {
							if c == conn {
								conns = append(conns[:i], conns[i+1:]...)
								break
							}
						}
					}
				}
			}
		}()
	}()

	wait.Wait()
	return send, receive, err
}

func (a *factory) Client(ip string, port int) (types.SendChan, types.ReceiveChan, error) {
	var err error
	var conn net.Conn

	conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil, nil, err
	}

	go func() {
		defer conn.Close()
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Client read error: ", err)
				return
			}
			receive <- []byte(message)
		}
	}()

	go func() {
		for data := range send {
			_, err := conn.Write(append(data, '\n'))
			if err != nil {
				fmt.Println("Client write error: ", err)
				return
			}
		}
	}()

	return send, receive, nil
}