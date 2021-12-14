package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server demo

func process(conn net.Conn) {
	defer conn.Close() //处理完之后要关闭这个链接
	// 针对当前的链接做数据的发送和接收操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, error:%v\n", err)
			break
		}
		recv := buf[:n]
		fmt.Printf("接收到的数据:%v\n", recv)

		conn.Write([]byte("ok")) // 返回客户端
	}
}

func main() {

	// 1. 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, error:%v\n", err)
		return
	}

	for {
		// 2.等待客户端来连杰
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, error:%v\n", err)
			continue
		}

		// 3.启动一个单独的goruting去处理链接
		go process(conn)
	}
}
