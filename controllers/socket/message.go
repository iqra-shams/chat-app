package socket

import "fmt"

func (s *Server) HandleMessages() {
	for {
		msg := <-s.broadcast

		for client, username := range s.clients {
			if username == msg.Receiver {
				err := client.WriteJSON(msg.Message)
				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(s.clients, client)
				}
			}

		}
	}
}
