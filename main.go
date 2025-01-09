package main

import (
	"fmt"
	"net"
	"time"
)

func scanPort(ip string, port int) {
	address := fmt.Sprintf("%s:%d", ip, port)
	//接続のタイムアウト時間を設定
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		//エラー発生時
		fmt.Printf("Port %d is closed or filterd\n", port)
		return
	}
	defer conn.Close()

	fmt.Printf("Port %d is open\n", port)
}

func scanPorts(ip string, startPort, endPort int){
	for port := startPort; port <= endPort; port++ {
		go scanPort(ip, port)
	}
}

func main(){
	ip := "127.0.0.1"
	startPort := 1
	endPort := 100
	scanPorts(ip, startPort, endPort)

	time.Sleep(2 * time.Second)
}
