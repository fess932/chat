package models

import (
	"encoding/json"
	"log"
)

// actions
const (
	SendCommand = iota
	SendMessage

	UserJoined
	UserLeft
)

type Message struct {
	Action  int    `json:"action"`
	Message string `json:"message"`
	Target  string `json:"target"`
	Sender  User   `json:"sender"`
}

func (m *Message) Encode() []byte {
	jsn, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
	}

	return jsn
}

func MessageParse(bin []byte) (m Message, err error) {
	err = json.Unmarshal(bin, &m)
	return
}
