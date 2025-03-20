package realtime

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type RealTimeManager struct {
	clients map[string]*websocket.Conn
	mutex   sync.Mutex // Mutex for concurrent access
}

func NewRealTimeManager() *RealTimeManager {
	return &RealTimeManager{
		clients: make(map[string]*websocket.Conn),
	}
}

func (m *RealTimeManager) RegisterClient(userID string, conn *websocket.Conn) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.clients[userID] = conn
}

func (m *RealTimeManager) UnregisterClient(userID string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if conn, exists := m.clients[userID]; exists {
		conn.Close()
		delete(m.clients, userID)
	}
}

func (m *RealTimeManager) Broadcast(message interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// Send the JSON message to all clients
	for _, client := range m.clients {
		err := client.WriteJSON(message)
		if err != nil {
			fmt.Println("Error writing JSON to WebSocket:", err)
		}
	}
}
