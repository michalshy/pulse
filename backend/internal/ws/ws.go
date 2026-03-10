package ws

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"pulse/internal/db"
	"pulse/internal/models"
	"pulse/internal/session"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Handler struct {
	Manager *session.Manager
}

func (h *Handler) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// first message must be handshake!
	_, data, err := conn.ReadMessage()
	if err != nil {
		return
	}
	var msg models.HandshakeMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		return
	}
	sessionID, err := db.CreateSession(context.Background(), msg.GameID, msg.Metadata)
	if err != nil {
		return
	}
	h.Manager.RegisterClient(sessionID, conn)

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// unmarshall the message
		var base models.BaseMessage
		if err := json.Unmarshal(data, &base); err != nil {
			log.Println(err)
			continue
		}

		switch base.Type {
		case "flush":
			// handle flush
		case "trigger":
			// handle trigger
		}
	}

	h.Manager.UnregisterClient(sessionID)
}
