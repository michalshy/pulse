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
		log.Println(err)
		return
	}
	var msg models.HandshakeMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		log.Println(err)
		return
	}
	sessionID, err := db.CreateSession(context.Background(), msg.ProjectID, msg.Metadata)
	if err != nil {
		log.Println(err)
		return
	}
	h.Manager.RegisterSession(sessionID, conn)

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
			ctx := context.Background()
			var flush models.FlushPayload
			if err := json.Unmarshal(data, &flush); err != nil {
				log.Println(err)
				continue
			}

			captureID, err := db.CreateCapture(ctx, sessionID)
			if err != nil {
				log.Println(err)
				continue
			}

			for _, metric := range flush.Metrics {
				if _, err := db.CreateMetric(ctx, captureID, metric.RecordedAt, metric.Name, metric.ValueType, metric.Value); err != nil {
					log.Println(err)
				}
			}

			for _, event := range flush.Events {
				if _, err := db.CreateEvent(ctx, captureID, event.RecordedAt, event.Name, event.Metadata); err != nil {
					log.Println(err)
				}
			}
		}
	}

	h.Manager.UnregisterSession(sessionID)
	db.EndSession(context.Background(), sessionID)
}
