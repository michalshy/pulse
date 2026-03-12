package main

import (
	"bufio"
	"encoding/json"
	"fakesdk/models"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	//connect to websocket
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("Connected to server")

	// write message
	var Message models.HandshakeMessage
	Message.Type = "handshake"
	Message.ProjectID = "test_game"

	data, err := json.Marshal(Message)
	if err != nil {
		log.Fatal(err)
	}
	conn.WriteMessage(websocket.TextMessage, data)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		flush := models.FlushPayload{
			Type: "flush",
			Metrics: []models.MetricPayload{
				{Name: "fps", RecordedAt: time.Now(), ValueType: "float", Value: json.RawMessage(`87.4`)},
				{Name: "memory_mb", RecordedAt: time.Now(), ValueType: "int", Value: json.RawMessage(`412`)},
			},
			Events: []models.EventPayload{},
		}
		data, _ = json.Marshal(flush)
		conn.WriteMessage(websocket.TextMessage, data)
		log.Println("Flush sent")
	}
}
