package main

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
)

type Client struct {
	id       int
	addr     string
	username string
	conn     net.Conn
	send     chan []byte
	last_contract_id int
}

var (
	start_id int
	all_user map[int]*Client
)

func Handler(client *Client) {
	defer delete(all_user, client.id)
	defer close(client.send)
	defer ALLNotify([]byte("from 系统：" + strconv.Itoa(client.id) + " 下线了"))
	client.conn.Write([]byte("欢迎光临很大的龙的WW聊天室，您的WW号是" + strconv.Itoa(client.id)))
	for {
		message := make([]byte, 100)
		length, err := client.conn.Read(message)
		if err != nil {
			return
		}
		go MessageFunc(client.id, message[:length])
		fmt.Println(string(message[:length]))
	}
}

func MessageFunc(id int, msg []byte) {
	if string(msg[0]) == "g" {
		ALLNotify([]byte("from 公共频道 ：" +  string(msg[1:])), id)
		return
	}
	re, _ := regexp.Compile("^\\d+")
	ret := re.Find(msg)
	fmt.Println("ret  " + string(ret))
	if length := len(ret); length > 0 {
		to_id, _ := strconv.Atoi(string(ret))
		if to_user, ok := all_user[to_id]; ok {
			//存在
			post := fmt.Sprintf("from %d: %s", id, string(msg[length:]))
			all_user[id].last_contract_id = to_user.id
			to_user.send <- []byte(post)
		} else {
			all_user[id].conn.Write([]byte("from 系统：找不到这个人"))
		}
	} else {
		if all_user[id].last_contract_id > 0{
			if to_user, ok := all_user[all_user[id].last_contract_id]; ok {
				//存在
				post := fmt.Sprintf("from %d: %s", id, string(msg[length:]))
				to_user.send <- []byte(post)
			} else {
				all_user[id].conn.Write([]byte("from 系统：找不到这个人"))
			}
		}else{
			all_user[id].send <- []byte("from 系统：指令无效 格式为：id 内容")
		}
	}
}

func SendFunc(id int) {
	if client, ok := all_user[id]; ok {
		//存在
		for {
			message := <-client.send
			client.conn.Write(message)
		}
	}
}

func main() {
	all_user = make(map[int]*Client)
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		start_id++
		addr := conn.RemoteAddr().String()
		client := &Client{start_id, addr, addr, conn, make(chan []byte), 0}
		ALLNotify([]byte("from 系统：" + strconv.Itoa(start_id) + " 上线了"))
		all_user[start_id] = client
		fmt.Println(client)
		go Handler(client)
		go SendFunc(start_id)
	}
}

func ALLNotify(msg []byte, ex ... int) {
	for _, v := range all_user {
		send := true
		for _, id := range ex {
			if v.id == id{
				send = false
			}
		}
		if send{
			v.send <- msg
		}
	}
}
