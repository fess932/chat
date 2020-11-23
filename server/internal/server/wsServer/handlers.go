package wsServer

import "chat-server/internal/server/models"

func (s *WsServer) handleUserJoined(message models.Message) {
	s.users = append(s.users, message.Sender)
	s.broadcastToClients(message.Encode())
}

func (s *WsServer) handleUserLeft(message models.Message) {
	// remove the user from the slice
	for i, user := range s.users {
		if user.GetID() == message.Sender.GetID() {
			s.users[i] = s.users[len(s.users)-1]
			s.users = s.users[:len(s.users)-1]
		}
	}

	s.broadcastToClients(message.Encode())
}
