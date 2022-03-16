package main
import (
	 
	"fmt"
	"log"
	"net/http" 
	"golang.org/x/net/websocket"
)

func add(ws *websocket.Conn) {
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:n])

	sendMsg := "订单添加：" + string(msg[:n])
	_, err = ws.Write([]byte(sendMsg))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", sendMsg)
}

func del(ws *websocket.Conn) {
	var msg string
	err := websocket.Message.Receive(ws, &msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg)

	sendMsg := "订单删除：" + msg
	err = websocket.Message.Send(ws, sendMsg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", sendMsg)
}

func main() {
	http.Handle("/add", websocket.Handler(add))
	http.Handle("/del", websocket.Handler(del))
	fmt.Println("开始监听8888端口...")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}
