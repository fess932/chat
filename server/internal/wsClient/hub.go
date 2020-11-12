package wsClient

import (
	"encoding/json"
	"log"

	"github.com/looplab/fsm"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	// State is sample state, need refactoring
	//State *State
}

//func sendState(s *Message, h *Hub) {
//	h.Bo
//
//	for client := range h.Clients {
//		select {
//		case client.Send <- bState:
//		default:
//			close(client.Send)
//			delete(h.Clients, client)
//		}
//	}
//}

// отсылаем в горутине стейт всем по броадкасту
//go sendState(hub)

func (h *Hub) Run() {
	log.Println("hub runnning")

	state := createFSM()

	//var state = Message{
	//	Type: "null",
	//	Body: map[string]string{"null": "null"},
	//}

	// очень страшно, передача указателя в горутину... каналы.. надо думать
	//go sendState(&state, h)

	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			log.Println("client registered, total clients:", len(h.Clients))
			sendMessage("henlo frind", client.Send)
			sendState(client.Send, state)

		case client := <-h.Unregister:
			log.Println("client unregister")

			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case message := <-h.Broadcast:

			messageHandle(message, state, h)

			//bState, err := json.Marshal(state)
			//if err != nil {
			//	log.Fatal(err)
			//}

		}
	}
}

func sendState(ch chan []byte, state *fsm.FSM) {
	cmd := &Message{
		Type: typeCommand,
		Body: map[string]string{
			"move": state.Current(),
		},
	}

	bCmd, err := json.Marshal(cmd)
	if err != nil {
		log.Fatal(err)
	}

	ch <- bCmd

}

//func sendCommand(cmd string) {
//	command := &Message{
//		Type: typeCommand,
//		Body: map[string]string{
//			"move": cmd,
//		},
//	}
//	bMsg,
//}

func sendMessage(msg string, ch chan []byte) {
	log.Println("send mesage..", msg)

	message := &Message{
		Type: typeMessage,
		Body: map[string]string{"message": msg},
	}
	bMsg, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err, "95")
	}

	ch <- bMsg
}

func messageHandle(message []byte, state *fsm.FSM, h *Hub) {

	log.Println("handle message:", string(message))
	m := messageParse(message)

	switch m.Type {
	case typeMessage:
		for client := range h.Clients {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(h.Clients, client)
			}
		}

	case typeCommand:
		if err := state.Event(m.Body["move"]); err != nil {
			log.Println(err)
		}

		cmd := &Message{
			Type: typeCommand,
			Body: map[string]string{
				"move": state.Current(),
			},
		}

		bCmd, err := json.Marshal(cmd)
		if err != nil {
			log.Fatal(err)
		}

		for client := range h.Clients {
			select {
			case client.Send <- bCmd:
			default:
				close(client.Send)
				delete(h.Clients, client)
			}
		}
	}

}

const (
	typeCommand = "command"
	typeMessage = "message"
)

type Message struct {
	Type string            `json:"type"`
	Body map[string]string `json:"body"`
}

func messageParse(m []byte) Message {
	var message Message
	if err := json.Unmarshal(m, &message); err != nil {
		log.Fatal(err)
	}
	return message
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte, 512),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}
