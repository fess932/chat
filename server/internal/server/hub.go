package server

import (
	"chat-server/config"
	"chat-server/internal/server/models"
	"chat-server/internal/server/wsServer"
	"context"
	"fmt"
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*wsServer.Client]bool

	// Broadcast messages to clients.
	Broadcast chan models.Message

	// Register requests from the clients.
	Register chan *wsServer.Client

	// Unregister requests from clients.
	Unregister chan *wsServer.Client
}

const welcomeMessage = "%v вошел в хаб"

func (h *Hub) notifyClientJoined(c *wsServer.Client) {
	message := &models.Message{
		Action:  models.SendMessage,
		Target:  hubName,
		Message: fmt.Sprintf(welcomeMessage, c.GetName()),
	}

	h.broadcastToClientsInHub(message.Encode())
}

func (h *Hub) registerClientInHub(c *wsServer.Client) {

	h.notifyClientJoined(c)
	h.Clients[c] = true
}

var ctx = context.Background()
var hubName = "hub"

func (h *Hub) publishHubMessage(message []byte) {
	err := config.Redis.Publish(ctx, hubName, message).Err()
	if err != nil {
		log.Println(err)
	}
}

func (h *Hub) subscribeToHubMessages() {
	pubsub := config.Redis.Subscribe(ctx, hubName)

	ch := pubsub.Channel()

	for msg := range ch {
		h.broadcastToClientsInHub([]byte(msg.Payload))
	}
}

func (h *Hub) broadcastToClientsInHub(message []byte) {
	for client := range h.Clients {
		client.Send <- message
	}
}

func (h *Hub) Run() {
	log.Println("hub runnning")

	//state := createFSM()

	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			log.Println("client registered, total clients:", len(h.Clients))
			h.notifyClientJoined(client)

		case client := <-h.Unregister:
			log.Println("client unregister")

			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case message := <-h.Broadcast:
			log.Println("broadcast message:", message)
			h.publishHubMessage(message.Encode())
		}
	}
}
