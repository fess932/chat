package main

import (
	"encoding/json"
	"log"
)

const SendMessageAction = "send-message"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"

type Message struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	Target  string `json:"target"`
	Sender  *Client
}

func (message *Message) encode() []byte {
	jsn, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return jsn
}
