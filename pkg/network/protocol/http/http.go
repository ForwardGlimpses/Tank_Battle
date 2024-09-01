package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ForwardGlimpses/Tank_Battle/pkg/network"
	"github.com/ForwardGlimpses/Tank_Battle/pkg/types"
)

func init() {
	network.RegisterProtocol("http", &factory{})
}

type factory struct{}

var (
	sendQueue []byte
	send      = make(chan []byte, 10)
	receive   = make(chan []byte, 10)
)

func (a *factory) Server(ip string, port int) (types.SendChan, types.ReceiveChan, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Read error", http.StatusInternalServerError)
				return
			}
			receive <- data
			w.WriteHeader(http.StatusOK)
		} else if r.Method == http.MethodGet {
			if len(sendQueue) > 0 {
				w.Write(sendQueue)
				sendQueue = nil
			} else {
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
	Out:
		for {
			select {
			case data := <-send:
				if len(data) == 0 {
					break Out
				}
				sendQueue = data
			case <-time.After(time.Second * 5):
				fmt.Println("等待数据超时")
				break Out
			}
		}
	}()

	return send, receive, nil
}

func (a *factory) Client(ip string, port int) (types.SendChan, types.ReceiveChan, error) {
	client := &http.Client{}

	go func() {
		for data := range send {
			_, err := client.Post(fmt.Sprintf("http://%s:%d/data", ip, port), "application/octet-stream", bytes.NewReader(data))
			if err != nil {
				fmt.Println("Client write error: ", err)
			}
		}
	}()

	go func() {
		for {
			resp, err := client.Get(fmt.Sprintf("http://%s:%d/data", ip, port))
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
			if len(data) > 0 {
				receive <- data
			}
		}
	}()

	return send, receive, nil
}
