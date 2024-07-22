package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	var buf = make([]byte, 512)
	reader := bufio.NewReader(conn)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	data := string(buf[0:n])
	fmt.Println("收到Client端发来的数据：", data)

	conn.Write([]byte("tcp:127.0.0.1:9999"))
}

func main() {
	// 2.2.2.2
	listen, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}
	fmt.Println("服务器启动")
	for {
		conn, err := listen.Accept() // 监听客户端的连接请求
		if err != nil {
			fmt.Println("Accept() failed, err: ", err)
			continue
		}
		go process(conn) // 启动一个goroutine来处理客户端的连接请求
	}
}
