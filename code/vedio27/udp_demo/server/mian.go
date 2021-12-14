package main

import (
	"fmt"
	"net"
)

func main() {
	// 1. 开启udp服务
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("listen udp failed,error:%v\n", err)
		return
	}

	// 2. 延迟关闭udp服务
	defer listen.Close()

	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("listen read from udp failed,error:%v\n", err)
			return
		}
		fmt.Printf("接收到的数据：%v\n", string(buf[:n]))

		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Printf("write to udp %v failed,error:%v\n", addr, err)
			return
		}
	}
}
