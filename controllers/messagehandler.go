package controllers

import "fmt"

func HandleMessages() {
	for {
		msg := <-broadcast

		for client, username := range clients {
			if username == msg.Receiver {
				err := client.WriteJSON(msg.Message)
				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(clients, client)
				}
			}

		}
	}
}
