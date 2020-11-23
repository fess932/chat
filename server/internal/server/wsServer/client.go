package wsServer

import (
	"chat-server/internal/server/models"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/gorilla/websocket"
)

const (
	// Время доступное для записи сообщения на клиента
	writeWait = 10 * time.Second

	// Время доступное на чтение следующего понг сообщения от клиента
	pongWait = 60 * time.Second

	// Отправка пингов к клиенту в этот период. должно быть меньше чем pongWait
	pingPeriod = (pongWait * 9) / 10

	// Максимальный размер сообщений от клиента
	maxMessageSize = 512
)

// Client is a middleman between the server connection and the hub.
type Client struct {

	// The server connection.
	conn     *websocket.Conn
	wsServer *WsServer

	// Buffered channel of outbound messages.
	Send chan []byte

	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func newClient(conn *websocket.Conn, wsServer *WsServer) *Client {
	return &Client{
		conn:     conn,
		wsServer: wsServer,
		Send:     make(chan []byte, 256),
		ID:       uuid.New(),
		Name:     "noname",
	}
}

////////////////////////////////////////////////////////////////////////

func (c *Client) handleNewMessage(jsonMessage []byte) {
	message, err := models.MessageParse(jsonMessage)
	if err != nil {
		log.Println(err)
		return
	}
	message.Sender = c

	c.wsServer.Broadcast <- message
}

func (c *Client) disconnect() {
	c.wsServer.Unregister <- c
	close(c.Send)
	c.conn.Close()
}

func (c *Client) GetID() string {
	return c.ID.String()
}

func (c *Client) GetName() string {
	return c.Name
}
