package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// TCP 客户端
func main() {

	fmt.Println("客户端启动")
	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n')  // 读取用户输入
		input = strings.TrimSuffix(input, "\r\n") // 删除换行
		if input == "q" {
			return
		}

		conn, err := net.Dial("tcp", "127.0.0.1:9999")
		if err != nil {
			panic(err)
		}
		defer conn.Close() // 关闭TCP连接

		_, err = conn.Write([]byte(input))
		if err != nil {
			panic(err)
		}

		var byteBuf bytes.Buffer
		var buf = make([]byte, 512)
		var read = bufio.NewReader(conn)
		for {
			n, err := read.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			byteBuf.Write(buf[0:n])
		}

		fmt.Println(string(byteBuf.Bytes()))
	}
}
