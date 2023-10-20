package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Setup the server interface for the websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Accepting all requests
	},
}

var clients map[*websocket.Conn]bool

type WebSocketController interface {
	StartConnect(w http.ResponseWriter, r *http.Request)
}

type webSocketController struct{}

func (c *webSocketController)StartConnect(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	clients[connection] = true // Add new client to the list of clients
	for {
		mt, message, err := connection.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			fmt.Println("==> Has 1 connection closed")
			break // Exit the loop if the client tries to close the connection or the connection with the interrupted client
		}

		go messageHandler(message)
		go writeMessage(message)
	}

	// Close the connection when the loop ends
	connection.Close()
	delete(clients, connection)
}

func messageHandler(message []byte) {
	fmt.Println(string(message))
}

func writeMessage(message []byte) {
	for conn := range clients {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}

func NewWebSocketController() WebSocketController {
	return &webSocketController{}
}