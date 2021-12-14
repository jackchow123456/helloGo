package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp client demo
func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("dial from udp failed,error:%v\n", err)
		return
	}
	defer conn.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		s, _ := input.ReadString('\n')
		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Printf("send to server failed,error:%v\n", err)
			return
		}

		// 接收数据
		var buf [1024]byte
		n, addr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Printf("recv from server failed,error:%v\n", err)
			return
		}
		fmt.Printf("read from %s, msg:%v\n", addr, string(buf[:n]))
	}

}
