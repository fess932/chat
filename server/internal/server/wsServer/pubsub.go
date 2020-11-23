package wsServer

import (
	"chat-server/config"
	"chat-server/internal/server/models"
	"log"
)

func (s *WsServer) publishClientJoined(c *Client) {
	message := &models.Message{Action: models.UserJoined, Sender: c}
	if err := config.Redis.Publish(ctx, PubSubGeneralChannel, message.Encode()).Err(); err != nil {
		log.Println(err)
	}
}

func (s *WsServer) publishClientLeft(c *Client) {
	message := &models.Message{Action: models.UserLeft, Sender: c}
	if err := config.Redis.Publish(ctx, PubSubGeneralChannel, message.Encode()).Err(); err != nil {
		log.Println(err)
	}
}
