package utils

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	RoomName string `json:"roomname"`
	Message  string `json:"message"`
	Receiver string `json:"to"`
	Action   string `json:"action"`
}
type RoomMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	
}

type Memeber struct {
	Username   string
	Connection *websocket.Conn
	Mu         sync.Mutex
}
