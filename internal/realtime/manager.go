package realtime

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type RealTimeManager struct {
	clients       map[string]*websocket.Conn
	mutex         sync.Mutex // Mutex for concurrent access
	connections   int
	totalMessages int
}

func NewRealTimeManager() *RealTimeManager {
	return &RealTimeManager{
		clients: make(map[string]*websocket.Conn),
	}
}

func (m *RealTimeManager) RegisterClient(userID string, conn *websocket.Conn) {
	m.mutex.Lock()
	m.clients[userID] = conn
	m.mutex.Unlock()

	// Send updated user list to all clients
	m.broadcastUserList()
}

func (m *RealTimeManager) UnregisterClient(userID string) {
	m.mutex.Lock()
	if conn, exists := m.clients[userID]; exists {
		conn.Close()
		delete(m.clients, userID)
	}
	m.mutex.Unlock()

	// Send updated user list to all clients
	m.broadcastUserList()
}

// Send a message to a specific user
func (m *RealTimeManager) SendPrivateMessage(senderID, receiverID string, message interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	fmt.Println(senderID, receiverID, message)

	if conn, exists := m.clients[receiverID]; exists {
		err := conn.WriteJSON(map[string]interface{}{
			"type":    "private_message",
			"from":    senderID,
			"message": message,
		})
		if err != nil {
			fmt.Println("Error sending private message:", err)
		}
	}
}

// Get a list of connected users
func (m *RealTimeManager) GetOnlineUsers() []string {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	users := make([]string, 0, len(m.clients))
	for userID := range m.clients {
		users = append(users, userID)
	}
	return users
}

func (m *RealTimeManager) broadcastUserList() {
	users := m.GetOnlineUsers()
	for _, conn := range m.clients {
		err := conn.WriteJSON(map[string]interface{}{
			"type":  "user_list",
			"users": users,
		})
		if err != nil {
			fmt.Println("Error broadcasting user list:", err)
		}
	}
}

func (m *RealTimeManager) Broadcast(senderID string, message interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	fmt.Printf("Broadcasting to %d clients\n", len(m.clients))
	m.totalMessages++

	// Send the JSON message to all clients
	for userID, client := range m.clients {
		if userID == senderID {
			continue
		}

		err := client.WriteJSON(message)
		if err != nil {
			fmt.Println("Error writing JSON to WebSocket:", err)
		}
	}
}
