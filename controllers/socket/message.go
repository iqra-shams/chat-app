package socket

import (
	"github.com/iqra-shams/chat-app/utils"
)

func (s *Server) HandlePrivateMessages(msg *utils.Message, member *utils.Memeber) {
	receiver, found := s.clients[msg.Receiver]
	if !found {
		member.Connection.WriteJSON("Receiver not found")
	}
	receiver.Connection.WriteJSON(msg.Message)
}
