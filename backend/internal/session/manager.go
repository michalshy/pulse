package session

import (
	"encoding/json"
	"fmt"
	"pulse/internal/models"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	clients map[int64]*websocket.Conn
	mu      sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(map[int64]*websocket.Conn),
	}
}

func (m *Manager) RegisterSession(sessionID int64, conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[sessionID] = conn
}

func (m *Manager) UnregisterSession(sessionID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.clients, sessionID)
}

func (m *Manager) GetSession(sessionID int64) *websocket.Conn {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.clients[sessionID]
}

func (m *Manager) SendTrigger(sessionID int64) error {
	m.mu.RLock()
	conn := m.clients[sessionID]
	m.mu.RUnlock()

	if conn == nil {
		return fmt.Errorf("session %d not active", sessionID)
	}

	data, err := json.Marshal(models.TriggerMessage{Type: "trigger"})
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.TextMessage, data)
}
