package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	wsServer *WsServer
	send     chan []byte
	rooms    map[*Room]bool
}

func newClient(conn *websocket.Conn, wsServer *WsServer) *Client {
	return &Client{
		conn:     conn,
		wsServer: wsServer,
		send:     make(chan []byte, 256),
		rooms:    make(map[*Room]bool),
	}
}

func (client *Client) disconnect() {
	client.wsServer.unregister <- client

	for room := range client.rooms {
		room.unregister <- client
	}

	close(client.send)
	client.conn.Close()
}

func (client *Client) handleNewMessage(jsonMessage []byte) {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Println("Error unmarshal json message:", err)
	}
	message.Sender = client
	switch message.Action {
	case SendMessageAction:
		roomName := message.Target
		if room := client.wsServer.findRoomByName(roomName); room != nil {
			room.broadcast <- &message
		}
	case JoinRoomAction:
		client.handleJoinRoomMessage(message)

	case LeaveRoomAction:
		client.handleLeaveRoomMessage(message)
	}
}

func (client *Client) handleJoinRoomMessage(message Message) {
	roomName := message.Message
	room := client.wsServer.findRoomByName(roomName)
	if room == nil {
		room = client.wsServer.createRoom(roomName)
	}
	client.rooms[room] = true

	room.register <- client
}

func (client *Client) handleLeaveRoomMessage(message Message) {
	room := client.wsServer.findRoomByName(message.Message)
	if _, ok := client.rooms[room]; ok {
		delete(client.rooms, room)
	}
	room.unregister <- client
}
