package socket

import (
	"github.com/gorilla/websocket"
	"github.com/iqra-shams/chat-app/utils"
)

type Server struct {
	broadcast chan utils.Message
	clients   map[*websocket.Conn]string
}

func (s *Server) InitServer() {
	s.clients = make(map[*websocket.Conn]string)
	s.broadcast = make(chan utils.Message)
}
