package socket

import (
	"github.com/iqra-shams/chat-app/utils"
)

type Server struct {
	broadcast chan utils.Message
	clients   map[string]*utils.Memeber
	room      map[string]*Users
}

func (s *Server) InitServer() {
	s.clients = make(map[string]*utils.Memeber)
	s.broadcast = make(chan utils.Message)
	s.room = make(map[string]*Users)
}

type Users struct {
	Members map[*utils.Memeber]utils.Message
}
