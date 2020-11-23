package wsServer

import (
	"chat-server/config"
	"chat-server/internal/server/models"
	"context"
	"log"
	"net/http"
)

type WsServer struct {
	users   []models.User
	clients map[*Client]bool

	Register   chan *Client
	Unregister chan *Client

	Broadcast chan models.Message
}

func (s *WsServer) registerClient(c *Client) {

	s.publishClientJoined(c)
	s.listOnlineClients(c)

	s.clients[c] = true
}

func (s *WsServer) unregisterClient(c *Client) {
	if _, ok := s.clients[c]; ok {
		delete(s.clients, c)

		//Publish user left in PubSub
		s.publishClientJoined(c)
	}
}

func (s *WsServer) broadcastToClients(message []byte) {
	for c := range s.clients {
		c.Send <- message
	}
}

var ctx = context.Background()

const PubSubGeneralChannel = "general"

func (s *WsServer) listenPubSubChannel() {
	pubsub := config.Redis.Subscribe(ctx, PubSubGeneralChannel)

	ch := pubsub.Channel()

	for msg := range ch {

		message, err := models.MessageParse([]byte(msg.Payload))
		if err != nil {
			log.Print(err)
			return
		}

		switch message.Action {
		case models.UserJoined:
			s.handleUserJoined(message)
		case models.UserLeft:
			s.handleUserLeft(message)
		}
	}
}

func (s *WsServer) Run() {
	go s.listenPubSubChannel()
	for {
		select {
		case c := <-s.Register:
			s.registerClient(c)
		case c := <-s.Unregister:
			s.unregisterClient(c)
		}
	}
}

func (s *WsServer) listOnlineClients(c *Client) {
	for _, user := range s.users {
		message := models.Message{Action: models.UserJoined, Sender: user}
		c.Send <- message.Encode()
	}
}

func NewWebsocketServer() *WsServer {
	wsServer := &WsServer{
		clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		users:      []models.User{},
	}

	return wsServer
}

func (s *WsServer) ServeWS(w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	c := newClient(conn, s)

	go c.WritePump()
	go c.ReadPump()

	s.Register <- c

	return nil
}
