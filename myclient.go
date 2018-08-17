package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8000")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			time.Sleep(300 * time.Microsecond)
			message := make([]byte, 100)
			length, err := conn.Read(message)
			if err != nil{
				fmt.Println("服务器停止运行")
				return
			}
			fmt.Println(string(message[:length]))
		}
	}()
	f := bufio.NewReader(os.Stdin)
	for {
		//fmt.Print("请输入发送的信息：")
		Input, _ := f.ReadString('\n')
		_, err := conn.Write([]byte(Input))
		if err != nil{
			fmt.Println("服务器停止运行")
			return
		}
	}
}
