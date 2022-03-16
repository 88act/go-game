package main

import (
 
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

func add(param string) {
	ws, err := websocket.Dial("ws://127.0.0.1:8888/add", "", "http://127.0.0.1:8888/")
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close() //关闭连接

	sendMsg := []byte(param)
	_, err = ws.Write(sendMsg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", string(sendMsg))

	msg := make([]byte, 512)
	m, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:m])
}

func del(param string) {
	ws, err := websocket.Dial("ws://127.0.0.1:8888/del", "", "http://127.0.0.1:8888/")
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close() //关闭连接

	sendMsg := param
	err = websocket.Message.Send(ws, sendMsg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", sendMsg)

	var msg string
	err = websocket.Message.Receive(ws, &msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg)
}

func main() {
	add("小龙虾")
	add("鱼香肉丝")
	del("小龙虾")
}
