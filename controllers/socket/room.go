package socket

import (
	"fmt"

	"github.com/iqra-shams/chat-app/utils"
)

func (s *Server) CreateRoom(name string, member *utils.Memeber) {
	room, exist := s.room[name]

	if !exist {
		room = &Users{
			Members: make(map[*utils.Memeber]utils.Message),
		}
		s.room[name] = room

	}

	room.Members[member] = utils.Message{
		Receiver: member.Username,
		RoomName: name,
		Message:  fmt.Sprintf("You have created and joined the room %s", name),
	}

	err := member.Connection.WriteJSON(room.Members[member].Message)
	if err != nil {
		fmt.Println("Error sending room creation notification to", member.Username, ":", err)
	}

}

func (s *Server) JoinRoom(member *utils.Memeber, name string) {

	room, exist := s.room[name]
	if !exist {
		member.Connection.WriteJSON("chat room not existed")

	} else {
		room.Members[member] = utils.Message{
			Receiver: member.Username,
			RoomName: name,
			Message:  fmt.Sprintf("You have joined room %s", name),
		}

		err := member.Connection.WriteJSON(room.Members[member].Message)
		if err != nil {
			fmt.Println("Error sending room creation notification to", member.Username, ":", err)
		}
	}
}
func (s *Server) RoomChat(msg *utils.Message, member *utils.Memeber) {
	usersInroom, exist := s.room[msg.RoomName]
	if !exist {
		member.Connection.WriteJSON("chat room not existed")
		return
		

	}
	_,exist = usersInroom.Members[member]
	if !exist {
		member.Connection.WriteJSON("first join the room ")
		return
	}

	for roomMember := range usersInroom.Members {
		if roomMember.Username == member.Username {
			continue

		}
		
		
		message:= utils.RoomMessage{
			Username: member.Username,
			Message: msg.Message,
			
		}
		err := roomMember.Connection.WriteJSON(message)
		if err != nil {
			fmt.Println("Error sending broadcast message to", roomMember.Username, ":", err)
		}

	}

}
