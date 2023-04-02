package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Msg struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Body     string `json:"body"`
}

type Connection struct {
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[string]*Connection)

func ServeChatting(w http.ResponseWriter, r *http.Request) {
	// connect
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	msg := new(Msg)
	for {
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Read Error : %s", err)
			break
		}

		log.Printf("[%s] %s : %s", msg.Type, msg.Username, msg.Body)

		switch {
		case msg.Type == "join":
			clients[msg.Username] = &Connection{conn: conn}
		case msg.Type == "leave":
			delete(clients, msg.Username)
		default:
			log.Println(clients)
		}
		// boardcast
		for username, client := range clients {
			err := client.conn.WriteJSON(msg)
			if err != nil {
				log.Printf("[%s] Write Error : %s", username, err)
				delete(clients, username)
			}
		}
		if err != nil {
			log.Printf("Write Error : %s", err)
			break
		}
	}
}
