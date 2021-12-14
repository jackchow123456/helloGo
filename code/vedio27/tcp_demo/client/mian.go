package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client demo
func main() {
	// 1. 与服务端建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("connet failed,error:%v\n", err)
		return
	}

	// 2.利用该链接惊醒数据的发送和接收
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}

		// 给服务端发消息
		_, err := conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send failed,error:%v\n", err)
			return
		}

		// 从服务端接收消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed,error:%v\n", err)
			return
		}

		fmt.Printf("收到服务端回复：%v\n", string(buf[:n]))
	}
}
