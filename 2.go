//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//type Server struct {
//	ServerName string
//	ServerIP   string
//	aaa string
//}
//
//type Serverslice struct {
//	Servers []Server
//}
//
//func main() {
//	var s Serverslice
//	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
//	json.Unmarshal([]byte(str), &s)
//	fmt.Println(s)
//}

package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"servername"`
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"},
		Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
