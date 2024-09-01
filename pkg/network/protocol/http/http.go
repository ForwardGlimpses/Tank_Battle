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

func (a *factory) Server(ip string, port int) (types.SendChan, types.ReceiveChan, error) {
	var sendData []byte
	go func() {
		for sendData = range send {
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Read error", http.StatusInternalServerError)
			return
		}
		receive <- data
		w.WriteHeader(http.StatusOK)
		w.Write(sendData)
	})

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", ip, port),
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return send, receive, nil
}

func (a *factory) Client(ip string, port int) (types.SendChan, types.ReceiveChan, error) {
	client := &http.Client{}

	go func() {
		for data := range send {
			resp, err := client.Post(fmt.Sprintf("http://%s:%d/data", ip, port), "application/octet-stream", bytes.NewReader(data))
			if err != nil {
				fmt.Println("Client write error: ", err)
			}

			receiveData, err := io.ReadAll(resp.Body)
			resp.Body.Close() // 确保关闭响应体
			if err != nil {
				fmt.Println("Client read error: ", err)
				continue
			}
			receive <- receiveData
		}
	}()

	return send, receive, nil
}
