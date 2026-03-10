package session

import (
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

func (m *Manager) RegisterClient(sessionID int64, conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[sessionID] = conn
}

func (m *Manager) UnregisterClient(sessionID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.clients, sessionID)
}

func (m *Manager) GetSession(sessionID int64) *websocket.Conn {
	return m.clients[sessionID]
}
