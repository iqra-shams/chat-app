package socket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/iqra-shams/chat-app/utils"
	"github.com/iqra-shams/chat-app/validation"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Server) HandleSocketConnection(w http.ResponseWriter, r *http.Request) {

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer"):]
	claims, err := validation.ValidateToken(tokenString)
	if err != nil {
		fmt.Fprintln(w, err)
		http.Error(w, "err in Validate token", http.StatusUnauthorized)
		return

	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	
	member := utils.Memeber{
		Username:   claims.Username,
		Connection: conn,
	}

	s.clients[""] =  &utils.Memeber{
		Username:   claims.Username,
		Connection: conn,
	}
	

	for {
		var msg utils.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(s.clients, claims.Username)
			return
		}
		switch msg.Action {
		case "creatroom":
			s.CreateRoom(msg.RoomName, &member)
		case "DM":
			s.HandlePrivateMessages(&msg, &member)
		case "joinroom":
			s.JoinRoom(&member, msg.RoomName)
		case "roomchat":
			s.RoomChat(&msg,&member)

		}

	}

}
