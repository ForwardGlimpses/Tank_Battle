package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/network"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/types"
)

func init() {
	network.RegisterProtocol("http", &factory{})
}

type factory struct{}

var (
	send    = make(chan []byte, 10)
	receive = make(chan []byte, 10)
)

// 服务器端处理逻辑
func (a *factory) Server(ip string, port int) (types.SendChan, types.ReceiveChan, error) {
	mux := http.NewServeMux()

	// 处理接收数据的 HTTP 请求
	mux.HandleFunc("/receive", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            data, err := io.ReadAll(r.Body)
            if err != nil {
                http.Error(w, "Read error", http.StatusInternalServerError)
                return
            }
            select {
            case receive <- data:
                w.WriteHeader(http.StatusOK)
            default:
                http.Error(w, "Receive channel is full", http.StatusInternalServerError)
            }
        } else if r.Method == http.MethodGet {
            select {
            case data := <-receive:
                w.Write(data)
            default:
                w.WriteHeader(http.StatusNoContent)
            }
        } else {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        }
    })
    

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", ip, port),
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error: ", err)
		}
	}()

	go func() {
		for data := range send {
			_, err := http.Post(fmt.Sprintf("http://%s:%d/receive", ip, port), "application/octet-stream", bytes.NewReader(data))
			if err != nil {
				fmt.Println("Send error: ", err)
			}
		}
	}()

	return send, receive, nil
}

// 客户端处理逻辑
func (a *factory) Client(ip string, port int) (types.SendChan, types.ReceiveChan, error) {
	client := &http.Client{}

	go func() {
		for data := range send {
			_, err := client.Post(fmt.Sprintf("http://%s:%d/receive", ip, port), "application/octet-stream", bytes.NewReader(data))
			if err != nil {
				fmt.Println("Client write error: ", err)
			}
		}
	}()

	go func() {
		for {
			resp, err := client.Get(fmt.Sprintf("http://%s:%d/receive", ip, port))
			if err != nil {
				fmt.Println("Client read error: ", err)
				continue
			}
			if resp.StatusCode == http.StatusNoContent {
				resp.Body.Close()
				continue
			}
			data, err := io.ReadAll(resp.Body)
			resp.Body.Close() // 确保关闭响应体
			if err != nil {
				fmt.Println("Client read error: ", err)
				continue
			}
			receive <- data
		}
	}()

	return send, receive, nil
}
