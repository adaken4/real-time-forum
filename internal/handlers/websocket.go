package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"real-time-forum/internal/auth"
	"real-time-forum/internal/realtime"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// To study and implement origin checking
		return true
	},
	HandshakeTimeout: 10 * time.Second,
	EnableCompression: true,
}

var rtManager = realtime.NewRealTimeManager()

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	userID, ok := auth.GetUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println(userID)

	// Upgrade to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Websocket upgrade error: %v\n", err)
		http.Error(w, "Failed to upgrade", http.StatusInternalServerError)
		return
	}

	// Configure connection
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// Start ping ticker
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}()

	fmt.Printf("New WebSocket connection: User %s (%s)\n", userID, r.RemoteAddr)
	defer func() {
		rtManager.UnregisterClient(userID)
		conn.Close()
		fmt.Printf("Connection closed: User %s (%s)\n", userID, r.RemoteAddr)
	}()

	rtManager.RegisterClient(userID, conn)

	// Handle messages
	for {
		_, rawMessage, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("WebSocket read error (user %s): %v\n", userID, err)
			break
		}

		// Parse JSON message
		var message map[string]interface{}
		err = json.Unmarshal(rawMessage, &message)
		if err != nil {
			fmt.Println("Invalid JSON recieved:", err)
			continue
		}
		fmt.Println(message)
		rtManager.Broadcast(message)
	}
}
